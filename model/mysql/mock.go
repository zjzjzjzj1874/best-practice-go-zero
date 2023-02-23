package mysql

import (
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/sqlx"
	"gorm.io/gorm"
)

// Mock mock表
//go:generate goctl model mysql datasource "mysql://root:admin123@tcp(127.0.0.1:3306)/test" -table mock -dir .
type Mock struct {
	gorm.Model
	TestName string           `gorm:"test_name,comment:测试名称" json:"test_name" description:"测试名称"`
	Hobbies  sqlx.StringSlice `gorm:"hobbies;comment:爱好" json:"hobbies" description:"爱好"`
	SlaTimes SlaTimes         `gorm:"sla_times;comment:对象测试字段" json:"sla_times" description:"对象测试"`
}

func (Mock) TableName() string {
	return "t_mock"
}
