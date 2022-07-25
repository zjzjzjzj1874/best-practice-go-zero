package test

import (
	"context"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) TestGet(req *types.ListReq) (resp *types.MongoTest, err error) {
	mt, err := l.svcCtx.MongoTestModel.FindOne(req.ID)
	if err != nil {
		return nil, err
	}
	resp = &types.MongoTest{}
	if err = copier.CopyWithOption(resp, mt, helper.OutOption()); err != nil {
		return nil, err
	}
	return
}
