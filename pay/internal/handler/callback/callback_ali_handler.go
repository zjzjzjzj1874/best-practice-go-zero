package callback

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/logic/callback"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
)

func CallbackAliHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := callback.NewCallbackAliLogic(r.Context(), svcCtx)
		err := l.CallbackAli(r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
