package ioutil

import (
	"io"

	"github.com/recallsong/go-utils/errorx"
)

// CloseMulti 关闭多个io
func CloseMulti(cs ...io.Closer) error {
	var errs errorx.Errors
	var err error
	for _, c := range cs {
		err = c.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// MerageCloser 合并多个io.Closer为一个
func MerageCloser(cs ...io.Closer) MultiCloser {
	return MultiCloser(cs)
}

// MultiCloser 多个io.Closer的集合
type MultiCloser []io.Closer

func (mc MultiCloser) Close() error {
	var errs errorx.Errors
	var err error
	for _, c := range mc {
		err = c.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (mc *MultiCloser) Append(c io.Closer) {
	*mc = append(*mc, c)
}
