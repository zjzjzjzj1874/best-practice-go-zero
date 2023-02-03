package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/middlewares"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	LogTrace rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		LogTrace: middlewares.NewLogTraceMiddleware().Handle,
	}
}
