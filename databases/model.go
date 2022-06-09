package databases

import (
	"database/sql/driver"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/helper/sqlx"
	"gorm.io/gorm"
)

type Mock struct {
	gorm.Model
	TestName string           `gorm:"test_name" json:"test_name" description:"测试名称"`
	Hobbies  sqlx.StringSlice `gorm:"hobbies" json:"hobbies" description:"爱好"`
	SlaTimes SlaTimes         `gorm:"sla_times" json:"sla_times" description:"对象测试"`
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
