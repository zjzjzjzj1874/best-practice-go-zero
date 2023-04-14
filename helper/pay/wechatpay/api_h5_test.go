package wechatpay

import (
	"context"
	"testing"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/h5"
)

var conf = Conf{
	IsProd:                     false,
	NotifyUrl:                  "https://www.baidu.com/pay/callback/wechat",
	MchPrivateKeyPath:          "./apiclient_key.pem",
	AppID:                      "wx1234567abcdefg2b",
	MchID:                      "1492583690",
	MchCertificateSerialNumber: "123456789ABCDEFGB71AF448684A5E00A60505",
	MchAPIv3Key:                "wwwBAIDU9N6Gbaidubd123456bd",
}

func TestMustNewWxH5PayClient(t *testing.T) {
	t.Run("#WXClient", func(t *testing.T) {
		ctx := context.Background()
		client := MustNewWxH5PayClient(ctx, conf)

		request := h5.PrepayRequest{
			Appid:         core.String(conf.AppID),
			Mchid:         core.String(conf.MchID),
			Description:   core.String("Image形象店-深圳腾大-QQ公仔"),
			OutTradeNo:    core.String("1217752501201407033233368018"),
			TimeExpire:    core.Time(time.Now()),
			Attach:        core.String("自定义数据说明"),
			NotifyUrl:     core.String(conf.NotifyUrl), // 微信回调地址
			GoodsTag:      core.String("WXG"),
			LimitPay:      []string{"LimitPay_example"},
			SupportFapiao: core.Bool(false),
			Amount: &h5.Amount{
				Currency: core.String("CNY"),
				Total:    core.Int64(1),
			},
			Detail: &h5.Detail{
				GoodsDetail: []h5.GoodsDetail{h5.GoodsDetail{
					GoodsName:        core.String("iPhoneX 256G"),
					MerchantGoodsId:  core.String("ABC"),
					Quantity:         core.Int64(1),
					UnitPrice:        core.Int64(1), // 订单待完善
					WechatpayGoodsId: core.String("1001"),
				}},
				InvoiceId: core.String("wx123"),
			},
			SceneInfo: &h5.SceneInfo{
				PayerClientIp: core.String("1.170.210.98"), // 必须
				H5Info: &h5.H5Info{
					AppName: core.String("应用名称"),
					AppUrl:  core.String("https://www.baidu.com"),
					Type:    core.String("iOS"),
				},
			},
		}
		resp, _, err := client.H5ApiService.Prepay(ctx, request)
		if err != nil {
			panic(err)
		}
		t.Log(resp.H5Url)
	})

}
