package logic

import (
	"context"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/rpc/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoResp, error) {
	// TODO do your self logic
	return &user.UserInfoResp{
		Id:     1,
		Name:   "Test",
		Email:  "test@gmail.com",
		Phone:  "18398340843",
		Gender: "1",
	}, nil
}
