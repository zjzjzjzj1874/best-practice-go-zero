package task

import (
	"context"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LivenessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLivenessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LivenessLogic {
	return &LivenessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LivenessLogic) Liveness() (resp *types.LivenessResponse, err error) {
	return &types.LivenessResponse{Msg: "ok"}, nil
	//return &types.LivenessResponse{Msg: "ok"}, errors.New("mock err")
}
