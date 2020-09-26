package conv_test

import (
	"fmt"

	"github.com/recallsong/go-utils/conv"
)

func Example_parse() {
	fmt.Println(conv.ParseInt("100", -1))
	fmt.Println(conv.ParseInt("error", -1))
	fmt.Println(conv.ParseInt8("108", -1))
	fmt.Println(conv.ParseInt16("116", -1))
	fmt.Println(conv.ParseInt32("133", -1))
	fmt.Println(conv.ParseInt64("164", -1))

	fmt.Println(conv.ParseUint("100", 0))
	fmt.Println(conv.ParseUint8("108", 0))
	fmt.Println(conv.ParseUint16("116", 0))
	fmt.Println(conv.ParseUint32("132", 0))
	fmt.Println(conv.ParseUint64("164", 0))

	fmt.Println(conv.ParseFloat32("100.999", -1))
	fmt.Println(conv.ParseFloat64("104.666", -1))

	fmt.Println(conv.ParseBool("true", true))
	fmt.Println(conv.ParseBool("1", true))
	fmt.Println(conv.ParseBool("0", true))
	fmt.Println(conv.ParseBool("False", true))
	fmt.Println(conv.ParseBool("FALSE", true))

	// Output:
	// 100
	// -1
	// 108
	// 116
	// 133
	// 164
	// 100
	// 108
	// 116
	// 132
	// 164
	// 100.999
	// 104.666
	// true
	// true
	// false
	// false
	// false
}

func Example_convert() {
	var val interface{}
	val = 123
	fmt.Println(conv.ToInt(val, -1))
	fmt.Println(conv.ToInt64(val, -1))
	fmt.Println(conv.ToUint(val, 0))
	fmt.Println(conv.ToUint64(val, 0))
	fmt.Println(conv.ToFloat32(val, -1))
	fmt.Println(conv.ToFloat64(val, -1))
	fmt.Println(conv.ToBool(val, false))
	val = true
	fmt.Println(conv.ToBool(val, false))
	fmt.Println(conv.ToString(val))

	// Output:
	// 123
	// 123
	// 123
	// 123
	// 123
	// 123
	// false
	// true
	// true
}

func Example_slice() {
	is := []int{1, 2, 3, 4, 5}
	var inters []interface{} = conv.Ints(is).ToInterfaces()
	fmt.Println(inters)
	fs := []float32{1, 2, 3, 4, 5}
	var ints []int = conv.Float32s(fs).ToInts()
	fmt.Println(ints)

	// Output:
	// [1 2 3 4 5]
	// [1 2 3 4 5]
}

func Example_formatFloat() {
	// conv.CleanFloat("")
	fmt.Println(conv.FormatFloat64(1.25678, 2))
	fmt.Println(conv.FormatFloat64(1.02567, 2))
	fmt.Println(conv.FormatFloat64(1.0000, 2))

	// Output:
	// 1.26
	// 1.03
	// 1
}
