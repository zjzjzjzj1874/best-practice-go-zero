package errors

import "net/http"

type Status int

// 接口状态错误
type StatusError struct {
	code Status
	msg  string
}

type StatusErrorResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (v StatusError) Response() *StatusErrorResponse {
	return &StatusErrorResponse{
		Code: v.Code(),
		Msg:  v.Error(),
	}
}

func ErrBadRequest(msg string) StatusError {
	return NewStatusError(StatusBadRequestError).WithMsg(msg)
}

func NewUnauthorizedErr(msg string) StatusError {
	return NewStatusError(StatusUnauthorized).WithMsg(msg)
}

func NewStatusError(code Status) StatusError {
	return StatusError{code: code}
}

func (v StatusError) Error() string {
	if v.msg == "" {
		return v.code.Msg()
	}
	return v.msg
}

func (v StatusError) Code() int {
	return int(v.code) / 1e3
}

func (v StatusError) WithMsg(message string) StatusError {
	v.msg = message
	return v
}

// 400
const (
	StatusBadRequestError Status = http.StatusBadRequest*1e3 + iota + 1
)

// 401
const (
	StatusUnauthorized Status = http.StatusUnauthorized*1e3 + iota + 1
	// @errTalk 无访问权限
	AccessNotAllowedError
)

// 403
const (
	// Forbidden
	ForbiddenError Status = http.StatusForbidden*1e3 + iota + 1
	// @errTalk 无效的请求来源
	InvalidRequestOrigin
)

// NotFoundError 404
const (
	NotFoundError Status = http.StatusNotFound*1e3 + iota + 1
)

// ConflictError 409
const (
	ConflictError Status = http.StatusConflict*1e3 + iota + 1
)

// InternalServerError 500
const (
	InternalServerError Status = http.StatusInternalServerError*1e3 + iota + 1
)

func (v Status) Msg() string {
	switch v {
	case StatusBadRequestError:
		return "错误请求"
	case StatusUnauthorized:
		return "未授权"
	case AccessNotAllowedError:
		return "无访问权限"
	case ForbiddenError:
		return "Forbidden"
	case InvalidRequestOrigin:
		return "无效的请求来源"
	case NotFoundError:
		return "未找到"
	case ConflictError:
		return "请求冲突"
	case InternalServerError:
		return "服务内部错误"
	}
	return "未知错误"
}
