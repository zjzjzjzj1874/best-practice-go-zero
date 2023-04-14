package pay

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	types2 "github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types"
	pay "github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/logic"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPayResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PayResultLogic {
	return &PayResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PayResultLogic) PayResult(req *types.PayQueryReq) (resp *types.PayResultResp, err error) {
	resp = &types.PayResultResp{}
	// 	1.查询数据库支付结果和状态;
	// 	1.1 收到微信/支付宝回调通知,直接返回订单结果
	// 	1.2 没有收到通知,去微信/支付宝查询订单
	// 	2. 调用微信/支付宝支付查询API查询订单支付状态
	//  2.1 微信和支付宝的金额单位不同,到时记得数据转换

	// 只有支付中才去支付宝或者微信那边查询
	payClient, err := pay.GetPay(l.ctx, l.svcCtx, types2.PAY_CHANNEL__ALI) // 支付类型 1支付宝 2微信
	if err != nil {
		logrus.Errorf("【Alipay PayResult】PayOrderQuery GetPay failure:[id:%d,err:%s]", req.ID, err.Error())
		return nil, err
	}

	payRes, err := payClient.QueryPayResult(l.ctx, "订单ID")
	if err != nil {
		logrus.Errorf("【Alipay PayResult】PayOrderQuery QueryPayResult failure:[id:%d,orderNo:%s,err:%s]", req.ID, req.ID, err.Error())
		return nil, err
	}

	// TODO 查询后
	// 	1.更新相关表
	err = copier.Copy(resp, &payRes)
	return
}
