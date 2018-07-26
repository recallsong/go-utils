package uintset_test

import (
	"fmt"

	"github.com/recallsong/go-utils/container/set/uintset"
	"github.com/recallsong/go-utils/container/slice/uints"
)

func ExampleUintSet() {
	set := uintset.UintSet{}
	// 向集合添加多个元素
	set.AddList(1, 2, 3, 4, 5, 6, 5, 5, 5, 6, 6, 6)
	// 向集合添加单个元素
	set.Add(7)
	// 排序并打印
	fmt.Println(uints.Uints(set.ToList()).Sort())

	// 集合运算
	s1 := uintset.UintSet{}
	s1.AddList(1, 2, 3)
	s2 := uintset.UintSet{}
	s2.AddList(3, 4, 5)
	// 并集
	s3 := s1.Copy()
	s3.AddSet(s2)
	fmt.Println(uints.Uints(s3.ToList()).Sort())
	// 差集
	s3 = s1.Copy()
	s3.RemoveSet(s2)
	fmt.Println(uints.Uints(s3.ToList()).Sort())
	// 交集
	s3 = s1.Copy()
	s3.RetainSet(s2)
	fmt.Println(uints.Uints(s3.ToList()).Sort())

	// Output:
	// [1 2 3 4 5 6 7]
	// [1 2 3 4 5]
	// [1 2]
	// [3]
}
