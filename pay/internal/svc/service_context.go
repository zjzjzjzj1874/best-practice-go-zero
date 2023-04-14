package svc

import (
	"context"
	
	"github.com/zeromicro/go-zero/rest"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/middlewares"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/pay/alipay"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/pay/wechatpay"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Log    rest.Middleware

	*wechatpay.WxNativePayClient
	*alipay.AliPayClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Log:    middlewares.NewLogMiddleware(c.Name).Handle,

		WxNativePayClient: wechatpay.MustNewWxNativePayClient(context.Background(), c.WechatPay),
		AliPayClient:      alipay.MustNewAliPayClient(c.AliPay),
	}
}
