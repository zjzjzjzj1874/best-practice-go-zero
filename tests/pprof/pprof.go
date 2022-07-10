package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Hour)

	go func() {
		con := true
		for con {
			select {
			case <-ctx.Done():
				fmt.Println("完成任务")
				con = false
			default:
				fmt.Println("任务进行时")
				time.Sleep(time.Millisecond * 100)
			}
		}
		fmt.Println("任务完成,准备退出")
	}()

	go func() {
		for {
			log.Println("hello pprof")
			time.Sleep(time.Second)
			cancel()
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:8899", nil)
}
