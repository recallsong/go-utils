package errorx

import "fmt"

// CodedError 具有code的错误类型
type CodedError interface {
	error
	Code() int       // 返回错误代码
	Message() string // 返回错误消息
}

type codedError struct {
	code    int
	message string
}

// NewCodedError 创建一个具有编码的错误
func NewCodedError(code int, msg string) CodedError {
	return &codedError{code: code, message: msg}
}

func (err *codedError) Error() string {
	return fmt.Sprintf("%d, %s", err.code, err.message)
}

func (err *codedError) Code() int {
	return err.code
}

func (err *codedError) Message() string {
	return err.message
}
