package swagger

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
)

type SwaggerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwaggerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwaggerLogic {
	return &SwaggerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwaggerLogic) Swagger() error {
	// todo: add your logic here and delete this line

	return nil
}
