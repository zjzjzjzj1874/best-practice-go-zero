package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/logic/task"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
)

func SwaggerGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := task.NewSwaggerGetLogic(r.Context(), svcCtx)
		resp, err := l.SwaggerGet()
		if err != nil {
			httpx.Error(w, err)
		} else {
			_, _ = w.Write(resp)
			httpx.Ok(w)
		}
	}
}
