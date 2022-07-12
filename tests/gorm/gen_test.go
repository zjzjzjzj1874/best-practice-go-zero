package gorm

import (
	"context"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/tests/gorm/dal/model"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/tests/gorm/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"testing"
	"time"
)

var dsn = "root:admin123@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 使用db生成model
func generateModel() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
	})

	g.UseDB(DB)

	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

func TestName(t *testing.T) {
	generateModel()
}

func TestName1(t *testing.T) {
	test := model.TTest{
		TestName:     "this is a test",
		Emotion:      "(灬ꈍ ꈍ灬)",
		ReceiveTime:  time.Now(),
		UpdateTime:   time.Now(),
		CallbackTime: time.Now(),
		TestTime:     time.Now(),
		TestTime1:    time.Now(),
		//Emotion:  "😭",
	}

	tt := query.Use(DB).TTest

	expr, ok := tt.GetFieldByName("test_name")
	if !ok {
		return
	}
	t.Log(expr)

	err := tt.WithContext(context.TODO()).Create(&test)
	if err != nil {
		t.Fatal(err)
	}
}
