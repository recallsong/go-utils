package emitter_test

import (
	"fmt"

	"github.com/recallsong/go-utils/emitter"
)

func Example_fileByPath() {
	emit := emitter.New()
	w1 := emit.Watch("connect")
	w1.Callback(func(e *emitter.Event) {
		fmt.Println("w1", e.Name)
	})

	w2 := emit.Watch("connect")
	w2.Callback(func(e *emitter.Event) {
		fmt.Println("w2", e.Name)
	})

	w3 := emit.Watch("disconnect")
	w3.Callback(func(e *emitter.Event) {
		fmt.Println("w3", e.Name)
	})

	w1.Close()

	w4 := emit.Watch("connect")
	w4.Callback(func(e *emitter.Event) {
		fmt.Println("w4", e.Name)
	})

	emit.Emit(&emitter.Event{
		Name: "connect",
		Data: 123,
	})
	emit.Emit(&emitter.Event{
		Name: "disconnect",
		Data: 123,
	})
	// Output:
	// w4 connect
	// w2 connect
	// w3 disconnect
}
