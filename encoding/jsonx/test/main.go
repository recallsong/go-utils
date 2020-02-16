package main

import (
	"encoding/json"
	"fmt"

	"github.com/recallsong/go-utils/encoding/jsonx"
)

// JSONBytes .
type JSONBytes []byte

// MarshalJSON .
func (bs JSONBytes) MarshalJSON() ([]byte, error) {
	return bs, nil
}

// UnmarshalJSON .
func (bs *JSONBytes) UnmarshalJSON(b []byte) error {
	*bs = b
	return nil
}

// Unmarshal .
func (bs JSONBytes) Unmarshal(out interface{}) error {
	return json.Unmarshal([]byte(bs), out)
}

type Event struct {
	Name   string    `json:"name"`
	Params JSONBytes `json:"params"`
}

func main() {
	var e Event
	// e.Name = "xxxx"
	// e.Params = []byte("\"abc\"")
	// fmt.Println(jsonx.MarshalAndIntend(e))

	body := `{"name":"abc","params":"eeeddd"}`
	err := jsonx.Unmarshal(body, &e)
	fmt.Println(err)
	fmt.Println(e, string(e.Params))
	var xd string
	e.Params.Unmarshal(&xd)
	fmt.Println(xd)
}
