package svc

import (
	"reflect"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/config"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/db"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    helper.Client
	MysqlDB        *db.Mysql // 本服务自己封装的一个mysql client
	MongoTestModel model.MongoTestModel
	//GormDB      *gorm.DB// database中的gorm.DB客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		RedisClient:    helper.MustNewClient(c.Mode, c.CacheRedis),
		MysqlDB:        db.NewMysqlClient(model.MustNewDB(c.MysqlConf)),
		MongoTestModel: model.NewMongoTestModel(c.MongoDB.URL, reflect.TypeOf(model.MongoTest{}).Name(), c.CacheRedis),
	}
}
