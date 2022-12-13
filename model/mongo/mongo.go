// Package mongo mongoDB数据库相关
package mongo

import (
	"github.com/globalsign/mgo/bson"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/constants/types"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/sqlx"
	"time"
)

// Test mongoDB的测试表 => TODO go-zero不带-c的生成有问题,因为go-zero/core/stores/mongo/collection中有一些方法没有实现,所以只能用这个
//go:generate goctl model mongo -t MongoTest -c --style go_zero -d .
type Test struct {
	ID        bson.ObjectId    `bson:"_id" json:"id" description:"主键ID"`
	TestName  string           `bson:"test_name" json:"test_name" description:"测试名称"`
	CreatedAt int64            `bson:"created_at" json:"created_at" description:"创建时间"`
	Hobbies   sqlx.StringSlice `bson:"hobbies" json:"hobbies" description:"爱好"`
}

//go:generate goctl model mongo -t TestService --home ../ --project my_zero -c --style go_zero -d .
type TestService struct {
	ID        bson.ObjectId    `bson:"_id" json:"id" description:"主键ID"`
	TestName  string           `bson:"test_name" json:"test_name" description:"测试名称"`
	CreatedAt int64            `bson:"created_at" json:"created_at" description:"创建时间"`
	Hobbies   sqlx.StringSlice `bson:"hobbies" json:"hobbies" description:"爱好"`
}

// ExportTask 导出任务表
//go:generate goctl model mongo -type ExportTask -c -style goZero -d .
type ExportTask struct {
	ID           bson.ObjectId     `json:"id" bson:"_id" description:"ObjectID"`
	ExportType   types.ExportType  `json:"export_type" bson:"export_type" description:"导出类型"`
	ExportState  types.ExportState `json:"export_state" bson:"export_state" description:"导出状态"`
	RequestArgs  string            `json:"request_args" bson:"request_args" description:"请求参数"`
	TargetEmails []string          `json:"target_emails" bson:"target_emails" description:"导出到邮箱"`
	Msg          string            `json:"msg" bson:"msg" description:"信息"`
	EmailUrl     string            `json:"email_url" bson:"email_url" description:"报表obs地址:带过期时间"`
	PostNum      int               `json:"post_num" bson:"post_num" description:"发送邮件次数"`
	CreateTime   time.Time         `json:"createTime" bson:"createTime" description:"创建时间"`
	UpdateTime   time.Time         `json:"updateTime" bson:"updateTime" description:"更新时间"`
	Operator     *Operator         `json:"operator" bson:"operator" description:"操作人"`
}

type Operator struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name" description:"姓名"`
}
