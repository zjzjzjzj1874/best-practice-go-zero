package task

import (
	"context"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/cron"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManualRunLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManualRunLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManualRunLogic {
	return &ManualRunLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManualRunLogic) ManualRun(req *types.ManualRunReq) (resp *types.ManualRunResp, err error) {
	logx.WithContext(l.ctx).Infof("手动触发任务:%s", req.Name)
	msg := cron.RunWithName(l.ctx, l.svcCtx, req.Name)
	return &types.ManualRunResp{Msg: msg}, nil
}
