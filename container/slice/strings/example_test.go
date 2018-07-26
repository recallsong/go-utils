package strings_test

import (
	"fmt"

	"github.com/recallsong/go-utils/container/slice/strings"
)

func ExampleStrings() {
	s := strings.Strings{}
	s.Append("abc", "abcd")
	s.Prepend("123", "456", "1234")
	s.Insert(1, "efg", "12")
	fmt.Println(s, s.Len())

	s1 := s.Copy()
	removed := s1.Remove(1, 1)
	fmt.Println(s1, s1.Len(), removed)
	// Output:
	// [123 efg 12 456 1234 abc abcd] 7
	// [123 12 456 1234 abc abcd] 6 1
}

func ExampleStrings_order() {
	s := strings.Strings{}
	s.Append("5", "4", "6", "3", "1", "2")
	s.Sort()
	fmt.Println(s)
	s.ReverseSort()
	fmt.Println(s)

	s.Shuffle() //随机打乱数据
	// fmt.Println(s)

	// Output:
	// [1 2 3 4 5 6]
	// [6 5 4 3 2 1]
}
