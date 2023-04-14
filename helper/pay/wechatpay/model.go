package wechatpay

import "time"

// Conf 微信支付配置
type Conf struct {
	IsProd                     bool   `json:",default=true"`                                      // 是否是生产环境:默认是
	NotifyUrl                  string `json:",default=https://www.baidu.com/pay/callback/wechat"` // 微信支付回调通知url,必须是https
	MchPrivateKeyPath          string `json:",default=/app/pay/apiclient_key.pem"`                // 商户私钥位置
	AppID                      string // = "190000****"                               // appid
	MchID                      string // = "190000****"                               // 商户号
	MchCertificateSerialNumber string // = "3775************************************" // 商户证书序列号
	MchAPIv3Key                string // = "2ab9****************************"         // 商户APIv3密钥
}

// NotifyReq 微信支付回调接口 see:https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_5.shtml
type NotifyReq struct {
	ID           string    `json:"id" description:"通知的唯一ID"`
	CreateTime   time.Time `json:"create_time" description:"通知创建的时间，遵循rfc3339标准格式"`
	ResourceType string    `json:"resource_type" description:"通知的资源数据类型，支付成功通知为encrypt-resource"`
	EventType    string    `json:"event_type" description:"通知的类型，支付成功通知的类型为TRANSACTION.SUCCESS"`
	Summary      string    `json:"summary" description:"回调摘要"`
	Resource     struct {
		OriginalType   string `json:"original_type" description:"原始回调类型，为transaction"`
		Algorithm      string `json:"algorithm" description:"对开启结果数据进行加密的加密算法，目前只支持AEAD_AES_256_GCM"`
		Ciphertext     string `json:"ciphertext" description:"Base64编码后的开启/停用结果数据密文"`
		AssociatedData string `json:"associated_data" description:"附加数据"`
		Nonce          string `json:"nonce" description:"加密使用的随机串"`
	} `json:"resource" description:"通知资源数据,json格式"`
}

// NotifyDecryptResource 微信回调解密数据结构 see:https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_3_5.shtml
type NotifyDecryptResource struct {
	TransactionID string `json:"transaction_id" description:"微信订单号"`
	Amount        struct {
		PayerTotal    int    `json:"payer_total"`
		Total         int    `json:"total"`
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"  description:"订单金额信息"`
	Mchid           string `json:"mchid"`
	TradeState      string `json:"trade_state"`
	BankType        string `json:"bank_type"`
	PromotionDetail []struct {
		Amount              int    `json:"amount"`
		WechatpayContribute int    `json:"wechatpay_contribute"`
		CouponID            string `json:"coupon_id"`
		Scope               string `json:"scope"`
		MerchantContribute  int    `json:"merchant_contribute"`
		Name                string `json:"name"`
		OtherContribute     int    `json:"other_contribute"`
		Currency            string `json:"currency"`
		StockID             string `json:"stock_id"`
		GoodsDetail         []struct {
			GoodsRemark    string `json:"goods_remark"`
			Quantity       int    `json:"quantity"`
			DiscountAmount int    `json:"discount_amount"`
			GoodsID        string `json:"goods_id"`
			UnitPrice      int    `json:"unit_price"`
		} `json:"goods_detail"`
	} `json:"promotion_detail"`
	SuccessTime time.Time `json:"success_time"`
	Payer       struct {
		Openid string `json:"openid" description:"用户在直连商户appid下的唯一标识"`
	} `json:"payer" description:"支付者信息"`
	OutTradeNo     string `json:"out_trade_no"`
	Appid          string `json:"appid"`
	TradeStateDesc string `json:"trade_state_desc"`
	TradeType      string `json:"trade_type"`
	Attach         string `json:"attach"`
	SceneInfo      struct {
		DeviceID string `json:"device_id"`
	} `json:"scene_info"`
}
