package bools_test

import (
	"fmt"

	"github.com/recallsong/go-utils/container/slice/bools"
)

func ExampleBools() {
	s := bools.Bools{}
	s.Append(false, false)
	s.Prepend(true, true, true, false, true)
	fmt.Println(s, s.Len())

	s1 := s.Copy()
	s1.Sort()
	fmt.Println(s1, s1.Len())

	s2 := s.Copy()
	s2.ReverseSort()
	fmt.Println(s2, s2.Len())

	// Output:
	// [true true true false true false false] 7
	// [false false false true true true true] 7
	// [true true true true false false false] 7
}
