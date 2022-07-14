package order

import (
	"context"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/order/api/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/order/api/internal/types"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// TODO add your own logic
	userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.OrderReply{
		Id:    userInfo.Id,
		Name:  userInfo.Name,
		Email: userInfo.Email,
	}, nil
}
