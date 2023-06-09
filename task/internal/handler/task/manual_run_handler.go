package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/logic/task"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/types"
)

func ManualRunHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ManualRunReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := task.NewManualRunLogic(r.Context(), svcCtx)
		resp, err := l.ManualRun(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
