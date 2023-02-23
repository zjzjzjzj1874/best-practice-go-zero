package mysql

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

// Test test表
//go:generate goctl model mysql datasource "mysql://root:admin123@tcp(127.0.0.1:3306)/test" -table t_test -dir .
type Test struct {
	gorm.Model
	ReceiveTime  sql.NullTime `gorm:"column:receive_time;comment:收货时间"` // 收货时间
	UpdateTime   time.Time    `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CallbackTime sql.NullTime `gorm:"column:callback_time;type:timestamp;comment:回调时间"`
	TestTime     sql.NullTime `gorm:"column:test_time;comment:测试时间"`
	TestName     string       `gorm:"comment:测试名称"`
	Emotion      string       `gorm:"comment:表情"`
}

func (Test) TableName() string {
	return "t_test"
}
