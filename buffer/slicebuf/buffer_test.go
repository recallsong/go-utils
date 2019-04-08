package slicebuf

import (
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	buffer := New(10)
	buf := buffer.Require(1)
	buf[0] = 1
	if buffer.Len() != 1 || len(buffer.Bytes()) != 1 {
		t.Error(`buffer.Len() != 1 || len(buffer.Bytes()) != 1`)
	}
	if buffer.Bytes()[0] != 1 {
		t.Error(`buffer.Bytes()[0] != 1 `)
	}
	buf = buffer.Require(999)
	buf[998] = 2
	if buffer.Len() != 1000 || len(buffer.Bytes()) != 1000 {
		t.Error(`buffer.Len() != 1000 || len(buffer.Bytes()) != 1000`)
	}
	if buffer.Bytes()[999] != 2 {
		t.Error(`buffer.Bytes()[999] != 2`)
	}
}

func ExampleBuffer() {
	buffer := New(10)
	buf := buffer.Require(4)
	buf[0] = 1
	buf[1] = 2
	buf[2] = 3
	buf[3] = 4
	fmt.Println(buffer.Bytes())

	// Output:
	// [1 2 3 4]
}
