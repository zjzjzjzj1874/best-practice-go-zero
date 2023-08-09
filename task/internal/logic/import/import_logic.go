package import

import (
	"context"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportLogic {
	return &ImportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportLogic) Import(req *types.ImportReq) (resp *types.ImportResp, err error) {
	// todo: add your logic here and delete this line

	return
}
