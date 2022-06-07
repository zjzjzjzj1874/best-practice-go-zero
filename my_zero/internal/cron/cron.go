package cron

import (
	"sync"

	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
)

var (
	cr            *cron.Cron
	cronLockInMem *sync.Map // cron任务在内存中的锁 =>即:当上一次任务没跑完,本次不需要执行
)

func InitTask(key, spec string) {
	cronLockInMem.Store(key, spec)
}

func LoadTask(key string) (value interface{}, ok bool) {
	return cronLockInMem.Load(key)
}

func ExecutedTask(key string) {
	cronLockInMem.Delete(key)
}

func Init(ctx *svc.ServiceContext) {
	cronLockInMem = &sync.Map{}
	cr = cron.New()

	if err := cr.AddJob(ctx.Config.Cron.TaskTimeoutSpec, &TaskTimeoutCron{
		ctx:  ctx,
		Spec: ctx.Config.Cron.TaskTimeoutSpec,
		Name: "超时订单清理任务",
	}); err != nil {
		panic(err)
	}

	cr.Start()
	logrus.Infof("cron task is running...")
}

func GetCronLock() *sync.Map {
	return cronLockInMem
}
