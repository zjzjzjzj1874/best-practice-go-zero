package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello panic")

	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recover failure:[err:%s]", err)
			}
		}()
		fmt.Println("hello panic in goroutine")
		time.Sleep(time.Second)
		panic("崩溃测试")
	}()

	for i := 0; i < 50; i++ {
		fmt.Println("hello:", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Int31n(100)))
	}

	time.Sleep(time.Hour)
}
