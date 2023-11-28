package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		MaxRetries:   3,
		DialTimeout:  time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
		MaxConnAge:   time.Hour,
		PoolTimeout:  time.Minute,
	})
}

var (
	Str  = "Test_%d_%d"
	ZSet = "ZSet_Test"
)

func Test_Incr(t *testing.T) {
	client.IncrBy(context.Background(), fmt.Sprintf(Str, 1, 1), 38)
}

func Test_Zadd(t *testing.T) {
	client.ZAdd(context.Background(), ZSet, &redis.Z{
		Score:  11,
		Member: "11",
	})
	client.ZAdd(context.Background(), ZSet, &redis.Z{
		Score:  4,
		Member: "44",
	})
	client.ZAdd(context.Background(), ZSet, &redis.Z{
		Score:  1,
		Member: "1",
	})
	client.ZAdd(context.Background(), ZSet, &redis.Z{
		Score:  4,
		Member: "14",
	})
	client.ZAdd(context.Background(), ZSet, &redis.Z{
		Score:  321,
		Member: "13",
	})
	client.ZAdd(context.Background(), ZSet, &redis.Z{
		Score:  0,
		Member: "2",
	})
	client.ZAdd(context.Background(), ZSet, &redis.Z{
		Score:  0,
		Member: "3",
	})
	client.ZAdd(context.Background(), ZSet, &redis.Z{
		Score:  0,
		Member: "6",
	})
	client.ZAdd(context.Background(), ZSet, &redis.Z{
		Score:  9,
		Member: "9",
	})
	client.ZAdd(context.Background(), ZSet, &redis.Z{
		Score:  0,
		Member: "10",
	})
}
func Test_ZScore(t *testing.T) {
	res, err := client.ZScore(context.Background(), ZSet, "14").Result()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("res:%v\n", res)
}

func Test_ZRang(t *testing.T) {
	res, err := client.ZRangeByScoreWithScores(context.Background(), ZSet, &redis.ZRangeBy{
		Min:    "1",
		Max:    "99999999",
		Offset: 0,
		Count:  0,
	}).Result()
	if err != nil {
		t.Fatal(err)
	}

	for _, re := range res {
		t.Logf("res: member:%v,score:%v\n", re.Member, re.Score)
	}
}
