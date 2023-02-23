package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/config"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/cron"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/handler"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/task.yaml", "the config file")

func main() {
	flag.Parse()

	logx.DisableStat()
	helper.InitLogrus()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.Swagger = swagger

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	cron.InitCron(ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//go:embed swagger.json
var swagger []byte
