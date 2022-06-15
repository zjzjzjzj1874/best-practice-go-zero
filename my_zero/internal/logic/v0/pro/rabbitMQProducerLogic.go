package pro

import (
	"context"
	"fmt"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/rabbitmq"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RabbitMQProducerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRabbitMQProducerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RabbitMQProducerLogic {
	return &RabbitMQProducerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RabbitMQProducer rabbitmq生产数据
func (l *RabbitMQProducerLogic) RabbitMQProducer(req *types.RabbitmqProRequest) (resp *types.RabbitmqProResponse, err error) {
	rabbitmq.ProduceData(rabbitmq.PublishMetaData{
		Name:    req.Name,
		Age:     req.Age,
		Hobbies: req.Hobbies,
	})

	fmt.Println("当前任务长度:", rabbitmq.ProducerLen())

	return
}
