package uintset

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Uint16Set 定义uint16集合
type Uint16Set map[uint16]struct{}

// String 将集合输出为字符串
func (s Uint16Set) String() string {
	bytes, err := json.Marshal(s.ToList())
	if err != nil {
		return fmt.Sprint(map[uint16]struct{}(s))
	}
	return string(bytes)
}

// Size 返回集合的大小
func (s Uint16Set) Size() int {
	return len(s)
}

// Add 向集合添加元素，如果已存在，则返回false, 否则返回true
func (s Uint16Set) Add(elem uint16) bool {
	_, ok := s[elem]
	s[elem] = struct{}{}
	return ok
}

// AddList 向集合中添加一组元素，自动过滤重复元素，返回实际添加的元素数量
func (s Uint16Set) AddList(elems ...uint16) int {
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
func (s Uint16Set) Remove(elem uint16) bool {
	_, ok := s[elem]
	delete(s, elem)
	return ok
}

// Clear 清空集合元素
func (s Uint16Set) Clear() {
	for k := range s {
		delete(s, k)
	}
}

// Contains 判断集合是否包含该元素
func (s Uint16Set) Contains(elem uint16) bool {
	_, ok := s[elem]
	return ok
}

// IsEmpty 判断集合是否为空
func (s Uint16Set) IsEmpty() bool {
	return len(s) == 0
}

// Equals 对两个集合进行深度比较，返回是否相等
func (s Uint16Set) Equals(set interface{}) bool {
	return reflect.DeepEqual(s, set)
}

// Copy 对集合进行浅拷贝
func (s Uint16Set) Copy() Uint16Set {
	cp := Uint16Set{}
	for k := range s {
		cp[k] = struct{}{}
	}
	return cp
}

// ToList 将集合转换为[]uint16
func (s Uint16Set) ToList() []uint16 {
	keys := make([]uint16, len(s))
	i := 0
	for k := range s {
		keys[i] = k
		i = i + 1
	}
	return keys
}

// Foreach 遍历集合，如果f返回 false ，则终止遍历
func (s Uint16Set) Foreach(f func(v uint16) bool) {
	for k := range s {
		if !f(k) {
			return
		}
	}
}

// AddSet 集合并集
func (s Uint16Set) AddSet(set Uint16Set) {
	for elem := range set {
		s[elem] = struct{}{}
	}
}

// removeSet 集合差集
func (s Uint16Set) RemoveSet(set Uint16Set) {
	for elem := range s {
		if _, ok := set[elem]; ok {
			delete(s, elem)
		}
	}
}

// retainAll 集合交集
func (s Uint16Set) RetainSet(set Uint16Set) {
	for elem := range s {
		if _, ok := set[elem]; !ok {
			delete(s, elem)
		}
	}
}
