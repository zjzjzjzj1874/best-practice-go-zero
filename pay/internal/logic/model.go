package pay

import (
	"errors"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types"
	"time"
)

var (
	ErrPayNotRegister = errors.New("支付方式未注册")
	ErrTradeQuery     = errors.New("查询账单失败")
	ErrTradePaying    = errors.New("查询支付中,请稍后查看")
)

type (
	OrderPayResult struct { // 订单支付结果
		TradeNo    string           // 支付交易号(支付宝:TradeNo;微信:TransactionId == 支付订单号)
		OrderNo    string           // 商户订单号(支付宝:OutTradeNo;微信:OutTradeNo)
		OrderState types.TradeState // 订单支付结果

		ReceiptAmount  string `json:"receipt_amount"`   // 实收金额，单位为元，两位小数
		BuyerPayAmount string `json:"buyer_pay_amount"` // 买家实付金额，单位为元，两位小数。
		PointAmount    string `json:"point_amount"`     // 积分支付的金额，单位为元，两位小数。
		InvoiceAmount  string `json:"invoice_amount"`   // 交易中用户支付的可开具发票的金额，单位为元，两位小数。		BuyerUserId           string           `json:"buyer_user_id"`                 // 买家在支付宝的用户id
		BuyerUserId    string `json:"buyer_user_id"`    // 买家在支付宝的用户id
	}

	CallbackResult struct { // 订单回调结果
		ID         string    `json:"id" description:"通知的唯一ID"`
		CreateTime time.Time `json:"create_time" description:"通知创建的时间，遵循rfc3339标准格式"`
		EventType  string    `json:"event_type" description:"通知的类型，支付成功通知的类型为TRANSACTION.SUCCESS"`

		Algorithm      string `json:"algorithm" description:"对开启结果数据进行加密的加密算法，目前只支持AEAD_AES_256_GCM"`
		Ciphertext     string `json:"ciphertext" description:"Base64编码后的开启/停用结果数据密文"`
		AssociatedData string `json:"associated_data" description:"附加数据"`
		Nonce          string `json:"nonce" description:"加密使用的随机串"`
	}
)
