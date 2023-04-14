package callback

import (
	"context"
	"github.com/sirupsen/logrus"
	pay "github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/logic"
	"time"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackWxLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCallbackWxLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackWxLogic {
	return &CallbackWxLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CallbackWx 通知应答
//  接收成功：HTTP应答状态码需返回200或204，无需返回应答报文。
//  接收失败：HTTP应答状态码需返回5XX或4XX，同时需返回应答报文，格式如下：
func (l *CallbackWxLogic) CallbackWx(req *types.WechatNotifyReq) error {
	logrus.Infof("【WECHAT NOTIFY】[Info:%+v]", req)

	wechatPay := pay.NewWechatPay(l.ctx, l.svcCtx)
	ct, _ := time.Parse(time.RFC3339, req.CreateTime)
	payRes, err := wechatPay.Callback(l.ctx, pay.CallbackResult{
		ID:             req.ID,
		CreateTime:     ct,
		EventType:      req.EventType,
		Algorithm:      req.Resource.Algorithm,
		Ciphertext:     req.Resource.Ciphertext,
		AssociatedData: req.Resource.AssociatedData,
		Nonce:          req.Resource.Nonce,
	})
	if err != nil { // 将错误包装成符合微信那边入参的
		logrus.Errorf("【WECHAT NOTIFY】Callback Logic Solve failure:[err:%v]", err.Error())
	}

	logrus.Infof("【WECHAT NOTIFY】PayRes:%+v", payRes)
	return err
}
