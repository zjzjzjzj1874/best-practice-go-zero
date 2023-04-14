package pay

import (
	"context"
	"github.com/sirupsen/logrus"
	types2 "github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types"
	pay "github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/logic"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/pay/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	host   string
}

func NewPayLogic(ctx context.Context, svcCtx *svc.ServiceContext, host string) *PayLogic {
	return &PayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		host:   host,
	}
}

func (l *PayLogic) Pay(req *types.PayReq) (resp *types.PayResp, err error) {
	// 订单信息预处理
	payClient, err := pay.GetPay(l.ctx, l.svcCtx, types2.PAY_CHANNEL__WECHAT)
	if err != nil {
		logrus.Errorf("pay.GetPay failure:[id:%d,err:%s]", req.ID, err.Error())
		return nil, err
	}

	redirectURL, err := payClient.WithHost(l.host).WithReturnUrl(req.ReturnUrl).Prepay(l.ctx, "订单信息")
	if err != nil {
		logrus.Errorf("pay.Prepay failure:[id:%d,err:%s]", req.ID, err.Error())
		return nil, err
	}

	// TODO 获取成功后修改数据库状态
	logrus.Infof("PayUrl:%s", redirectURL)

	return &types.PayResp{PayUrl: redirectURL}, nil
}
