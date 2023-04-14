package pay

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/logic/pay"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/types"
)

func PayResultHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PayQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := pay.NewPayResultLogic(r.Context(), svcCtx)
		resp, err := l.PayResult(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
