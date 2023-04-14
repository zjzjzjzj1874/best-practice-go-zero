package alipay

import "time"

// Conf 支付宝支付配置
type Conf struct {
	NotifyUrl string `json:",default=https://www.baidu.com/pay/callback/ali"` // 微信支付回调通知url,必须是https
	IsProd    bool   `json:",default=true"`                                   // 是否是生产环境:默认是,否则为沙箱
	AppId     string // = "190000****"                               // 商户号
}

// NotifyReq 微信支付回调接口
type NotifyReq struct {
	ID           string    `json:"id"`
	CreateTime   time.Time `json:"create_time"`
	ResourceType string    `json:"resource_type"`
	EventType    string    `json:"event_type"`
	Summary      string    `json:"summary"`
	Resource     struct {
		OriginalType   string `json:"original_type"`
		Algorithm      string `json:"algorithm"`
		Ciphertext     string `json:"ciphertext"`
		AssociatedData string `json:"associated_data"`
		Nonce          string `json:"nonce"`
	} `json:"resource"`
}
