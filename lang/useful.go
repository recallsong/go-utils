package lang

import "reflect"

// IsNil 判断v是否为nil
func IsNil(v interface{}) bool {
	return (v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()))
}
