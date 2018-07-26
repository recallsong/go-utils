package set_test

import (
	"fmt"

	"github.com/recallsong/go-utils/container/set"
	"github.com/recallsong/go-utils/container/slice/ints"
	"github.com/recallsong/go-utils/container/slice/strings"
	"github.com/recallsong/go-utils/conv"
)

func ExampleSet() {
	s := set.Set{}
	// 向集合添加元素列表
	s.AddList(1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9)
	// 向集合添加单个元素
	s.Add(10)
	// 排序并打印
	fmt.Println(ints.Ints(conv.Interfaces(s.ToList()).ToInts()).Sort())

	// 创建保存Persion结构体的集合
	type Persion struct {
		Name string
	}
	ps := set.Set{}
	ps.Add(Persion{Name: "Song"})
	ps.Add(Persion{Name: "Hello"})
	ps.Add(Persion{Name: "Song"})
	var plist []Persion
	// 将集合元素输出为[]Persion
	ps.Dump(&plist)
	fmt.Println(len(plist))

	// 集合运算
	s1 := set.Set{}
	s1.AddList("A", "B", "C")
	s2 := set.Set{}
	s2.AddList("A", "D")
	// 并集
	s3 := s1.Copy()
	s3.AddSet(s2)
	fmt.Println(strings.Strings(conv.Interfaces(s3.ToList()).ToStrings()).Sort())
	// 差集
	s3 = s1.Copy()
	s3.RemoveSet(s2)
	fmt.Println(strings.Strings(conv.Interfaces(s3.ToList()).ToStrings()).Sort())
	// 交集
	s3 = s1.Copy()
	s3.RetainSet(s2)
	fmt.Println(strings.Strings(conv.Interfaces(s3.ToList()).ToStrings()).Sort())

	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
	// 2
	// [A B C D]
	// [B C]
	// [A]
}
