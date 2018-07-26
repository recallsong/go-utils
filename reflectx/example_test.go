package reflectx_test

import (
	"fmt"
	"testing"

	"github.com/recallsong/go-utils/reflectx"
)

func ExampleStructToMap() {
	type Inner struct {
		Age       int `api:"age"`
		NotOutput int
	}
	type Detail struct {
		Address   string `api:"address"`
		Inner     Inner  `api:"inner"`
		NotOutput int
	}
	type Person struct {
		Name      string  `api:"name"`
		Detail    Detail  `api:"detail"`
		DetailPtr *Detail `api:"detail_ptr"`
		Inner     Inner   `api:"inner"`
		NotOutput int     `api:"-"`
	}
	detail := Detail{Address: "xxx@xxx.com", Inner: Inner{Age: 18}, NotOutput: 100}
	s := Person{Name: "Song", Detail: detail, DetailPtr: &detail, Inner: Inner{Age: 20}}
	// 把Person转换成map
	m := reflectx.StructToMap(s, 0, "api")
	fmt.Println(m["detail"])
	fmt.Println(m["detail_ptr"])
	fmt.Println(m["inner"])
	// 把Person转换成map，递归1层，把Detail、DetailPtr、Inner也转换成map
	m = reflectx.StructToMap(s, 1, "api")
	fmt.Println(m["inner"])
	// Output:
	// {xxx@xxx.com {18 0} 100}
	// &{xxx@xxx.com {18 0} 100}
	// {20 0}
	// map[age:20]
}

var benchmark_str_data string = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
var benchmark_bytes_data = []byte(benchmark_str_data)

func BenchmarkBytesToString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = reflectx.BytesToString(benchmark_bytes_data)
	}
}

func BenchmarkBytesToString_cmpWithCopy(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = string(benchmark_bytes_data)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	b.ReportAllocs()
	str := benchmark_str_data
	for i := 0; i < b.N; i++ {
		_ = reflectx.StringToBytes(str)
	}
}

func BenchmarkStringToBytes_cmpWithCopy(b *testing.B) {
	b.ReportAllocs()
	str := benchmark_str_data
	for i := 0; i < b.N; i++ {
		_ = []byte(str)
	}
}
