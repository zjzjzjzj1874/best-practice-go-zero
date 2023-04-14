package types

//go:generate tools gen enum PayChannel
// PayChannel 支付渠道
type PayChannel uint8

const (
	PAY_CHANNEL_UNKNOWN PayChannel = iota
	PAY_CHANNEL__WECHAT            // 微信支付
	PAY_CHANNEL__ALI               // 支付宝支付
)
