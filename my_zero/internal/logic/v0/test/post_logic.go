package test

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"github.com/jinzhu/copier"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mongo"
	"time"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostLogic {
	return &PostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostLogic) TestPost(req *types.ListReq) (resp *types.MongoTest, err error) {
	t := &mongo.Test{
		ID:        bson.NewObjectId(),
		TestName:  "Mongo test",
		CreatedAt: time.Now().Unix(),
		Hobbies:   []string{"羽毛球", "台球"},
	}

	// 添加一个基于函数式编程处理error的closure
	errFunc := func(fn func() error) {
		if err != nil {
			return
		}

		err = fn()
	}

	errFunc(func() error {
		return l.svcCtx.MongoTestModel.Insert(t)
	})

	resp = &types.MongoTest{}
	errFunc(func() error {
		return copier.CopyWithOption(resp, t, helper.OutOption())
	})

	return
}
