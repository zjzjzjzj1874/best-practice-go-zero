package task

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/cron"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManualRunTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManualRunTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManualRunTaskLogic {
	return &ManualRunTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManualRunTaskLogic) ManualRunTask(req *types.ManualExecTaskRequest) (resp *types.ManualExecTaskResponse, err error) {
	logrus.Infof("手动触发任务入口")
	logx.WithContext(l.ctx).Infof("手动触发任务入口")
	msg := cron.RunWithName(l.ctx, l.svcCtx, req.Name)
	return &types.ManualExecTaskResponse{Msg: msg}, nil
}
