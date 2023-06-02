package swagger

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/svc"
)

func SwaggerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(svcCtx.Config.Swagger)
		httpx.Ok(w)
	}
}
