package databases

import (
	"fmt"
	"testing"
)

func TestNewDB(t *testing.T) {
	orm := MustNewDB(MysqlConfig{
		DSN:           "root:admin123@tcp(localhost:3306)/scana_audit?charset=utf8&parseTime=True&loc=Local",
		SlowThreshold: 200,
	})
	fmt.Println(orm)
	mock := Mock{
		TestName: "this is test",
	}
	_ = orm.AutoMigrate(&mock)

	fmt.Println(orm.Create(&mock).Error)
}
