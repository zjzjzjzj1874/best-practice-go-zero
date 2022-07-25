package svc

import (
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mongo"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/model/mysql"
	"reflect"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/config"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/my_zero/internal/db"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    helper.Client
	MysqlDB        *db.Mysql // 本服务自己封装的一个mysql client
	MongoTestModel mongo.MongoTestModel
	//GormDB      *gorm.DB// database中的gorm.DB客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		RedisClient:    helper.MustNewClient(c.Mode, c.CacheRedis),
		MysqlDB:        db.NewMysqlClient(mysql.MustNewDB(c.MysqlConf)),
		MongoTestModel: mongo.NewMongoTestModel(c.MongoDB.URL, reflect.TypeOf(mongo.Test{}).Name(), c.CacheRedis),
	}
}
