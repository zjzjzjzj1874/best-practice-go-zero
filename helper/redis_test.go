package helper

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"sync"
	"testing"
	"time"
)

var client Client

func init() {
	client = MustNewClient("dev", cache.CacheConf{cache.NodeConf{
		RedisConf: redis.RedisConf{
			Host: "127.0.0.1:6379",
			Pass: "123456a",
		}}})

}

func Test_Basic_redis(t *testing.T) {
	t.Run("#模拟并发任务", func(t *testing.T) {
		task := "1"
		wg := sync.WaitGroup{}
		wg.Add(4)
		for i := 0; i < 4; i++ {
			go func(idx int) {
				wg.Done()
				keepsake, lock := client.LockedDistributeTask(task) // 将任务1锁住
				t.Logf("Idx:%d,res:%v", idx, lock)
				if !lock {
					t.Logf("Idx:%d,锁住任务失败,等待下次尝试", idx)
					return
				}
				t.Logf("Idx:%d,成功获取事务锁.", idx)
				defer client.UnLockDistributeTask(task, keepsake)
			}(i)
		}
		wg.Wait()
		time.Sleep(time.Second)
		// 期望:四个协程,一个能够获取到锁,其他三个获取不到.
	})

	t.Run("SetNx", func(t *testing.T) {
		task := "1"
		keepsake, lock := client.LockedDistributeTask(task) // 将任务1锁住
		if !lock {
			t.Log("锁住任务失败,等待下次尝试")
			return
		}
		defer client.UnLockDistributeTask(task, keepsake)
		t.Log(lock)
	})
}
