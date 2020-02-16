package jsonx

import (
	"bytes"
	"encoding/json"
)

// Marshal 将interface{}类型的值转换成json字符串，失败将产生panic
func Marshal(value interface{}) string {
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// Unmarshal .
func Unmarshal(str string, out interface{}) error {
	return json.Unmarshal([]byte(str), out)
}

// ToMap 解析字符串到结构体或者map[string]interface{}中，out为指针类型
func ToMap(str string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &out)
	if err != nil {
		return nil, err
	}
	return out, err
}

// MarshalAndIntend .
func MarshalAndIntend(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "\t")
	if err != nil {
		panic(err)
	}
	return out.String()
}

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
