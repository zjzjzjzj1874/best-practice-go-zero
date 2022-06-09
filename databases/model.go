package databases

import (
	"database/sql"
	"database/sql/driver"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/sqlx"
	"gorm.io/gorm"
	"time"
)

// Mock mock表
type Mock struct {
	gorm.Model
	TestName string           `gorm:"test_name,comment:测试名称" json:"test_name" description:"测试名称"`
	Hobbies  sqlx.StringSlice `gorm:"hobbies;comment:爱好" json:"hobbies" description:"爱好"`
	SlaTimes SlaTimes         `gorm:"sla_times;comment:对象测试字段" json:"sla_times" description:"对象测试"`
}

func (Mock) TableName() string {
	return "t_mock"
}

// Test test表
type Test struct {
	gorm.Model
	ReceiveTime sql.NullTime `gorm:"column:receive_time;comment:收货时间"` // 收货时间
	UpdateTime  time.Time    `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	TestName    string       `gorm:"uniqueIndex:idx_name;comment:测试名称" json:"test_name"`
}

func (Test) TableName() string {
	return "t_test"
}

// User 用户表 1.联合索引构建 => 不同字段上index名字一样,即可创建联合索引;2.唯一索引的两种方式(邮箱和姓名字段)
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(64);not null;default:'';column:name;uniqueIndex:idx_name;comment:姓名"`
	Email    string `gorm:"column:email;index:idx_email,unique;type:varchar(255);not null;default:'';comment:邮箱"`
	NickName string `gorm:"index:idx_nick_name_age;type:varchar(64);not null;default:'';comment:昵称"`
	Age      int    `gorm:"index:idx_nick_name_age;type:int(11);not null;default:0;comment:年龄"`
	Phone    string `gorm:"index:idx_nick_name_age;type:varchar(18);not null;default:'';comment:电话号码"`
}

func (User) TableName() string {
	return "t_user"
}

type (
	SlaTimes []SlaTime
	SlaTime  struct {
		StartTime int `json:"start_time"`
		EndTime   int `json:"end_time"`
		During    int `json:"during"`
	}
)

func (s *SlaTimes) Scan(src interface{}) error {
	return sqlx.JSONScan(src, s)
}

func (s SlaTimes) Value() (driver.Value, error) {
	return sqlx.JSONValue(s)
}
