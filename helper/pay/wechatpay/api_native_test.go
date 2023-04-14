package wechatpay

import (
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
	"testing"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
)

var nativeConf = Conf{
	IsProd:                     false,
	NotifyUrl:                  "https://www.baidu.com/pay/callback/wechat",
	MchPrivateKeyPath:          "./apiclient_key.pem",
	AppID:                      "wx1234567abcdefg2b",
	MchID:                      "1492583690",
	MchCertificateSerialNumber: "123456789ABCDEFGB71AF448684A5E00A60505",
	MchAPIv3Key:                "wwwBAIDU9N6Gbaidubd123456bd",
}

func TestMustNewWxNativePayClient(t *testing.T) {
	t.Run("#WXClient", func(t *testing.T) {
		ctx := context.Background()
		client := MustNewWxNativePayClient(ctx, nativeConf)

		request := native.PrepayRequest{
			Appid:       core.String(nativeConf.AppID),
			Mchid:       core.String(nativeConf.MchID),
			Description: core.String("Image形象店-深圳腾大-QQ公仔"),
			OutTradeNo:  core.String("1217752501201407033233368018"),
			TimeExpire:  core.Time(time.Now()),
			Attach:      core.String("自定义数据说明"),
			NotifyUrl:   core.String(conf.NotifyUrl), // 微信回调地址
			GoodsTag:    core.String("WXG"),
			//LimitPay:    []string{"LimitPay_example"},
			SupportFapiao: core.Bool(false),
			Amount: &native.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(1),
			},
			Detail: &native.Detail{
				GoodsDetail: []native.GoodsDetail{native.GoodsDetail{
					GoodsName:        core.String("iPhoneX 256G"),
					MerchantGoodsId:  core.String("ABC"),
					Quantity:         core.Int64(1),
					UnitPrice:        core.Int64(1), // 订单待完善
					WechatpayGoodsId: core.String("1001"),
				}},
			},
			SceneInfo: &native.SceneInfo{
				PayerClientIp: core.String("1.170.210.98"), // 必须
				DeviceId:      core.String("商户端设备号"),
				StoreInfo: &native.StoreInfo{
					Id:   core.String("商户侧门店编号"),
					Name: core.String("商户侧门店名称"),
				},
			},
		}
		resp, result, err := client.NativeApiService.Prepay(ctx, request)
		if err != nil {
			panic(err)
		}
		t.Log(resp)
		t.Log(result.Response)
	})

}
