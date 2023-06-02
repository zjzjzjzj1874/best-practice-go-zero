package customer

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/logic/customer"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/types"
)

func LoginWechatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginWechatReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := customer.NewLoginWechatLogic(r.Context(), svcCtx)
		resp, err := l.LoginWechat(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
