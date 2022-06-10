package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			log.Println("hello pprof")
			time.Sleep(time.Minute)
		}
	}()

	_ = http.ListenAndServe("0.0.0.0:8899", nil)
}
