package middlewares

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type LogMiddleware struct {
	systemName string // 系统名
}

func NewLogMiddleware(systemName string) *LogMiddleware {
	return &LogMiddleware{systemName: systemName}
}

func (m *LogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet { // get请求不做任何操作
			next(w, r)
			return
		}

		nb, _ := r.GetBody() // 这里读取之后不需要再放回去
		body := make([]byte, 0)
		_, _ = nb.Read(body)
		logrus.Println("收到 body = ", string(body))

		//bodyRes, _ := ioutil.ReadAll(r.Body)
		//if len(bodyRes) != 0 {
		//	r.Body = ioutil.NopCloser(bytes.NewReader(bodyRes)) // 读取之后给放回去
		//}
		logrus.Infof("中间件处理日志")
		next(w, r)
	}
}
