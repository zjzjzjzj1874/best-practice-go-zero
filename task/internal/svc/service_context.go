package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/obs"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/middlewares"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mongo"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mysql"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	LogTrace        rest.Middleware
	ExportTaskModel mongo.ExportTaskModel
	MysqlDB         *gorm.DB // 本服务自己封装的一个mysql client
	RedisClient     helper.Client
	HwObsClient     *obs.HwObsClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		LogTrace:    middlewares.NewLogTraceMiddleware().Handle,
		MysqlDB:     mysql.MustNewDB(c.MysqlConf),
		RedisClient: helper.MustNewClient(c.Mode, c.CacheRedis),
		HwObsClient: obs.NewHWObsClient(c.HwObs, c.RestConf),
		//ExportTaskModel: mongo.NewExportTaskModel(c.MongoDB.URL, reflect.TypeOf(mongo.ExportTask{}).Name(), c.CacheRedis),
	}
}
