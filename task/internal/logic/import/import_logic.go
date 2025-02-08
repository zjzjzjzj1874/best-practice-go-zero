/*
 * @Author: zjzjzjzj1874 zjzjzjzj1874@gmail.com
 * @Date: 2023-12-12 11:27:59
 * @LastEditors: zjzjzjzj1874 zjzjzjzj1874@gmail.com
 * @LastEditTime: 2025-02-08 11:42:23
 * @FilePath: /best-practice-go-zero/task/internal/logic/import/import_logic.go
 * @Description: package name modify
 */
package ipt

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
