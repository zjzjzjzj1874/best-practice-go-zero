package svc

import "github.com/zjzjzjzj1874/best-pracrice-go-zero/user/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
