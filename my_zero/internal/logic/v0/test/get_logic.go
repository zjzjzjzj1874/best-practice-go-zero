package test

import (
	"context"
	"fmt"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/__test__/gorm/dal/model"
	"time"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) TestGet(req *types.ListReq) (resp *types.MongoTest, err error) {
	//mt, err := l.svcCtx.MongoTestModel.FindOne(req.ID)
	//if err != nil {
	//	return nil, err
	//}
	//resp = &types.MongoTest{}
	//if err = copier.CopyWithOption(resp, mt, helper.OutOption()); err != nil {
	//	return nil, err
	//}

	// 测试异步是否有问题
	go l.AsyncGet(req)

	return &types.MongoTest{
		ID:        req.ID,
		TestName:  "测试",
		CreatedAt: time.Now().Unix(),
	}, nil
}

func (l *GetLogic) AsyncGet(req *types.ListReq) {
	for i := 10; i >= 0; i-- {
		fmt.Printf("i == %d, goto next loop.\n", i)
		user, err := l.getOne(l.ctx)
		if err != nil {
			fmt.Printf("error ==== %s\n", err.Error())
			return
		}
		fmt.Println(user)
		time.Sleep(time.Second * 5)
	}
}

func (l *GetLogic) getOne(ctx context.Context) (*model.TUser, error) {
	user := model.TUser{}
	err := l.svcCtx.MysqlDB.First(&user).Error
	if err != nil {
		fmt.Printf("error ==== %s\n", err.Error())
		return nil, err
	}
	return &user, nil
}
