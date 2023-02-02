package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/obs"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mongo"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/config"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/db"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/middleware"
)

type ServiceContext struct {
	Config          config.Config
	RedisClient     helper.Client
	MysqlDB         *db.Mysql // 本服务自己封装的一个mysql client
	MongoTestModel  mongo.MongoTestModel
	FlowLimit       rest.Middleware
	LogTrace        rest.Middleware
	ExportTaskModel mongo.ExportTaskModel
	HwObsClient     *obs.HwObsClient
	//GormDB      *gorm.DB// database中的gorm.DB客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//RedisClient:     helper.MustNewClient(c.Mode, c.CacheRedis),
		//MysqlDB:         db.NewMysqlClient(mysql.MustNewDB(c.MysqlConf)),
		//HwObsClient:     obs.NewHWObsClient(c.HwObs, c.RestConf),
		//ExportTaskModel: mongo.NewExportTaskModel(c.MongoDB.URL, reflect.TypeOf(mongo.ExportTask{}).Name(), c.CacheRedis),
		////MongoTestModel: mongo.NewMongoTestModel(c.MongoDB.URL, reflect.TypeOf(mongo.Test{}).Name(), c.CacheRedis),
		//FlowLimit: middleware.NewFlowLimitMiddleware(c.CacheRedis[0].RedisConf).Handle,
		LogTrace: middleware.NewLogTraceMiddleware().Handle,
	}
}
