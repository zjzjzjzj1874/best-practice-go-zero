// Package cron 本包用来跑定时任务
package cron

import (
	"sync"
	"time"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
)

var (
	cr      *cron.Cron
	taskMap sync.Map
)

type task struct {
	name string                        // 任务名称
	spec string                        // cron执行表达式
	do   func(ctx *svc.ServiceContext) // 真正任务执行函数
}

type tasks []task

func InitTask(key string) bool {
	if v, ok := taskMap.Load(key); ok {
		if v == true {
			return false
		}
	}

	taskMap.Store(key, true)
	return true
}

func ExecutedTask(key string) {
	taskMap.Store(key, false)
}

func InitCron(ctx *svc.ServiceContext) {
	logrus.Infof("任务初始化...")
	cr = cron.New()
	tasks := tasks{
		task{
			name: "超时任务",
			spec: ctx.Config.Cron.TaskTimeoutSpec,
			do:   taskTimeout,
		},
	}

	for i := range tasks {
		task := tasks[i]
		err := cr.AddFunc(task.spec, func() {
			if InitTask(task.name) {
				logrus.Debugf("%s start running on %s", task.name, time.Now().Format(time.RFC3339))
				task.do(ctx)
				ExecutedTask(task.name)
			} else {
				logrus.Infof("%s is running", task.name)
			}
		})
		if err != nil {
			logrus.Fatalf("start ask fail;[name:%s,spec:%s,err:%s]", task.name, task.spec, err)
		}
	}

	cr.Start()
}
