package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mysql"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/rabbitmq"
)

type Config struct {
	rest.RestConf
	helper.PprofConf

	CacheRedis cache.CacheConf

	MysqlConf mysql.MysqlConfig
	MongoDB   struct {
		URL string // MongoDB数据库链接url
	}

	Cron struct {
		TaskTimeoutSpec string
	}
	RabbitMQ rabbitmq.Config
}
