package pro

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/logic/v0/pro"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"
)

func RabbitMQProducerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RabbitmqProRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := pro.NewRabbitMQProducerLogic(r.Context(), svcCtx)
		resp, err := l.RabbitMQProducer(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
