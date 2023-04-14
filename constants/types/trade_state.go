package types

//go:generate tools gen enum TradeState
// TradeState 支付状态
type TradeState uint8

const (
	TRADE_STATE_UNKNOWN     TradeState = iota
	TRADE_STATE__NOTPAY                // 未支付
	TRADE_STATE__USERPAYING            // 用户支付中
	TRADE_STATE__SUCCESS               // 支付成功
	TRADE_STATE__REFUND                // 转入退款
	TRADE_STATE__CLOSED                // 已关闭
	TRADE_STATE__REVOKED               // 已撤销
	TRADE_STATE__PAYERROR              // 支付失败
)

// region 支付宝状态转化

type AliTradeState string

const (
	WAIT_BUYER_PAY AliTradeState = "WAIT_BUYER_PAY" // 交易创建，等待买家付款
	TRADE_CLOSED   AliTradeState = "TRADE_CLOSED"   // 未付款交易超时关闭，或支付完成后全额退款
	TRADE_SUCCESS  AliTradeState = "TRADE_SUCCESS"  // 交易支付成功
	TRADE_FINISHED AliTradeState = "TRADE_FINISHED" // 交易结束，不可退款
)

func (a AliTradeState) Transfer() (ts TradeState) {
	switch a {
	case TRADE_CLOSED:
		ts = TRADE_STATE__CLOSED
	case TRADE_FINISHED, TRADE_SUCCESS: // 在支付宝的业务通知中，只有交易通知状态为 TRADE_SUCCESS 或 TRADE_FINISHED 时，支付宝才会认定为买家付款成功。
		ts = TRADE_STATE__SUCCESS
	case WAIT_BUYER_PAY:
		ts = TRADE_STATE__NOTPAY
	default:
		ts = TRADE_STATE__PAYERROR
	}

	return
}

// endregion 支付宝状态转化
