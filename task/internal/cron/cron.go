// Package cron 本包用来跑定时任务
package cron

import (
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/exporter"
	"sync"
	"time"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

var (
	cr      *cron.Cron
	taskMap sync.Map
)

type task struct {
	name string                // 任务名称
	spec string                // cron执行表达式
	do   func(ctx interface{}) // 真正任务执行函数
	//do   func(ctx *svc.ServiceContext) // 真正任务执行函数
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
			name: "手动测试任务",
			spec: ctx.Config.Cron.TaskTestSpec,
			do:   manualTrigger,
		},
		task{
			name: "导出任务",
			spec: ctx.Config.Cron.TaskExportSpec,
			do:   exporter.Export,
		},
	}

	for i := range tasks {
		task := tasks[i]
		set(task.name, &tasks[i]) // 将任务注册到全局map
		err := cr.AddFunc(task.spec, func() {
			if InitTask(task.name) { // cron任务增加redis分布式事务锁 => 这样这个服务就不仅仅只能部署单节点
				keepsake, ok := ctx.RedisClient.LockedDistributeTask(task.name)
				if !ok {
					logrus.Warnf("没有获取到分布式事务锁:%s start running on %s other Node", task.name, time.Now().Format(time.RFC3339))
					return
				}
				defer ctx.RedisClient.UnLockDistributeTask(task.name, keepsake)
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
