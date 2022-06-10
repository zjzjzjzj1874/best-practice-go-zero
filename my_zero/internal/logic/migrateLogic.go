package logic

import (
	"context"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MigrateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMigrateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MigrateLogic {
	return &MigrateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Migrate 自动化生产表
func (l *MigrateLogic) Migrate(req *types.MigrateRequest) (resp *types.MigrateResponse, err error) {
	return nil, l.svcCtx.MysqlDB.MigrateWithApi([]string{})
	//return nil, l.svcCtx.MysqlDB.MigrateWithApi(req.TableNames)
}
