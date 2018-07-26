package floatset_test

import (
	"fmt"

	"github.com/recallsong/go-utils/container/set/floatset"
	"github.com/recallsong/go-utils/container/slice/floats"
)

func ExampleFloat32Set() {
	set := floatset.Float32Set{}
	// 向集合添加多个元素
	set.AddList(1, 2, 3, 4, 5, 6, 5, 5, 5, 6, 6, 6)
	// 向集合添加单个元素
	set.Add(7)
	// 排序并打印
	fmt.Println(floats.Float32s(set.ToList()).Sort())

	// 集合运算
	s1 := floatset.Float32Set{}
	s1.AddList(1, 2, 3)
	s2 := floatset.Float32Set{}
	s2.AddList(3, 4, 5)
	// 并集
	s3 := s1.Copy()
	s3.AddSet(s2)
	fmt.Println(floats.Float32s(s3.ToList()).Sort())
	// 差集
	s3 = s1.Copy()
	s3.RemoveSet(s2)
	fmt.Println(floats.Float32s(s3.ToList()).Sort())
	// 交集
	s3 = s1.Copy()
	s3.RetainSet(s2)
	fmt.Println(floats.Float32s(s3.ToList()).Sort())

	// Output:
	// [1 2 3 4 5 6 7]
	// [1 2 3 4 5]
	// [1 2]
	// [3]
}

func ExampleFloat64Set() {
	set := floatset.Float64Set{}
	// 向集合添加多个元素
	set.AddList(1, 2, 3, 4, 5, 6, 5, 5, 5, 6, 6, 6)
	// 向集合添加单个元素
	set.Add(7)
	// 排序并打印
	fmt.Println(floats.Float64s(set.ToList()).Sort())

	// 集合运算
	s1 := floatset.Float64Set{}
	s1.AddList(1, 2, 3)
	s2 := floatset.Float64Set{}
	s2.AddList(3, 4, 5)
	// 并集
	s3 := s1.Copy()
	s3.AddSet(s2)
	fmt.Println(floats.Float64s(s3.ToList()).Sort())
	// 差集
	s3 = s1.Copy()
	s3.RemoveSet(s2)
	fmt.Println(floats.Float64s(s3.ToList()).Sort())
	// 交集
	s3 = s1.Copy()
	s3.RetainSet(s2)
	fmt.Println(floats.Float64s(s3.ToList()).Sort())

	// Output:
	// [1 2 3 4 5 6 7]
	// [1 2 3 4 5]
	// [1 2]
	// [3]
}
