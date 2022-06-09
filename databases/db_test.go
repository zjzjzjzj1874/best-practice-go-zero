package databases

import (
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
)

var orm = MustNewDB(MysqlConfig{
	DSN:           "root:admin123@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local",
	SlowThreshold: 200,
})

func TestNewDB(t *testing.T) {
	fmt.Println(orm)
	mock := Mock{
		TestName: "this is test",
		Hobbies:  []string{"badminton", "football"},
		SlaTimes: []SlaTime{
			{
				StartTime: 6,
				EndTime:   16,
				During:    10,
			},
			{
				StartTime: 0,
				EndTime:   6,
				During:    6,
			},
		},
	}
	_ = orm.AutoMigrate(&mock)

	fmt.Println(orm.Create(&mock).Error)
}

func TestGetFromDB(t *testing.T) {
	mock := Mock{}

	orm.First(&mock)
	fmt.Printf("result:%+v", mock)
}

func TestAddDataIntoTableTest(t *testing.T) {
	t.Run("create test time", func(t *testing.T) {
		test := Test{
			TestName: "this is test",
		}

		fmt.Println(orm.Create(&test).Error)
	})

	t.Run("update time", func(t *testing.T) {
		test := Test{
			Model: gorm.Model{
				ID: 1,
			},
		}
		update := map[string]interface{}{
			"receive_time": time.Now(),
		}

		fmt.Println(orm.Model(&test).Updates(update).Error)
	})

}
