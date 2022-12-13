package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/config"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/cron"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/handler"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/my-zero.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// todo rabbitmq学习pprof的开启和zero中prometheus的开启,将从这里取消哦,以后添加一个中间件初始化的方法,初始化pprof,mq,cron的任务等等
	helper.OpenPPROF(c.PprofConf)
	cron.Init(ctx) // 初始化轮询任务
	//rabbitmq.InitProducer(context.TODO(), ctx.Config.RabbitMQ) // 初始化消息队列生产者
	//rabbitmq.InitConsumer(context.TODO(), ctx.Config.RabbitMQ) // 初始化消息队列消费者

	handler.RegisterHandlers(server, ctx)
	//async()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func async() {
	for i := 0; i < 3; i++ {
		go func(idx int) {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

			for {
				time.Sleep(time.Second)
				fmt.Println("hello world")

				select {
				case <-sigs:
					fmt.Println("notify sigs,bye:", idx)
					fmt.Println("http shutdown")
					return
				default:

				}
			}
		}(i)
	}

}
