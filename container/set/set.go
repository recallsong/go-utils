package set

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrTypeCast = errors.New("failed to type cast")
)

// Set 定义集合
type Set map[interface{}]struct{}

// String 将集合输出为字符串
func (s Set) String() string {
	bytes, err := json.Marshal(s.ToList())
	if err != nil {
		return fmt.Sprint(map[interface{}]struct{}(s))
	}
	return string(bytes)
}

// Size 返回集合的大小
func (s Set) Size() int {
	return len(s)
}

// Add 向集合添加元素，如果已存在，则返回false, 否则返回true
func (s Set) Add(elem interface{}) bool {
	_, ok := s[elem]
	s[elem] = struct{}{}
	return ok
}

// AddList 向集合中添加一组元素，自动过滤重复元素，返回实际添加的元素数量
func (s Set) AddList(elems ...interface{}) int {
	new_elem := 0
	for _, elem := range elems {
		if _, ok := s[elem]; !ok {
			new_elem = new_elem + 1
		}
		s[elem] = struct{}{}
	}
	return new_elem
}

// Remove 从集合中删除元素，如果不存在，则返回false, 否则返回true
func (s Set) Remove(elem interface{}) bool {
	_, ok := s[elem]
	delete(s, elem)
	return ok
}

// Clear 清空集合元素
func (s Set) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Contains 判断集合是否包含该元素
func (s Set) Contains(elem interface{}) bool {
	_, ok := s[elem]
	return ok
}

// IsEmpty 判断集合是否为空
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Equals 对两个集合进行深度比较，返回是否相等
func (s Set) Equals(set interface{}) bool {
	return reflect.DeepEqual(s, set)
}

// Copy 对集合进行浅拷贝
func (s Set) Copy() Set {
	cp := Set{}
	for k := range s {
		cp[k] = struct{}{}
	}
	return cp
}

// Dump 将集合输出到list，失败则返回错误信息
func (s Set) Dump(list interface{}) (err error) {
	typ := reflect.TypeOf(list)
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Slice {
		return errors.New("Set.Dump args must be a slice pointer")
	}
	defer func() {
		e := recover()
		if e != nil {
			switch v := e.(type) {
			case error:
				err = v
			default:
				err = fmt.Errorf("%v", v)
			}
		}
	}()
	vlist := reflect.MakeSlice(typ.Elem(), 0, len(s))
	for k := range s {
		vlist = reflect.Append(vlist, reflect.ValueOf(k))
	}
	reflect.ValueOf(list).Elem().Set(vlist)
	return err
}

// ToList 将集合转换为[]interface{}
func (s Set) ToList() []interface{} {
	keys := make([]interface{}, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i = i + 1
	}
	return keys
}

// Foreach 遍历集合，如果f返回 false ，则终止遍历
func (s Set) Foreach(f func(v interface{}) bool) {
	for k := range s {
		if !f(k) {
			return
		}
	}
}

// AddSet 集合并集
func (s Set) AddSet(set Set) {
	for elem := range set {
		s[elem] = struct{}{}
	}
}

// removeSet 集合差集
func (s Set) RemoveSet(set Set) {
	for elem := range s {
		if _, ok := set[elem]; ok {
			delete(s, elem)
		}
	}
}

// retainAll 集合交集
func (s Set) RetainSet(set Set) {
	for elem := range s {
		if _, ok := set[elem]; !ok {
			delete(s, elem)
		}
	}
}
