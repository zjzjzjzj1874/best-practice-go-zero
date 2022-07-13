package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/order/api/internal/config"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
