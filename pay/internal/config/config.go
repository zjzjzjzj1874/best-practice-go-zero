package config

import (
	"github.com/zeromicro/go-zero/rest"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/pay/alipay"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/pay/wechatpay"
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Swagger []byte `json:",optional"` // swagger文件

	WechatPay wechatpay.Conf // 微信支付
	AliPay    alipay.Conf    // 支付宝支付
}
