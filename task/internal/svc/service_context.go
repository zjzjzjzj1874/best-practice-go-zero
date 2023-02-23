package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/obs"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/middlewares"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mongo"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	LogTrace        rest.Middleware
	ExportTaskModel mongo.ExportTaskModel
	HwObsClient     *obs.HwObsClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		LogTrace: middlewares.NewLogTraceMiddleware().Handle,
		//HwObsClient: obs.NewHWObsClient(c.HwObs, c.RestConf),
		//ExportTaskModel: mongo.NewExportTaskModel(c.MongoDB.URL, reflect.TypeOf(mongo.ExportTask{}).Name(), c.CacheRedis),
	}
}
