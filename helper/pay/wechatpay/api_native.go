package wechatpay

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

type WxNativePayClient struct {
	native.NativeApiService
	Conf
}

func MustNewWxNativePayClient(ctx context.Context, c Conf) *WxNativePayClient {
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

	return &WxNativePayClient{
		NativeApiService: native.NativeApiService{Client: client},
		Conf:             c,
	}
}

// DecryptAES256GCM 使用 AEAD_AES_256_GCM 算法进行解密
//
// 你可以使用此算法完成微信支付平台证书和回调报文解密，详见：
// https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/zheng-shu-he-hui-tiao-bao-wen-jie-mi
func DecryptAES256GCM(aesKey, associatedData, nonce, ciphertext string) (plaintext []byte, err error) {
	plaintext = make([]byte, 0)
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return
	}
	c, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return
	}

	return gcm.Open(nil, []byte(nonce), decodedCiphertext, []byte(associatedData))
}

// 微信证书文档:https://pay.weixin.qq.com/wiki/doc/apiv3/open/pay/wechatpay5_0.shtml
// 微信证书生成证书序列号文档:https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay7_0.shtml ==> openssl x509 -in apiclient_cert.pem -noout -serial
// 微信支付申请文档:https://pay.weixin.qq.com/wiki/doc/apiv3/open/pay/chapter2_5_1.shtml#part-4
