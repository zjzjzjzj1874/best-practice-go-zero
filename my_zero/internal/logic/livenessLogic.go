package logic

import (
	"context"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"

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

func (l *LivenessLogic) Liveness(req *types.LivenessRequest) (resp *types.LivenessResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
