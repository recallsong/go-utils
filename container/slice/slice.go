package slice

import (
	"fmt"
	"reflect"
	"sort"
)

// Sort 方便对切片进行排序
func Sort(slice interface{}, less func(i, j int) bool) {
	sort.Sort(SortInterface(slice, less))
}

// SortInterface 返回可以排序的sort.Interface接口
func SortInterface(slice interface{}, less func(i, j int) bool) sort.Interface {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Errorf("slice.Sort called with non-slice value of type %T", slice))
	}
	value := reflect.ValueOf(slice)
	return &sorter{
		length: sv.Len(),
		less:   less,
		swapper: &swapper{
			slice: value,
			tmp:   reflect.New(value.Type().Elem()).Elem(),
		},
	}
}

// Sort2 方便对切片进行排序
func Sort2(slice interface{}, less func(i, j int) bool, swap func(i, j int)) {
	sort.Sort(SortInterface2(slice, less, swap))
}

// SortInterface2 返回可以排序的sort.Interface接口
func SortInterface2(slice interface{}, less func(i, j int) bool, swap func(i, j int)) sort.Interface {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Errorf("slice.Sort called with non-slice value of type %T", slice))
	}
	return &funcs{
		length: sv.Len(),
		less:   less,
		swap:   swap,
	}
}

type sorter struct {
	length int
	less   func(i, j int) bool
	*swapper
}

func (f *sorter) Len() int           { return f.length }
func (f *sorter) Less(i, j int) bool { return f.less(i, j) }

type swapper struct {
	slice reflect.Value
	tmp   reflect.Value
}

func (s *swapper) Swap(i, j int) {
	v1 := s.slice.Index(i)
	v2 := s.slice.Index(j)
	s.tmp.Set(v1)
	v1.Set(v2)
	v2.Set(s.tmp)
}

type funcs struct {
	length int
	less   func(i, j int) bool
	swap   func(i, j int)
}

func (f *funcs) Len() int           { return f.length }
func (f *funcs) Less(i, j int) bool { return f.less(i, j) }
func (f *funcs) Swap(i, j int)      { f.swap(i, j) }
