package cron

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"

	"time"
)

type TaskTimeoutCron struct {
	ctx  *svc.ServiceContext // 上下文信息
	Name string              // 任务名称
	Spec string              // Cron表达式
}

func (c *TaskTimeoutCron) Run() {
	// 上次任务是否还在执行 ? 打日志不执行 : 再次执行任务
	if _, ok := LoadTask(c.Name); ok {
		logrus.Warnf("last %s is running,this time do not need run.", c.Name)
		return
	}

	InitTask(c.Name, c.Spec)
	logrus.Infof("Cron %s begin running...", c.Name)

	cacheTasksChan := make(chan []helper.CacheTaskQueueMetaData, 10)
	exit := make(chan struct{})
	// 起协程消费channel
	go func() {
		defer ExecutedTask(c.Name) // 任务完成后清理lock
		for {
			select {
			case cacheTasks, ok := <-cacheTasksChan:
				if !ok {
					logrus.Infof("no more data:ready to exit")
					return
				}
				if len(cacheTasks) == 0 {
					continue
				}
				userID := cacheTasks[0].UserID
				newTasks := make([]helper.CacheTaskQueueMetaData, 0, len(cacheTasks))
				expireTaskIDs := make([]string, 0, len(cacheTasks))
				for _, task := range cacheTasks {
					// 当前时间小于任务过期时间,表示还没有过期,不需要清理
					if time.Now().Unix() < task.ExpireAt {
						// 未过期
						newTasks = append(newTasks, task)
					} else {
						// 过期
						expireTaskIDs = append(expireTaskIDs, task.TaskID)
					}
				}
				logrus.Infof("UserID:%s,cacheTasks.len:%d,newTasks.len:%d,expireTaskIDs:%v", userID, len(cacheTasks), len(newTasks), expireTaskIDs)

				newstr, err := jsoniter.MarshalToString(newTasks)
				if err != nil {
					logrus.Errorf("jsoniter.MarshalToString failure:[newTasks:%+v,err:%v]", newTasks, err.Error())
					continue
				}

				// 重置缓存池中用户缓存
				err = c.ctx.RedisClient.CronUpdateTimeoutTasks(userID, newstr)
				if err != nil {
					logrus.Errorf("CronUpdateTimeoutTasks failure:[userID:%s,err:%v]", userID, err.Error())
					continue
				}
				if len(expireTaskIDs) == 0 {
					continue
				}

				// 删除用户个人缓存
				c.ctx.RedisClient.CronClearTimeoutTasks(userID, expireTaskIDs)
			case <-exit:
				logrus.Infof("receive exit singal,ready to exit")
				return
			}
		}
	}()
	// 自动任务清理缓存队列
	c.ctx.RedisClient.CronSweepTimeoutTaskQueueCacheWithScan(cacheTasksChan, exit)
}
