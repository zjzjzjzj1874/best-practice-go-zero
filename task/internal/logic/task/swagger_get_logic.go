package task

import (
	"context"
	"github.com/sirupsen/logrus"
	"io/ioutil"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
)

type SwaggerGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwaggerGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SwaggerGetLogic {
	return &SwaggerGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SwaggerGetLogic) SwaggerGet() (data []byte, err error) {
	logrus.Infof("读取swagger文件")
	return ioutil.ReadFile(l.svcCtx.Config.SwaggerPath)
}
