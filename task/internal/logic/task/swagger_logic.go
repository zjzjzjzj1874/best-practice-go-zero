package task

import (
	"context"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *SwaggerLogic) Swagger(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
