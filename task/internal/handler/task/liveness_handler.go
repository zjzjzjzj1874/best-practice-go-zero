package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/logic/task"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
)

func LivenessHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := task.NewLivenessLogic(r.Context(), svcCtx)
		resp, err := l.Liveness()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
