package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/rabbitmq"

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

	cron.Init(ctx) // 初始化轮询任务

	rabbitmq.InitProducer(context.TODO(), ctx.Config.RabbitMQ) // 初始化消息队列生产者
	rabbitmq.InitConsumer(context.TODO(), ctx.Config.RabbitMQ) // 初始化消息队列消费者
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
