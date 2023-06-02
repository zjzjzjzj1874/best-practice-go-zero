package customer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/wechat"
	"io"
	"net/http"
	"net/url"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginWechatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginWechatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginWechatLogic {
	return &LoginWechatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginWechatLogic) LoginWechat(req *types.LoginWechatReq) (resp *types.LoginWechatResp, err error) {
	// todo: add your logic here and delete this line

	return
}

// 微信扫码登陆
func (l *LoginWechatLogic) loginWechat(req *types.LoginWechatReq) (resp WebWXUserInfo, err error) {
	webWXUserInfo := WebWXUserInfo{}
	// 通过code获取access_token
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("code", req.Code)
	params.Add("appid", l.svcCtx.Config.WechatQR.AppId)
	params.Add("secret", l.svcCtx.Config.WechatQR.AppSecret)
	params.Add("redirect_uri", l.svcCtx.Config.WechatQR.RedirectURI)
	loginURL := fmt.Sprintf("%s?%s", wechat.LoginQRUrl, params.Encode())
	response, err := http.Get(loginURL)
	if err != nil {
		logx.Errorf("Get Access_token from WX failure:[Code:%s,err:%s]", req.Code, err.Error())
		return webWXUserInfo, err
	}

	defer response.Body.Close()
	bs, _ := io.ReadAll(response.Body)
	// 成功
	// {
	// "access_token":"ACCESS_TOKEN",
	// "expires_in":7200,
	// "refresh_token":"REFRESH_TOKEN",
	// "openid":"OPENID",
	// "scope":"SCOPE",
	// "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
	// }
	// 失败 {"errcode":40001, "errmsg":"***" }
	webWXToken := &WebWXToken{}
	err = json.Unmarshal(bs, webWXToken)
	if err != nil {
		logx.Errorf("Unmarshal from WebWXToken failure:[bs:%s,err:%s]", string(bs), err.Error())
		return webWXUserInfo, err
	}

	if webWXToken.Errcode > 0 {
		logx.Errorf("webWXToken.Errcode failure:[webWXToken:%+v,err:%s]", webWXToken, err.Error())
		return webWXUserInfo, errors.New("登录失败 Error_100006")
	}

	// 通过access_token调用接口 获取用户个人信息
	params2 := url.Values{}
	params2.Add("access_token", webWXToken.AccessToken)
	params2.Add("openid", webWXToken.Openid)
	loginURL = fmt.Sprintf("%s?%s", wechat.UserInfoUrl, params2.Encode())
	response2, err := http.Get(loginURL)
	if err != nil {
		logx.Errorf("Get UserInfo from WX failure:[Code:%s,err:%s]", req.Code, err.Error())
		return webWXUserInfo, err
	}
	defer response2.Body.Close()
	bs2, _ := io.ReadAll(response2.Body)
	// 成功 {
	// "openid":"***",
	// "nickname":"***",
	// "sex":0,
	// "headimgurl":"http...",
	// "unionid":"******"
	//}
	// 失败 {"errcode":40001, "errmsg":"***" }
	logx.Infof("UserInfo Body:%s", string(bs2))
	err = json.Unmarshal(bs2, &webWXUserInfo)
	if err != nil {
		logx.Errorf("Unmarshal from webWXUserInfo failure:[bs:%s,err:%s]", string(bs2), err.Error())
		return webWXUserInfo, err
	}
	if webWXUserInfo.Errcode > 0 {
		logx.Errorf("webWXUserInfo.Errcode failure:[webWXToken:%+v,err:%s]", webWXUserInfo, err.Error())
		return webWXUserInfo, errors.New("登录失败 Error_100008")
	}

	return webWXUserInfo, nil
}

// region 微信登陆

type WebWXLogin struct {
	AppId       string
	AppSecret   string
	RedirectURI string
	StatePrefix string
}

type WebWXToken struct {
	AccessToken string `json:"access_token"`
	Openid      string `json:"openid"`
	Unionid     string `json:"unionid"`
	Errcode     int    `json:"errcode"`
}

type WebWXUserInfo struct {
	HeadImgUrl string `json:"headimgurl"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	Nickname   string `json:"nickname"`
	Errcode    int    `json:"errcode"`
}

// endregion 微信登陆
