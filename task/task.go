/*
 * @Author: zjzjzjzj1874 zjzjzjzj1874@gmail.com
 * @Date: 2023-12-12 11:27:59
 * @LastEditors: zjzjzjzj1874 zjzjzjzj1874@gmail.com
 * @LastEditTime: 2025-02-08 14:11:43
 * @FilePath: /best-practice-go-zero/task/task.go
 * @Description: task服务入口
 */
package main

import (
	_ "embed"
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/errors"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/config"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/cron"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/handler"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
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

	server := rest.MustNewServer(c.RestConf, Opt)
	defer server.Stop()

	httpx.SetErrorHandler(errors.ErrorRestfulHandler) // 自定义错误返回,restful类型
	//httpx.SetErrorHandler(errors.ErrorHandler)// 自定义错误返回,非restful类型,兼容我方前端
	handler.RegisterHandlers(server, ctx)
	server.PrintRoutes() // 打印路由

	helper.OpenPPROF(c.PprofConf)
	cron.InitCron(ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//go:embed task.json
var swagger []byte

// Opt 设置自定义跨域请求,或者返回的header
var Opt = rest.WithCustomCors(func(header http.Header) {
	header.Set("Access-Control-Allow-Headers", "*")
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Expose-Headers", "*")
	header.Set("Access-Control-Allow-Methods", "*")
	header.Set("Access-Control-Allow-Credentials", "false")
}, nil, "*")
