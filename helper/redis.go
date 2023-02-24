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

	LockedDistributeTask(task string) (keepsake interface{}, ok bool) // 给redis任务上分布式事务锁;OK:是否加锁成功;keepsake:上锁的值,释放的时候只能释放这个值;
	UnLockDistributeTask(task string, keepsake interface{})           // 释放事务锁:keepsake:只能释放keepsake的锁.防止错误释放其他用户加的锁
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

// LockedDistributeTask 是否被分布式事务锁锁住
func (c defaultClient) LockedDistributeTask(task string) (interface{}, bool) {
	key := fmt.Sprintf(CronTaskLockerPrefix, c.env, task)
	keepsake := time.Now().UnixMilli()
	// 分布式事务锁,锁住3秒即可,原因:
	// 1.一般任务最低每5s触发一次,分布式事务锁住3s不影响下一次执行;
	// 2.服务器如果同步NTP服务器频率低,那么时差就会与晶振有关,误差应该也在3s以内,误差过大的情况咱们也没办法控制了.
	// TODO 待确认:是锁住3s还是锁住到下一次的时间?因为有可能任务确实在本周期没有跑完,所以这个其实是有待商榷的.
	return keepsake, c.Client.SetNX(key, 1, 3*time.Second).Val()
}

func (c defaultClient) UnLockDistributeTask(task string, keepsake interface{}) {
	key := fmt.Sprintf(CronTaskLockerPrefix, c.env, task)

	// Note:只能删除自己的key 使用lua来操作,属于自己的key才释放
	err := redis.NewScript(checkKeyValueBeforeDelKeyScript).Run(c.Client, []string{key}, keepsake).Err()
	if err != nil && err != redis.Nil {
		logrus.Errorf("Del Key failure:[Key:%s;err:%s]", key, err.Error())
	}
}

// 删除Key之前先校验key值
var checkKeyValueBeforeDelKeyScript = `
  local key = KEYS[1]
  local targetValue = ARGV[1]

  if redis.call("GET", key) == targetValue then
    redis.call("DEL", key)
    return true
  else
    return false
  end
`
