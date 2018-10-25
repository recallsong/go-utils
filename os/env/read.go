package env

import (
	"os"
	"reflect"
	"strings"
)

func Unmarshal(out interface{}) {
	typ := reflect.TypeOf(out)
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() == reflect.Ptr {
		panic("env.Unmarshal parameter must be one-level pointer")
	}
	typ = typ.Elem()
	val := reflect.ValueOf(out).Elem()
	for i, num := 0, typ.NumField(); i < num; i++ {
		field := typ.Field(i)
		key := field.Tag.Get("env")
		if len(key) == 0 {
			key = field.Name
		}
		key = strings.ToUpper(key)
		value := os.Getenv(key)
		if len(value) > 0 {
			v := val.Field(i)
			v.Set(reflect.ValueOf(value))
		}
	}
}
