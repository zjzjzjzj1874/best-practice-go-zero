package test

import (
	"context"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/__test__/gorm/dal/model"
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

func (l *PostLogic) TestPost(req *types.ListReq) (*types.MongoTest, error) {
	var (
	//err  error
	//resp = &types.MongoTest{}
	//t = &mongo.Test{
	//	ID:        bson.NewObjectId(),
	//	TestName:  "Mongo test",
	//	CreatedAt: time.Now().Unix(),
	//	Hobbies:   []string{"羽毛球", "台球"},
	//}
	)

	user := model.TUser{
		Name:     "测试",
		Email:    "test@163.com",
		NickName: "tester",
		Age:      18,
		Phone:    "13636353433",
	}

	return nil, l.svcCtx.MysqlDB.Create(&user).Error
	// 添加一个基于函数式编程处理error的closure
	//errFunc := func(fn func() error) {
	//	if err != nil {
	//		return
	//	}
	//
	//	err = fn()
	//}
	//
	//errFunc(func() error { return l.svcCtx.MongoTestModel.Insert(t) })
	//
	//errFunc(func() error { return copier.CopyWithOption(resp, t, helper.OutOption()) })
	//
	//return resp, err
}
