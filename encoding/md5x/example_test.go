package md5x_test

import (
	"fmt"
	"os"

	"github.com/recallsong/go-utils/encoding/md5x"
)

func Example_sum() {
	str := md5x.Sum([]byte("123")).String()
	fmt.Println(str)
	// Output:
	// 202cb962ac59075b964b07152d234b70
}

func Example_sumString() {
	str := md5x.SumString("123").String()
	fmt.Println(str)
	// Output:
	// 202cb962ac59075b964b07152d234b70
}

func Example_fileByPath() {
	val, err := md5x.SumFileByPath("md5x.go")
	fmt.Println(val.String(), err)
	// Output:
	// 25d81829d5e7219164dc00a297fc579b <nil>
}

func Example_file() {
	file, err := os.Open("md5x.go")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	val, err := md5x.SumFile(file)
	fmt.Println(val.String(), err)
	// Output:
	// 25d81829d5e7219164dc00a297fc579b <nil>
}

func Example_short() {
	val, err := md5x.SumFileByPath("md5x.go")
	fmt.Println(val.String16(), err)
	// Output:
	// 081a33d327980264 <nil>
}
