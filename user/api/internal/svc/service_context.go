package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/config"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/middleware"
)

type ServiceContext struct {
	Config    config.Config
	Log       rest.Middleware
	FlowLimit rest.Middleware
	Recover   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Log:       middleware.NewLogMiddleware().Handle,
		FlowLimit: middleware.NewFlowLimitMiddleware().Handle,
		Recover:   middleware.NewRecoverMiddleware().Handle,
	}
}
