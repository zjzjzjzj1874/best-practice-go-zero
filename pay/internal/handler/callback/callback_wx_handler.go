package callback

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/logic/callback"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/types"
)

func CallbackWxHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WechatNotifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := callback.NewCallbackWxLogic(r.Context(), svcCtx)
		err := l.CallbackWx(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
