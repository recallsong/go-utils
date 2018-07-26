package ints

import (
	"math/rand"
	"sort"
)

// Int64s 定义方便操作[]int64的类型
type Int64s []int64

// Len 返回slice长度
func (s Int64s) Len() int {
	return len(s)
}

// Less 比较两个位置上的数据
func (s Int64s) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap 交换两个位置上的数据
func (s Int64s) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Cap 返回slice容量
func (s Int64s) Cap() int {
	return cap(s)
}

// Copy 拷贝一份数据
func (s Int64s) Copy() Int64s {
	cp := make([]int64, len(s))
	copy(cp, s)
	return cp
}

// Sort 对切片进行排序
func (s Int64s) Sort() Int64s {
	sort.Sort(s)
	return s
}

func (s Int64s) ReverseSort() Int64s {
	sort.Sort(sort.Reverse(s))
	return s
}

// Shuffle 随机打乱slice数据
func (s Int64s) Shuffle() Int64s {
	ln := len(s)
	for i := 0; i < ln; i++ {
		j := rand.Intn(i + 1)
		s[j], s[i] = s[i], s[j]
	}
	return s
}

// Append 追加数据
func (s *Int64s) Append(data ...int64) {
	*s = append(*s, data...)
}

// Prepend 向前添加数据
func (s *Int64s) Prepend(data ...int64) {
	slice := *s
	ln, cp, num := len(slice), cap(slice), len(data)
	total := ln + num
	if total <= cp {
		if ln == 0 {
			copy(slice[0:total], data)
		} else if num > 0 {
			copy(slice[num:total], slice)
			copy(slice[:num], data)
		}
		*s = slice[:total]
	} else {
		*s = append(data[0:num:num], slice...)
	}
}

// Insert 插入数据
func (s *Int64s) Insert(index int, data ...int64) {
	if index == 0 {
		s.Prepend(data...)
	} else {
		slice := *s
		ln := len(slice)
		if index == ln {
			*s = append(slice, data...)
		} else {
			if index < 0 || index > ln {
				panic("runtime error: slice bounds out of range")
			}
			cp, num := cap(slice), len(data)
			total := ln + num
			if total > cp {
				slice = make([]int64, total)
				copy(slice, *s)
			}
			if num > 0 {
				copy(slice[index+num:], slice[index:])
				copy(slice[index:], data)
			}
			*s = slice
		}
	}
}

// Remove 从指定位置删除指定数量的数据，返回被删除的数据数量
func (s *Int64s) Remove(index, num int) int {
	if num == 0 {
		return 0
	}
	slice := *s
	ln := len(slice)
	if index < 0 || index > ln {
		panic("runtime error: slice bounds out of range")
	}
	end := index + num
	if end == ln {
		*s = slice[:index]
		return num
	} else {
		if end > ln {
			*s = slice[:index]
			return ln - index
		}
		copy(slice[index:], slice[index+num:])
		*s = slice[:ln-num]
		return num
	}
}
