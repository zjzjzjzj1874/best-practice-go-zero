package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"

	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

const (
	TimeoutTaskQueuePrefix = "TimeoutTaskQueue.%s"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "123456a", // no password set
	DB:       0,         // use default DB
})

func TestTaskTimeoutCron_Run(t *testing.T) {
	// 往redis中写入数据
	t.Run("#hset", func(t *testing.T) {
		key := fmt.Sprintf(TimeoutTaskQueuePrefix, "dev")
		cacheTasks := make([]helper.CacheTaskQueueMetaData, 0)

		for i := 0; i < 20; i++ {
			userID := i
			for j := 0; j < 10; j++ {
				cacheTasks = append(cacheTasks, helper.CacheTaskQueueMetaData{
					TaskID:   strconv.Itoa(j),
					UserID:   strconv.Itoa(userID),
					ExpireAt: time.Now().Add(time.Second * time.Duration(rand.Intn(j+10))).Unix(),
				})
			}
			newStr, err := jsoniter.Marshal(cacheTasks)
			if err != nil {
				t.Fatal(err)
			}
			ok, err := rdb.HSet(key, fmt.Sprintf("%d", userID), newStr).Result()
			t.Log(ok, err)
			cacheTasks = make([]helper.CacheTaskQueueMetaData, 0)
		}
	})

	t.Run("#HGetAll", func(t *testing.T) {
		resMap, err := rdb.HGetAll("myhash").Result()
		if err != nil {
			t.Fatal(err)
		}

		for key, val := range resMap {
			cacheTasks := make([]helper.CacheTaskQueueMetaData, 0)
			fmt.Println("key: ", key)
			err = jsoniter.UnmarshalFromString(val, &cacheTasks)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(cacheTasks)
		}

		t.Log(len(resMap))
	})

	t.Run("#HScan handle with channel", func(t *testing.T) {
		var (
			key = fmt.Sprintf(TimeoutTaskQueuePrefix, "dev")
		)

		cacheTasksChan := make(chan []helper.CacheTaskQueueMetaData, 10)
		exit := make(chan struct{})
		// 起协程消费channel
		go func() {
			for {
				select {
				case <-exit:
					fmt.Println("收到退出信号,准备退出")
					return
				case cacheTasks := <-cacheTasksChan:
					if len(cacheTasks) == 0 {
						continue
					}

					userID := cacheTasks[0].UserID
					fmt.Println("pre消费", userID)
					time.Sleep(time.Millisecond * 500)
					fmt.Println("sleep 消费", userID)

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

					err = rdb.HSet(key, userID, newstr).Err()
					if err != nil {
						logrus.Errorf("HSet failure:[key:%s,userID:%s,err:%v]", key, userID, err.Error())
						continue
					}
				}
			}
		}()

		loadCacheTasks(key, cacheTasksChan, exit)
		time.Sleep(time.Second * 10)
	})

}

func loadCacheTasks(key string, cacheTasksChan chan []helper.CacheTaskQueueMetaData, exit chan struct{}) {
	var (
		cursor uint64
		resMap = make(map[string]string)
	)
	for {
		var (
			keys []string
			err  error
		)
		keys, cursor, err = rdb.HScan(key, cursor, "", 10).Result()
		if err != nil {
			fmt.Println(err)
		}
		if len(keys)%2 != 0 {
			fmt.Println("hscan不能成对出现")
		}

		for idx, val := range keys {
			if idx%2 == 0 {
				resMap[val] = keys[idx+1]
			}
		}

		for userID, val := range resMap {
			var (
				cacheTasks = make([]helper.CacheTaskQueueMetaData, 0)
			)
			err = jsoniter.UnmarshalFromString(val, &cacheTasks)
			if err != nil {
				logrus.Errorf("UnmarshalFromString failure:[userID:%s,val:%s,err:%s]", userID, val, err.Error())
				continue
			}

			cacheTasksChan <- cacheTasks
		}
		if cursor == 0 {
			time.Sleep(1 * time.Second)
			exit <- struct{}{}
			break
		}
	}
}

func TestTaskTimeoutCron(t *testing.T) {
	t.Run("#Cron", func(t *testing.T) {
		c := cron.New()
		err := c.AddFunc("0 0 5 31 2 *", func() {
			fmt.Println("cron trigger")
		})
		if err != nil {
			t.Errorf("err:%s", err.Error())
			return
		}
		_ = c.AddFunc("@every 10s", func() {
			fmt.Println("every 10s")
		})
		c.Start()
		fmt.Println("start cron")
		select {}
	})
}
