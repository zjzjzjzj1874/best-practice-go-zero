package helper

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestValidate(t *testing.T) {
	//设置支持语言
	chinese := zh.New()
	//设置国际化翻译器
	uni := ut.New(chinese, chinese)
	//设置翻译语言
	trans, found := uni.GetTranslator("zh")
	if !found {
		log.Fatalf("no zh Translator")
	}
	// 创建validator
	validate_ := validator.New()
	// 注册获取字段名的方法
	validate_.RegisterTagNameFunc(func(field reflect.StructField) string {
		tagValue, ok := field.Tag.Lookup("description")
		if !ok {
			// 获取不到返回字段名
			return field.Name
		}
		return tagValue
	})
	//注册默认翻译
	if err := zhtrans.RegisterDefaultTranslations(validate_, trans); err != nil {
		log.Fatalf("validator register zh translation failed [err:%s]", err.Error())
	}
	type AddTest struct {
		Test1        string  `json:"test1" description:"测试1" validate:"required,min=3,max=66"`
		Count        uint    `json:"count" description:"量级(次)" validate:"gt=0"`
		Price        uint    `json:"price,optional" description:"价格"`
		Discount     float64 `json:"discount" description:"折扣" validate:"gt=0"`
		Desc         string  `json:"desc,optional" description:"描述"`
		EffectiveNum uint    `json:"effective_num" description:"有效期" validate:"gt=0"`
		Effective1   int     `json:"effective1" description:"有效期单位" validate:"eq=1|eq=2|eq=3"`
		Status       int     `json:"status" description:"状态" validate:"eq=1|eq=2"`
	}
	value := AddTest{
		Test1:        "",
		Count:        1,
		Price:        10,
		Discount:     10,
		Desc:         "1212",
		EffectiveNum: 10,
		Effective1:   1,
		Status:       1,
	}
	err := validate_.Struct(value)
	//获取翻译器
	// tran, _ := c.Get(TranslatorKey)
	// trans, _ := tran.(ut.Translator)
	// err := valid.Struct(params)
	//如果数据效验不通过，则将所有err以切片形式输出
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			t.Log(e.Translate(trans))
			//使用validator.ValidationErrors类型里的Translate方法进行翻译
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		t.Log(strings.Join(sliceErrs, ","))
	}
	t.Log(err)

}
