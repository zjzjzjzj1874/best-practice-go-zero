package pay

import (
	"context"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
)

func GetPay(ctx context.Context, svcCtx *svc.ServiceContext, pc types.PayChannel) (Pay, error) {
	switch pc {
	case types.PAY_CHANNEL__WECHAT:
		return NewWechatPay(ctx, svcCtx), nil
	case types.PAY_CHANNEL__ALI:
		return NewAliPay(ctx, svcCtx), nil
	default:
		return nil, ErrPayNotRegister
	}
}

type (
	Pay interface {
		WithHost(host string) Pay
		WithReturnUrl(url string) Pay
		Prepay(ctx context.Context, order interface{}) (string, error)                              // 预支付接口
		QueryPayResult(ctx context.Context, orderTradeId string) (payRes OrderPayResult, err error) // 根据交易号查询支付状态查询
		Callback(ctx context.Context, info CallbackResult) (payRes OrderPayResult, err error)       // 回调处理
	}
)
