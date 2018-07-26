package floats_test

import (
	"fmt"

	"github.com/recallsong/go-utils/container/slice/floats"
)

func ExampleFloat32s() {
	s := floats.Float32s{}
	s.Append(1.1, 1.2)
	s.Prepend(3.99, 3.98, 4.99, 5.69, 3.998)
	fmt.Println(s, s.Len())

	s1 := s.Copy()
	s1.Sort()
	fmt.Println(s1, s1.Len())

	s2 := s.Copy()
	s2.ReverseSort()
	fmt.Println(s2, s2.Len())

	// Output:
	// [3.99 3.98 4.99 5.69 3.998 1.1 1.2] 7
	// [1.1 1.2 3.98 3.99 3.998 4.99 5.69] 7
	// [5.69 4.99 3.998 3.99 3.98 1.2 1.1] 7
}

func ExampleFloa64s() {
	s := floats.Float64s{}
	s.Append(1.1, 1.2)
	s.Prepend(3.99, 3.98, 4.99, 5.69, 3.998)
	fmt.Println(s, s.Len())

	s1 := s.Copy()
	s1.Sort()
	fmt.Println(s1, s1.Len())

	s2 := s.Copy()
	s2.ReverseSort()
	fmt.Println(s2, s2.Len())

	// Output:
	// [3.99 3.98 4.99 5.69 3.998 1.1 1.2] 7
	// [1.1 1.2 3.98 3.99 3.998 4.99 5.69] 7
	// [5.69 4.99 3.998 3.99 3.98 1.2 1.1] 7
}
