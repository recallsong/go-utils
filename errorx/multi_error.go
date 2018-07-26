package errorx

import (
	"bytes"
	"fmt"
)

// MultiError 多个错误的集合类型
type MultiError interface {
	error
	Num() int // 返回错误的数量
	MaybeUnwrap() error
}

// Errors is a slice of errors implementing the error interface.
type Errors []error

// NewMultiError 创建错误集合
func NewMultiError(errs ...error) MultiError {
	merr := make(Errors, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			merr = append(merr, err)
		}
	}
	return merr
}

func (errs Errors) Error() string {
	if len(errs) == 0 {
		return ""
	} else if len(errs) == 1 {
		return errs[0].Error()
	}
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%d error(s) occurred:", len(errs))
	for _, err := range errs {
		fmt.Fprintf(buf, "\n* %s", err)
	}
	return buf.String()
}

func (errs Errors) Num() int {
	return len(errs)
}

func (errs *Errors) Append(err error) {
	*errs = append(*errs, err)
	return
}

// MaybeUnwrap returns nil if len(errs) is 0. It returns the first and only
// contained error as error if len(errs is 1). In all other cases, it returns
// the Errors directly. This is helpful for returning a Errors in a way
// that only uses the Errors if needed.
func (errs Errors) MaybeUnwrap() error {
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errs
	}
}
