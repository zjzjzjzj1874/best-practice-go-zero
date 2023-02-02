package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/appengine/log"
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

	logx.DisableStat()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	helper.OpenPPROF(c.PprofConf)
	cron.InitCron(ctx) // 初始化定时任务

	handler.RegisterHandlers(server, ctx)

	//bc, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//async(bc)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func async(ctx context.Context) {
	for i := 0; i < 3; i++ {
		go func(idx int) {
			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

			for {
				time.Sleep(time.Second)
				fmt.Println(time.Now())
				select {
				case <-sigs:
					fmt.Println("notify sigs,bye:", idx)
					fmt.Println("http shutdown")
					return
				case <-ctx.Done():
					log.Infof(ctx, "ctx.Done() ")
					return
				default:
				}
			}
		}(i)
	}

}
