package reflectx

import (
	"reflect"
)

// StructToMap 将结构体转为map类型,
// 参数：
// data为结构体地址变量，
// depth为转换的深度，0表示不递归转换struct, <0表示无限递归转换，>0表示递归转换的深度，
// tag为可选，不指定则全部字段输出，否则将指定tag的字段输出到map
func StructToMap(data interface{}, depth int, tag ...string) map[string]interface{} {
	m := make(map[string]interface{})
	values := reflect.ValueOf(data)
	types := reflect.TypeOf(data)
	for types.Kind() == reflect.Ptr {
		values = values.Elem()
		types = types.Elem()
	}
	num := types.NumField()
	depth = depth - 1
	if len(tag) <= 0 || tag[0] == "" {
		if depth == -1 {
			for i := 0; i < num; i++ {
				t := types.Field(i)
				v := values.Field(i)
				if v.CanInterface() {
					m[t.Name] = v.Interface()
				}
			}
		} else {
			for i := 0; i < num; i++ {
				t := types.Field(i)
				v := values.Field(i)
				v_struct := v
				v_struct_ptr := v
				for v_struct.Kind() == reflect.Ptr {
					v_struct_ptr = v_struct
					v_struct = v_struct.Elem()
				}
				if v_struct.Kind() == reflect.Struct && v_struct_ptr.CanInterface() {
					m[t.Name] = StructToMap(v_struct_ptr.Interface(), depth, tag[0])
				} else {
					if v.CanInterface() {
						m[t.Name] = v.Interface()
					}
				}
			}
		}
	} else {
		tagName := tag[0]
		if depth == -1 {
			for i := 0; i < num; i++ {
				t := types.Field(i)
				v := values.Field(i)
				tagVal := t.Tag.Get(tagName)
				if v.CanInterface() && tagVal != "" && tagVal != "-" {
					m[tagVal] = v.Interface()
				}
			}
		} else {
			for i := 0; i < num; i++ {
				t := types.Field(i)
				v := values.Field(i)
				tagVal := t.Tag.Get(tagName)
				if tagVal != "" && tagVal != "-" {
					v_struct := v
					v_struct_ptr := v
					for v_struct.Kind() == reflect.Ptr {
						v_struct_ptr = v_struct
						v_struct = v_struct.Elem()
					}
					if v_struct.Kind() == reflect.Struct && v_struct_ptr.CanInterface() {
						m[tagVal] = StructToMap(v_struct_ptr.Interface(), depth, tag[0])
						continue
					}
					if v.CanInterface() {
						m[tagVal] = v.Interface()
					}
				}
			}
		}
	}
	return m
}
