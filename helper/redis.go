package helper

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Client interface {
	Ping() (string, error)
	CronClearTimeoutTasks(userID string, expireTaskIDs []string)
	CronUpdateTimeoutTasks(userID, newstr string) error
	CronSweepTimeoutTaskQueueCacheWithScan(cacheTasksChan chan []CacheTaskQueueMetaData, exit chan struct{})
}

type defaultClient struct {
	env string // 环境名称 => (dev/test/pro/pre)
	*redis.Client
}

func MustNewClient(env string, conf cache.CacheConf) *defaultClient {
	if len(conf) == 0 {
		log.Fatal("no redis config")
	}
	client := redis.NewClient(&redis.Options{
		Addr:         conf[0].Host,
		Password:     conf[0].Pass,
		DB:           0,
		MaxRetries:   3,
		DialTimeout:  time.Second,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
		PoolSize:     10,
		MinIdleConns: 3,
	})
	dc := &defaultClient{
		env:    env,
		Client: client,
	}
	return dc
}

func (c defaultClient) Ping() (string, error) {
	return c.Client.Ping().Result()
}

// CronSweepTimeoutTaskQueueCacheWithScan 清除超时任务队列中的缓存
func (c defaultClient) CronSweepTimeoutTaskQueueCacheWithScan(cacheTasksChan chan []CacheTaskQueueMetaData, exit chan struct{}) {
	var (
		cursor uint64
		key    = fmt.Sprintf(TimeoutTaskQueuePrefix, c.env)
	)

	for {
		var (
			res    []string
			err    error
			resMap = make(map[string]string)
		)
		res, cursor, err = c.Client.HScan(key, cursor, "", 10).Result()
		if err != nil {
			logrus.Errorf("HScan failure:[key:%s,cursor:%d,err:%s]", key, cursor, err.Error())
			continue
		}

		if len(res)%2 != 0 {
			logrus.Errorf("HScan结果不能成对出现:[key.len:%d]", len(key))
			continue
		}

		for idx, val := range res {
			if idx%2 == 0 {
				resMap[val] = res[idx+1]
			}
		}

		for userID, val := range resMap {
			var (
				cacheTasks = make([]CacheTaskQueueMetaData, 0)
			)
			err = jsoniter.UnmarshalFromString(val, &cacheTasks)
			if err != nil {
				logrus.Errorf("UnmarshalFromString failure:[userID:%s,val:%s,err:%s]", userID, val, err.Error())
				continue
			}
			cacheTasksChan <- cacheTasks
		}

		// 游标扫描完毕,结束本次循环
		if cursor == 0 {
			break
		}
	}

	exit <- struct{}{} // 发送退出信号
}

// CronUpdateTimeoutTasks Cron自动更新用户过期任务
func (c defaultClient) CronUpdateTimeoutTasks(userID, newstr string) error {
	var key = fmt.Sprintf(TimeoutTaskQueuePrefix, c.env)
	return c.HSet(key, userID, newstr).Err()
}

// CronClearTimeoutTasks Cron自动清理过期任务
func (c defaultClient) CronClearTimeoutTasks(userID string, expireTaskIDs []string) {
	personKey := fmt.Sprintf(TaskPersonalPoolPrefix, c.env, userID)
	err := c.Client.HDel(personKey, expireTaskIDs...).Err() // 清理个人队列记录
	if err != nil {
		logrus.Errorf("CronClearTimeoutTasks failure:[key:%s,expireIDs:%v,err:%s]", personKey, expireTaskIDs, err.Error())
	}
}
