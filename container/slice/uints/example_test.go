package uints_test

import (
	"fmt"

	"github.com/recallsong/go-utils/container/slice/uints"
)

func ExampleUints() {
	s := uints.Uints{}
	s.Append(1, 2)
	s.Prepend(4, 3, 2, 1, 0)
	s.Insert(1, 3, 4, 5, 6)
	fmt.Println(s, s.Len())

	s1 := s.Copy()
	removed := s1.Remove(3, 6)
	fmt.Println(s1, s1.Len(), removed)
	// Output:
	// [4 3 4 5 6 3 2 1 0 1 2] 11
	// [4 3 4 1 2] 5 6
}

func ExampleUints_order() {
	s := uints.Uints{}
	s.Append(9, 1, 8, 2, 7, 3, 6, 4, 5)
	s.Sort()
	fmt.Println(s)
	s.ReverseSort()
	fmt.Println(s)

	s.Shuffle() //随机打乱数据
	// fmt.Println(s)

	// Output:
	// [1 2 3 4 5 6 7 8 9]
	// [9 8 7 6 5 4 3 2 1]
}
