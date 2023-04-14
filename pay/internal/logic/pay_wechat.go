// Package pay 微信支付
package pay

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/pay/wechatpay"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"

	"github.com/sirupsen/logrus"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/native"
)

var (
	wechatPay  *WechatPay
	wechatSync sync.Once
)

type WechatPay struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	host      string
	returnUrl string // 支付成功后跳转链接 ==> But:微信支付成功后不支持传入跳转链接,没办法了...
}

func NewWechatPay(ctx context.Context, svcCtx *svc.ServiceContext) *WechatPay {
	wechatSync.Do(func() {
		wechatPay = &WechatPay{
			ctx:    ctx,
			svcCtx: svcCtx,
		}
	})

	return wechatPay
}

func (p *WechatPay) WithHost(host string) Pay {
	if host != "" {
		p.host = host
	}
	return p
}

func (p *WechatPay) WithReturnUrl(url string) Pay {
	if url != "" {
		p.returnUrl = url
	}
	return p
}

func (p *WechatPay) Prepay(ctx context.Context, order interface{}) (string, error) {
	request := native.PrepayRequest{
		Appid:         core.String(p.svcCtx.Config.WechatPay.AppID),
		Mchid:         core.String(p.svcCtx.Config.WechatPay.MchID),
		Description:   core.String("描述"),
		OutTradeNo:    core.String("123456"),                       // 我方订单ID
		TimeExpire:    core.Time(time.Now().Add(15 * time.Minute)), // 默认15分钟过期
		Attach:        core.String("附件信息"),
		NotifyUrl:     core.String(p.svcCtx.Config.WechatPay.NotifyUrl), // 微信回调地址
		GoodsTag:      core.String("商品Tag"),
		SupportFapiao: core.Bool(false),
		Amount: &native.Amount{
			Currency: core.String("CNY"),
			Total:    core.Int64(1), // 单位:分
		},
		Detail: &native.Detail{
			GoodsDetail: []native.GoodsDetail{native.GoodsDetail{
				GoodsName:       core.String("商品名"),
				MerchantGoodsId: core.String("1"), // 商品ID
				Quantity:        core.Int64(1),
				UnitPrice:       core.Int64(1), // 商品单价:单位分
			}},
		},
		SceneInfo: &native.SceneInfo{
			PayerClientIp: core.String(p.host),
			DeviceId:      core.String("商户端设备号"),
			StoreInfo: &native.StoreInfo{
				Id:   core.String("商户侧门店编号"),
				Name: core.String("商户侧门店名称"),
			},
		},
	}
	resp, _, err := p.svcCtx.WxNativePayClient.Prepay(ctx, request)
	if err != nil || resp.CodeUrl == nil {
		logrus.Errorf("Prepay failure:[OrderId:%s,err:%v]", "order_id", err)
		return "", err
	}

	return *resp.CodeUrl, nil
}

func (p *WechatPay) QueryPayResult(ctx context.Context, orderTradeId string) (payRes OrderPayResult, err error) {
	request := native.QueryOrderByIdRequest{
		TransactionId: core.String(orderTradeId),
		Mchid:         core.String(p.svcCtx.Config.WechatPay.MchID),
	}

	// 调用微信查询API查询订单支付状态
	res, result, err := p.svcCtx.WxNativePayClient.QueryOrderById(p.ctx, request)
	if err != nil {
		logrus.Errorf("Wechatpay QueryOrderById failure:[orderTradeId:%s,err:%s]", orderTradeId, err.Error())
		return
	}

	if result.Response.StatusCode == http.StatusAccepted {
		// 用户支付中，需要输入密码 ==> 等待5秒，然后调用被扫订单结果查询API，查询当前订单的不同状态，决定下一步的操作
		logrus.Infof("Wechatpay QueryOrderById Paying:[orderTradeId:%s,result.Response.StatusCode:%d]", orderTradeId, result.Response.StatusCode)
		err = ErrTradePaying
		return
	}
	if result.Response.StatusCode >= http.StatusMultipleChoices { // 状态码 >= 300均表示失败
		logrus.Errorf("Wechatpay QueryOrderById Failure:[orderTradeId:%s,result.Response.StatusCode:%d]", orderTradeId, result.Response.StatusCode)
		err = ErrTradeQuery
		return
	}

	// 微信单位转化
	payRes.OrderNo = *res.OutTradeNo
	return OrderPayResult{
		TradeNo:        *res.TransactionId,
		OrderNo:        *res.OutTradeNo,
		OrderState:     wechatTradeState(*res.TradeState),
		ReceiptAmount:  strconv.Itoa(int(*res.Amount.Total)),
		BuyerPayAmount: strconv.Itoa(int(*res.Amount.PayerTotal)),
		BuyerUserId:    *res.Payer.Openid,
	}, nil
}

// Callback 微信通知回调
func (p *WechatPay) Callback(ctx context.Context, info CallbackResult) (payRes OrderPayResult, err error) {
	// 处理微信回调结果,根据加密反解body
	pt, err := wechatpay.DecryptAES256GCM(p.svcCtx.Config.WechatPay.MchAPIv3Key, info.AssociatedData, info.Nonce, info.Ciphertext)
	if err != nil {
		logrus.Errorf("【WECHAT NOTIFY】DecryptAES256GCM Callback Info failure:[err:%s]", err.Error())
		return
	}

	dr := wechatpay.NotifyDecryptResource{}
	err = json.Unmarshal(pt, &dr)
	if err != nil {
		logrus.Errorf("【WECHAT NOTIFY】Unmarshal Callback Info failure:[resource:%+v,err:%s]", pt, err.Error())
		return
	}
	payRes = OrderPayResult{
		TradeNo:        dr.TransactionID,
		OrderNo:        dr.OutTradeNo,
		OrderState:     wechatTradeState(dr.TradeState),
		ReceiptAmount:  strconv.Itoa(dr.Amount.Total),
		BuyerPayAmount: strconv.Itoa(dr.Amount.PayerTotal),
		BuyerUserId:    dr.Payer.Openid,
	}
	// TODO 逻辑处理
	return payRes, err
}

// 交易状态 	SUCCESS：支付成功
//			REFUND：转入退款
//			NOTPAY：未支付
//			CLOSED：已关闭
//			REVOKED：已撤销（付款码支付）
//			USERPAYING：用户支付中（付款码支付）
//			PAYERROR：支付失败(其他原因，如银行返回失败)
func wechatTradeState(tradeState string) types.TradeState {
	// 因为tradeState是根据微信支付的状态定义的,所以可以这么反解,支付宝需要手动映射
	ts, err := types.ParseTradeStateFromString(tradeState)
	if err != nil {
		return types.TRADE_STATE__PAYERROR
	}

	return ts
}
