package test

import (
	"context"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GETAPITestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGETAPITestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GETAPITestLogic {
	return &GETAPITestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GETAPITestLogic) GETAPITest(req *types.ListReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
