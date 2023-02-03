package cron

import (
	"github.com/sirupsen/logrus"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
)

// 手动触发测试
//func manualTrigger(ctx *svc.ServiceContext) {
func manualTrigger(iCtx interface{}) {
	if _, ok := iCtx.(*svc.ServiceContext); !ok {
		logrus.Errorf("ctx type error:%T", iCtx)
		return
	}
	logrus.Infof("pong")
}
