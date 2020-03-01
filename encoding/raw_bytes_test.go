package encoding

import (
	"encoding/json"
	"fmt"
)

func Example_rawbyte_json() {
	var obj struct {
		Name   string   `json:"name"`
		Params RawBytes `json:"params"`
	}
	body := `{"name":"abc","params":"eeeddd"}`
	err := json.Unmarshal([]byte(body), &obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(obj, string(obj.Params))
	// Output:
	// {abc [34 101 101 101 100 100 100 34]} "eeeddd"
}
