package svc

import (
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/databases"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient helper.Client
	MysqlDB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: helper.MustNewClient(c.Mode, c.CacheRedis),
		MysqlDB:     databases.MustNewDB(c.MysqlConf),
	}
}
