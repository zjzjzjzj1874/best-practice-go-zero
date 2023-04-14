package callback

import (
	"context"
	"github.com/sirupsen/logrus"
	pay "github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/logic"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
)

type CallbackAliLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallbackAliLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackAliLogic {
	return &CallbackAliLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CallbackAli 异步通知参数 ==> https://opendocs.alipay.com/open/203/105286
func (l *CallbackAliLogic) CallbackAli(r *http.Request) error {
	notification, err := l.svcCtx.Client.GetTradeNotification(r)
	if err != nil {
		logrus.Errorf("【ALIPAY NOTIFY】GetTradeNotification Failure:[err:%s]", err.Error())
		return err
	}

	alipay := pay.NewAliPay(l.ctx, l.svcCtx)
	payRes, err := alipay.WithNotification(notification).Callback(l.ctx, pay.CallbackResult{})
	if err != nil {
		logrus.Errorf("【ALIPAY NOTIFY】Callback Failure:[err:%s]", err.Error())
		return err
	}

	// 这是回调信息打印 ==> 后面可能需要周期轮询,轮询在支付中状态的订单,如果有过期或者什么,更新那个订单状态
	logrus.Infof("【ALIPAY NOTIFY】PayRes:%+v", payRes)
	return nil
}
