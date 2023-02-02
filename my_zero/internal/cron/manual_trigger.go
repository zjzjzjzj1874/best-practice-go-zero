package cron

import (
	"github.com/sirupsen/logrus"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
)

// 手动触发测试
func manualTrigger(ctx *svc.ServiceContext) {
	logrus.Infof("pong")
}
