package wechat

// 微信扫码登陆文档:https://developers.weixin.qq.com/doc/oplatform/Website_App/WeChat_Login/Wechat_Login.html

const (
	LoginQRUrl  = "https://api.weixin.qq.com/sns/oauth2/access_token"
	UserInfoUrl = "https://api.weixin.qq.com/sns/userinfo"
)

type LoginQRConf struct {
	AppId       string // 微信扫码登陆 =>注册应用的AppId
	AppSecret   string // 微信扫码登陆 =>注册应用的密钥
	RedirectURI string // 微信扫码登陆 =>登录后回调地址
}
