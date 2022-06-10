package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/databases"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/rabbitmq"
)

type Config struct {
	rest.RestConf
	helper.PprofConf

	CacheRedis cache.CacheConf

	MysqlConf databases.MysqlConfig

	Cron struct {
		TaskTimeoutSpec string
	}
	RabbitMQ rabbitmq.Config
}
