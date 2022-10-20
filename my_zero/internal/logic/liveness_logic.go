package logic

import (
	"context"
	"golang.org/x/sync/singleflight"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LivenessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLivenessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LivenessLogic {
	return &LivenessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Liveness singleflight => singleflight的原理是当同时有很多请求同时到来时，最终只有一个请求会最终访问到资源，
// 其他请求都会等待结果然后返回 => https://www.cyningsun.com/01-11-2021/golang-concurrency-singleflight.html
// 以及singleFlight的一些缺点:https://juejin.cn/post/7093859835694809125
func (l *LivenessLogic) Liveness(req *types.LivenessRequest) (resp *types.LivenessResponse, err error) {
	sf := singleflight.Group{}
	v, err, _ := sf.Do("", func() (interface{}, error) {

		return resp, nil
	})
	if err != nil {
		logx.Errorf("do failure:[err:%s]", err.Error())
		return nil, err
	}
	return v.(*types.LivenessResponse), nil
}
