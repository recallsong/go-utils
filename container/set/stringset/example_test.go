package stringset_test

import (
	"fmt"

	"github.com/recallsong/go-utils/container/set/stringset"
	"github.com/recallsong/go-utils/container/slice/strings"
)

func ExampleStringSet() {
	set := stringset.StringSet{}
	// 向集合添加多个元素
	set.AddList("a", "b", "c", "d", "c", "c", "d")
	// 向集合添加单个元素
	set.Add("e")
	// 排序并打印
	fmt.Println(strings.Strings(set.ToList()).Sort())

	// 集合运算
	s1 := stringset.StringSet{}
	s1.AddList("a", "b", "c")
	s2 := stringset.StringSet{}
	s2.AddList("c", "d", "e")
	// 并集
	s3 := s1.Copy()
	s3.AddSet(s2)
	fmt.Println(strings.Strings(s3.ToList()).Sort())
	// 差集
	s3 = s1.Copy()
	s3.RemoveSet(s2)
	fmt.Println(strings.Strings(s3.ToList()).Sort())
	// 交集
	s3 = s1.Copy()
	s3.RetainSet(s2)
	fmt.Println(strings.Strings(s3.ToList()).Sort())

	// Output:
	// [a b c d e]
	// [a b c d e]
	// [a b]
	// [c]
}
