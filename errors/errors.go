package errors

import "strings"

// ErrorType 代表错误类型。
type ErrorType string

// 错误类型常量。
const (
	//数据库连接错误
	ERROR_TYPE_DBCONNECTION ErrorType = "数据库连接失败"
)

type CommonError interface {
	// Type 用于获得错误的类型。
	Type() ErrorType
	// Error 用于获得错误提示信息。
	Error() string
}

type commonError struct {
	// errType 代表错误的类型。
	errType ErrorType
	// errMsg 代表错误的提示信息。
	errMsg string
	// fullErrMsg 代表完整的错误提示信息。
	fullErrMsg string
}

func (c *commonError) Type() ErrorType {
	return c.errType
}
func (c *commonError) Error() string {
	return c.errMsg
}

func NewCommonError(errType ErrorType, errMsg string) CommonError {
	return &commonError{
		errType: errType,
		errMsg:  strings.TrimSpace(errMsg),
	}
}
