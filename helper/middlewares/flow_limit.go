package middlewares

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/flow_limit"
	"net"
	"net/http"
	"strings"
)

type Limiter string

const (
	ObsFlowLimiter Limiter = "AntiFake-Obs-PeriodLimit-" // Obs上传限流器
	SMSFlowLimiter Limiter = "AntiFake-SMS-PeriodLimit-" // sms限流器
)

type (
	FlowLimitMiddleware struct {
		Limiter
		host            string // redis host
		pwd             string // redis pwd
		flow_limit.Conf        // 限流参数
	}

	Req struct { // 自定义字段
	}
)

var FlowLimitMap = map[Limiter]string{
	ObsFlowLimiter: string(ObsFlowLimiter),
	SMSFlowLimiter: string(SMSFlowLimiter),
}

func NewFlowLimitMiddleware(conf redis.RedisConf, conf2 flow_limit.Conf, limiter Limiter) *FlowLimitMiddleware {
	return &FlowLimitMiddleware{
		Limiter: limiter,
		host:    conf.Host,
		pwd:     conf.Pass,
		Conf:    conf2,
	}
}

func (m *FlowLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err error
		)

		pass, err := m.Check(m.resolveHost(r))
		if err != nil || !pass {
			logrus.Errorf("RemoteAddr:%s,realIP:%s:X-Forward-For:%s,err:%v", r.RemoteAddr, r.Header.Get("X-Real-IP"), r.Header.Get("X-Forward-For"), err)
			res, _ := json.Marshal(map[string]interface{}{
				"code": http.StatusTooManyRequests,
				"msg":  "请求次数超过限制,请稍后再尝试",
			})
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write(res)
			return
		}

		next(w, r)
	}
}

func (m *FlowLimitMiddleware) Check(key string) (bool, error) {
	l := limit.NewPeriodLimit(m.PeriodSec, m.Quota, redis.New(m.host, redis.WithPass(m.pwd)), FlowLimitMap[m.Limiter])
	code, err := l.Take(key)
	if err != nil {
		logrus.Errorf("take out key failure[key:%s,err:%s]", key, err.Error())
		return false, err
	}

	switch code {
	case limit.OverQuota:
		logrus.Errorf("OverQuota key: %v", key)
		return false, err
	case limit.Allowed:
		logrus.Infof("AllowedQuota key: %v", key)
		return true, nil
	case limit.HitQuota: // 刚好达到限额,不过也是允许用户访问的
		logrus.Infof("HitQuota key: %v", key)
		return true, err
	default:
		logrus.Errorf("DefaultQuota key: %v", key)
		// unknown response, we just let the request go
		return true, nil
	}
}

func (m *FlowLimitMiddleware) resolveHost(r *http.Request) (host string) {
	ip := r.Header.Get("X-Real-IP") // 只包含客户端机器的一个IP，如果为空，某些代理服务器（如Nginx）会填充此header
	if net.ParseIP(ip) != nil {
		return ip
	}
	ip = r.Header.Get("X-Forward-For") // 一系列的IP地址列表，以,分隔，每个经过的代理服务器都会添加一个IP。
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr) // 包含客户端的真实IP地址。 这是Web服务器从其接收连接并将响应发送到的实际物理IP地址。 但是，如果客户端通过代理连接，它将提供代理的IP地址。
	if err != nil {
		logrus.Errorf("net.SplitHostPort failure:{remoteAddr:%s,err:%s}", r.RemoteAddr, err.Error())
		return ""
	}

	if net.ParseIP(ip) != nil {
		return ip
	}

	//RemoteAddr是最可靠的，但是如果客户端位于代理之后或使用负载平衡器或反向代理服务器时，它将永远不会提供正确的IP地址，因此顺序是先是X-REAL-IP，然后是X-FORWARDED-FOR，然后是 RemoteAddr。
	//请注意，恶意用户可以创建伪造的X-REAL-IP和X-FORWARDED-FOR标头。
	return ""
}
