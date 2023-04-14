package liveness

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/logic/liveness"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
)

func LivenessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := liveness.NewLivenessLogic(r.Context(), svcCtx)
		err := l.Liveness()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
