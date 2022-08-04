/**
 * @author zjzjzjzj
 * @description 限流器=> 一般而言,不同业务的限流不太一样(字段),所以可能不同的路由组需要限流的字段不太一样
 * @date 5:20 PM 2022/8/3
**/

package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	seconds = 10 // 限流周期(s)
	quota   = 5  // 周期内请求次数,允许=,但是不允许>
)

var (
	ErrorNoBusiness = errors.New("请传入业务ID")
)

type (
	FlowLimitMiddleware struct {
		host string // redis host
		pwd  string // redis pwd
	}

	Req struct {
		BusinessID string `json:"business_id,optional,omitempty" form:"business_id,optional"`
	}
)

func NewFlowLimitMiddleware(conf redis.RedisConf) *FlowLimitMiddleware {
	return &FlowLimitMiddleware{
		host: conf.Host,
		pwd:  conf.Pass,
	}
}

func (m *FlowLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Req
		if err := Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if req.BusinessID == "" {
			res, _ := json.Marshal(map[string]interface{}{
				"code": http.StatusBadRequest,
				"msg":  "请传入必要的业务",
			})
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(res)
			return
		}

		pass, err := m.Check(req.BusinessID)
		if err != nil || !pass {
			res, _ := json.Marshal(map[string]interface{}{
				"code": http.StatusTooManyRequests,
				"msg":  "当前业务请求次数超过限制",
			})
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write(res)
			return
		}

		next(w, r)
	}
}

func (m *FlowLimitMiddleware) Check(key string) (bool, error) {
	l := limit.NewPeriodLimit(seconds, quota, redis.New(m.host, redis.WithPass(m.pwd)), "periodLimit")
	code, err := l.Take(key)
	if err != nil {
		logrus.Errorf("take out key failure[key:%s,err:%s]", key, err.Error())
		return false, err
	}

	// switch val =&gt; process request
	switch code {
	case limit.OverQuota:
		logrus.Errorf("OverQuota key: %v", key)
		return false, err
	case limit.Allowed:
		logrus.Infof("AllowedQuota key: %v", key)
		return true, nil
	case limit.HitQuota: // 刚好达到限额,不过也是允许用户访问的
		logrus.Errorf("HitQuota key: %v", key)
		return true, err
	default:
		logrus.Errorf("DefaultQuota key: %v", key)
		// unknown response, we just let the sms go
		return true, nil
	}

}

func Parse(r *http.Request, v *Req) error {
	// 一般就只确认Form表单和body的参数,header和path需要的话可以再添加
	if err := httpx.ParseForm(r, v); err != nil {
		return err
	}
	if v.BusinessID != "" {
		return nil
	}

	bodyRes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if len(bodyRes) == 0 {
		return ErrorNoBusiness
	}
	r.Body = ioutil.NopCloser(bytes.NewReader(bodyRes))

	return json.Unmarshal(bodyRes, v)
}
