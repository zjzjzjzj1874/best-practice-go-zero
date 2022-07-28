package helper

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate 参数校验器
func Validate() *validator.Validate {
	return validate
}
