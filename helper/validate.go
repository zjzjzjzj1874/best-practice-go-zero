package helper

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
)

var validate *validator.Validate

func init() {
	//设置支持语言
	chinese := zh.New()
	//设置国际化翻译器
	uni := ut.New(chinese, chinese)
	//设置翻译语言
	trans, found := uni.GetTranslator("zh")
	if !found {
		log.Fatalf("no zh Translator")
	}
	validate = validator.New()

	// 注册获取字段名的方法
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		tagValue, ok := field.Tag.Lookup("description")
		if !ok {
			// 获取不到返回字段名
			return field.Name
		}
		return tagValue
	})
	//注册默认翻译
	if err := zhtrans.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatalf("validator register zh translation failed [err:%s]", err.Error())
	}
}

// Validate 参数校验器
func Validate() *validator.Validate {
	return validate
}
