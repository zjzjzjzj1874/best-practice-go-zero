package pro

import (
	"context"
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

// RabbitMQProducer TODO 这里结构体可以使用copier来优化
func (l *RabbitMQProducerLogic) RabbitMQProducer(req *types.RabbitmqProRequest) (resp *types.RabbitmqProResponse, err error) {
	rabbitmq.Producer <- rabbitmq.PublishMetaData{
		Name:    req.Name,
		Age:     req.Age,
		Hobbies: req.Hobbies,
	}

	return
}
