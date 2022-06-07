package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/databases"
)

type Config struct {
	rest.RestConf

	CacheRedis cache.CacheConf

	MysqlConf databases.MysqlConfig

	Cron struct {
		TaskTimeoutSpec string
	}
}
