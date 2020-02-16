package main

import (
	"fmt"

	"github.com/recallsong/go-utils/emitter"
)

func main() {
	emit := emitter.New()
	w1 := emit.Watch("connect")
	w1.Callback(func(e *emitter.Event) {
		fmt.Println("w1", e.Name)
	})
	fmt.Println(emit)

	w2 := emit.Watch("connect")
	w2.Callback(func(e *emitter.Event) {
		fmt.Println("w2", e.Name)
	})
	fmt.Println(emit)

	w3 := emit.Watch("disconnect")
	w3.Callback(func(e *emitter.Event) {
		fmt.Println("w3", e.Name)
	})
	fmt.Println(emit)

	w1.Close()
	fmt.Println(emit)

	w4 := emit.Watch("connect")
	w4.Callback(func(e *emitter.Event) {
		fmt.Println("w4", e.Name)
	})
	fmt.Println(emit)

	emit.Emit(&emitter.Event{
		Name: "connect",
		Data: 123,
	})
	emit.Emit(&emitter.Event{
		Name: "disconnect",
		Data: 123,
	})
}
