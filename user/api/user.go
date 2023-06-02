package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/errors"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/config"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/handler"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	logx.DisableStat()
	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.Swagger = swagger

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(helper.UnAuthCallback), rest.WithCors())
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(errors.ErrorHandler) // 自定义错误

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//go:embed user.json
var swagger []byte

// 生成json文件: goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir .
