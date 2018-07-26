package errorx_test

import (
	"fmt"

	"github.com/recallsong/go-utils/errorx"
)

var _ error = (errorx.MultiError)(nil)
var _ error = (errorx.StringError)("")
var _ error = (errorx.CodedError)(nil)
var _ error = (errorx.TracedError)(nil)

func ExampleStringError() {
	var err error = errorx.New("this is error message")
	fmt.Println(err)
	// Output:
	// this is error message
}

func ExampleMultiError() {
	err0 := errorx.New("error 0")
	err1 := errorx.New("error 1")
	err2 := errorx.New("error 2")
	errs := errorx.NewMultiError(err0, err1, err2)

	fmt.Println(errs)
	// Output:
	// 3 error(s) occurred:
	// * error 0
	// * error 1
	// * error 2
}

func ExampleCodedError() {
	err := errorx.NewCodedError(400, "bad request")
	fmt.Println(err)
	// Output:
	// 400, bad request
}

func ExampleTracedError() {
	err := errorx.NewTracedError("this is error message")
	fmt.Println(err)
	// Output:
	// this is error message :
	// * [example_test.go:43]	errorx_test.ExampleTracedError
	// * [example.go:122]	testing.runExample
	// * [example.go:46]	testing.runExamples
	// * [testing.go:922]	testing.(*M).Run
	// * [_testmain.go:56]	main.main
}

func ExampleTracedError_2() {
	var err error = errorx.NewTracedError(nil)
	fmt.Println(err == nil)
	err = errorx.NewTracedError(errorx.NewCodedError(400, "bad request"))
	fmt.Println(err)
	// Output:
	// true
	// 400, bad request :
	// * [example_test.go:57]	errorx_test.ExampleTracedError_2
	// * [example.go:122]	testing.runExample
	// * [example.go:46]	testing.runExamples
	// * [testing.go:922]	testing.(*M).Run
	// * [_testmain.go:56]	main.main
}
