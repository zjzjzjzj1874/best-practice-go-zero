package pay

import (
	"github.com/sirupsen/logrus"
	"net"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/logic/pay"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/types"
)

func PayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PayReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

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
		l := pay.NewPayLogic(r.Context(), svcCtx, host)
		resp, err := l.Pay(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
