package test

import (
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/logic/v0/test"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := helper.Validate().StructCtx(r.Context(), &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := test.NewGetLogic(r.Context(), svcCtx)
		resp, err := l.TestGet(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
