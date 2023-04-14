// Package pay 支付宝支付
package pay

import (
	"context"
	"fmt"
	"sync"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"

	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay/v3"
)

var (
	aliPay  *AliPay
	aliSync sync.Once
)

type AliPay struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	*alipay.TradeNotification
	host      string
	returnUrl string // 支付成功后跳转链接
}

func NewAliPay(ctx context.Context, svcCtx *svc.ServiceContext) *AliPay {
	aliSync.Do(func() {
		aliPay = &AliPay{
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	})

	return aliPay
}

func (p *AliPay) WithNotification(notify *alipay.TradeNotification) *AliPay {
	if notify != nil {
		p.TradeNotification = notify
	}

	return p
}

func (p *AliPay) WithHost(host string) Pay {
	if host != "" {
		p.host = host
	}
	return p
}

func (p *AliPay) WithReturnUrl(url string) Pay {
	if url != "" {
		p.returnUrl = url
	}

	return p
}

func (p *AliPay) Prepay(ctx context.Context, order interface{}) (string, error) {
	ta := 0.01 // 单位转换:分=>元
	var param = alipay.TradePagePay{}
	param.NotifyURL = p.svcCtx.Config.AliPay.NotifyUrl
	param.Subject = "主题"
	param.OutTradeNo = "12345678" // 我方订单号
	param.TotalAmount = fmt.Sprintf("%.2f", ta)
	param.ProductCode = "FAST_INSTANT_TRADE_PAY" // Note:不可变更
	param.GoodsDetail = []*alipay.GoodsDetail{&alipay.GoodsDetail{
		GoodsId:   "1", // 货物ID
		GoodsName: "货物名称",
		Quantity:  1,
		Price:     ta,
	}}

	if p.returnUrl != "" {
		param.ReturnURL = p.returnUrl // 支付成功后跳转链接
	}

	result, err := p.svcCtx.Client.TradePagePay(param)
	if err != nil || result == nil {
		logrus.Errorf("Prepay failure:[OrderId:%s,err:%v]", param.OutTradeNo, err)
		return "", err
	}

	return result.String(), nil
}

func (p *AliPay) QueryPayResult(ctx context.Context, orderTradeId string) (payRes OrderPayResult, err error) {
	result, err := p.svcCtx.Client.TradeQuery(alipay.TradeQuery{OutTradeNo: orderTradeId})
	if err != nil || result == nil {
		logrus.Errorf("Alipay TradeQuery failure:[orderTradeId:%s,err:%v]", orderTradeId, err)
		err = ErrTradeQuery
		return
	}
	// 支付宝单位转化
	payRes = OrderPayResult{
		TradeNo:        result.Content.TradeNo,
		OrderNo:        result.Content.OutTradeNo,
		OrderState:     types.AliTradeState(result.Content.TradeStatus).Transfer(),
		ReceiptAmount:  result.Content.ReceiptAmount,
		BuyerPayAmount: result.Content.BuyerPayAmount,
		PointAmount:    result.Content.PointAmount,
		InvoiceAmount:  result.Content.InvoiceAmount,
		BuyerUserId:    result.Content.BuyerUserId,
	}
	return
}

func (p *AliPay) Callback(ctx context.Context, info CallbackResult) (payRes OrderPayResult, err error) {
	payRes = OrderPayResult{
		TradeNo:        p.TradeNo,
		OrderNo:        p.OutTradeNo,
		OrderState:     types.AliTradeState(p.TradeStatus).Transfer(),
		ReceiptAmount:  p.ReceiptAmount,
		BuyerPayAmount: p.BuyerPayAmount,
		PointAmount:    p.PointAmount,
		InvoiceAmount:  p.InvoiceAmount,
		BuyerUserId:    p.BuyerId,
	}

	// TODO 处理我方逻辑
	return payRes, err
}
