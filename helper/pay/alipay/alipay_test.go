package alipay

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"testing"
)

var conf = Conf{
	IsProd:    true,
	NotifyUrl: "https://www.baidu.com/pay/callback/ali",
	AppId:     AppId,
}

func TestMustNewAliPayClient(t *testing.T) {
	t.Run("#Alipay", func(t *testing.T) {
		client := MustNewAliPayClient(conf)

		var p = alipay.TradePagePay{}
		p.Subject = "支付测试"
		p.OutTradeNo = "123456788"
		p.TotalAmount = "0.01"
		p.ProductCode = "FAST_INSTANT_TRADE_PAY"

		p.GoodsDetail = []*alipay.GoodsDetail{&alipay.GoodsDetail{
			GoodsId:   "123",
			GoodsName: "xxx",
			Quantity:  1,
			Price:     0.01,
		}}

		res, err := client.TradePagePay(p)
		if err != nil {
			panic(err)
		}

		fmt.Println(res)
		// https://openapi.alipay.com/gateway.do?
		//app_id=2019010462816111&
		//biz_content=%7B%22subject%22%3A%22%E4%BF%AE%E6%AD%A3%E4%BA%86%E4%B8%AD%E6%96%87%E7%9A%84%22%2C%22
		//out_trade_no%22%3A%22trade_no_201706230111212%22%2C%22
		//total_amount%22%3A%220.01%22%2C%22
		//product_code%22%3A%22%E6%B5%8B%E8%AF%95%22%2C%22
		//goods_detail%22%3A%5B%7B%22goods_id%22%3A%22123%22%2C%22goods_name%22%3A%22%E6%88%91%E6%98%AF%E6%B5%8B%E8%AF%95%22%2C%22quantity%22%3A1%2C%22price%22%3A0.01%7D%5D%7D&
		//charset=utf-8&
		//format=JSON&
		//method=alipay.trade.page.pay&
		//notify_url=&
		//return_url=&
		//sign=Oh8tGcKZoHzLsX%2BS6dp26KSn5d3gJMTB%2BzbjLD23o%2BJMCfXIgOx6tBFLkDX10xkjh8j%2Bb1DNgqgFCbZWKwxHuvzIz2o3ANhJ7XrCh6oDejkgCfgnwzz%2FggRi9pQMC5fklkMGtGzKZmmE4eGBbibTiD87zF42ktJMcCl0NFM71p6nYe0dpDmpgNYX4Kq2HAmpWwdhEIIRcpOCHfPR%2BVCL4btzY2LyBdMS1FjGr812wprMgYMRn4N5EhYY2z%2BqmBAuuC3SsWZ7pRr2VTZt54SV14EO5UBhjjzUePOsN5FCvzXEekViaL7vU9D3%2B%2BY2hmCjW%2F0ltV6NFazKUjPqtL6dwA%3D%3D&
		//sign_type=RSA2&
		//timestamp=2023-04-12+16%3A27%3A09&
		//version=1.0
	})

}
