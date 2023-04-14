package wechatpay

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type WxH5PayClient struct {
	h5.H5ApiService
	Conf
}

func MustNewWxH5PayClient(ctx context.Context, c Conf) *WxH5PayClient {
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(c.MchPrivateKeyPath)
	if err != nil {
		panic(fmt.Sprintf("【WXPay】LoadPrivateKeyWithPath failure:%s", err.Error()))
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(c.MchID, c.MchCertificateSerialNumber, mchPrivateKey, c.MchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		panic(fmt.Sprintf("【WXPay】NewClient err:%s", err))
	}

	return &WxH5PayClient{
		H5ApiService: h5.H5ApiService{Client: client},
		Conf:         c,
	}
}

// 微信证书文档:https://pay.weixin.qq.com/wiki/doc/apiv3/open/pay/wechatpay5_0.shtml
// 微信证书生成证书序列号文档:https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay7_0.shtml ==> openssl x509 -in apiclient_cert.pem -noout -serial
// 微信支付申请文档:https://pay.weixin.qq.com/wiki/doc/apiv3/open/pay/chapter2_5_1.shtml#part-4
