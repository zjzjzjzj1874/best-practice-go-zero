package errors

import (
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ErrorHandler 定义非Restful错误处理函数
func ErrorHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case validator.ValidationErrors: // 参数校验错误
		return http.StatusOK, NewStatusError(StatusBadRequestError).WithMsg(err.Error()).Response()
	case *StatusError:
		return e.Code(), e.Response()
	default: // 未知错误则为服务器内部错误
		return http.StatusOK, NewStatusError(InternalServerError).WithMsg(err.Error()).Response()
	}
}

// ErrorRestfulHandler 定义Restful错误处理函数
func ErrorRestfulHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case validator.ValidationErrors: // 参数校验错误
		return http.StatusBadRequest, e.Error()
	case *StatusError:
		return e.Code(), e.Error()
	default: // 未知错误则为服务器内部错误
		return http.StatusInternalServerError, e.Error()
	}
}
