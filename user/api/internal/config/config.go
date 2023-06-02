package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/wechat"
)

type Config struct {
	rest.RestConf
	Swagger  []byte             `json:",optional"` // swagger文件
	WechatQR wechat.LoginQRConf // 微信扫码登陆配置
}
