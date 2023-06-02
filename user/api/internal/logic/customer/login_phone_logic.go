package customer

import (
	"context"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginPhoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginPhoneLogic {
	return &LoginPhoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginPhoneLogic) LoginPhone(req *types.LoginPhoneReq) (resp *types.LoginPhoneResp, err error) {
	// todo: add your logic here and delete this line

	return
}
