package customer

import (
	"context"
	"testing"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/wechat"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/config"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/types"
)

// 1.前端调用微信授权登陆 https://open.weixin.qq.com/connect/qrconnect?appid=appid&redirect_uri=https://www.baidu,com&response_type=code&scope=snsapi_login#wechat_redirect
func TestLoginWechatLogic_loginWechat(t *testing.T) {
	t.Run("#wechatLogin", func(t *testing.T) {
		l := &LoginWechatLogic{
			ctx: context.Background(),
			svcCtx: &svc.ServiceContext{
				Config: config.Config{
					WechatQR: wechat.LoginQRConf{
						AppId:       "appid",
						AppSecret:   "secret",
						RedirectURI: "www.baidu.com",
					},
				},
			},
		}
		wxUser, err := l.loginWechat(&types.LoginWechatReq{
			Code:               "0618C41w3htZK03hdF3w38Xfa618C41c",
			BrowserFingerprint: "",
		})
		if err != nil {
			panic(err)
		}

		t.Log(wxUser)
	})
}
