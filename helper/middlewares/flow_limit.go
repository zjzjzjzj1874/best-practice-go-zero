package middlewares

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/flow_limit"
	"net"
	"net/http"
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
			err  error
			host = r.Header.Get("X-Real-IP")
		)

		if host == "" {
			host, _, err = net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				logrus.Errorf("net.Split HostPort failure:[remoteAddr:%s,err:%s]", r.RemoteAddr, err.Error())
			}
		}
		pass, err := m.Check(host)
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
