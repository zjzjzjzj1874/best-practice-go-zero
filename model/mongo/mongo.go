// Package model mongo数据库相关
package mongo

import (
	"github.com/globalsign/mgo/bson"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/sqlx"
)

// MongoTest mongoDB的测试表 => TODO go-zero不带-c的生成有问题,因为go-zero/core/stores/mongo/collection中有一些方法没有实现,所以只能用这个
//go:generate goctl model mongo -t MongoTest -c --style go_zero -d ./mongo
type MongoTest struct {
	ID        bson.ObjectId    `bson:"_id" json:"id" description:"主键ID"`
	TestName  string           `bson:"test_name" json:"test_name" description:"测试名称"`
	CreatedAt int64            `bson:"created_at" json:"created_at" description:"创建时间"`
	Hobbies   sqlx.StringSlice `bson:"hobbies" json:"hobbies" description:"爱好"`
}
