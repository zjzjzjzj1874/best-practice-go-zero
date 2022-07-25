package mysql

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
)

var orm = MustNewDB(MysqlConfig{
	//DSN:           "root:admin123@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
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
	t.Run("#Create", func(t *testing.T) {
		test := Test{
			TestName: "akjbe",
		}

		fmt.Println(orm.Create(&test).Error)
	})

	t.Run("#Created && Migrate", func(t *testing.T) {
		test := Test{
			TestName: "jhvjhf",
			Emotion:  " ლ(ʘ▽ʘ)ლ   (๑ↀᆺↀ๑)✧",
		}
		orm.AutoMigrate(&test)

		fmt.Println(orm.Create(&test).Error)
	})
	t.Run("#Select", func(t *testing.T) {
		test := Test{
			TestName: "jhvjhf",
		}
		orm.AutoMigrate(&test)
		err := orm.First(&test).Error
		if err != nil {
			t.Log(err)
			return
		}
		fmt.Println(test)
	})

	t.Run("update time", func(t *testing.T) {
		test := Test{
			Model: gorm.Model{
				ID: 1,
			},
		}
		update := map[string]interface{}{
			"receive_time": time.Now(),
			"test_time":    time.Now(),
			"callback_time": sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		}

		fmt.Println(orm.Model(&test).Updates(update).Error)
	})

}
