package strings

import (
	"math/rand"
	"sort"
	"strings"
)

// Strings 定义方便操作[]string的类型
type Strings []string

// Len 返回slice长度
func (s Strings) Len() int {
	return len(s)
}

// Less 比较两个位置上的数据
func (s Strings) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap 交换两个位置上的数据
func (s Strings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Cap 返回slice容量
func (s Strings) Cap() int {
	return cap(s)
}

// Copy 拷贝一份数据
func (s Strings) Copy() Strings {
	cp := make([]string, len(s))
	copy(cp, s)
	return cp
}

// Sort 对切片进行排序
func (s Strings) Sort() Strings {
	sort.Sort(s)
	return s
}

func (s Strings) ReverseSort() Strings {
	sort.Sort(sort.Reverse(s))
	return s
}

// Shuffle 随机打乱slice数据
func (s Strings) Shuffle() Strings {
	ln := len(s)
	for i := 0; i < ln; i++ {
		j := rand.Intn(i + 1)
		s[j], s[i] = s[i], s[j]
	}
	return s
}

// Append 追加数据
func (s *Strings) Append(data ...string) {
	*s = append(*s, data...)
}

// Prepend 向前添加数据
func (s *Strings) Prepend(data ...string) {
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
func (s *Strings) Insert(index int, data ...string) {
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
				slice = make([]string, total)
				copy(slice, *s)
			} else {
				slice = slice[0:total:cp]
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
func (s *Strings) Remove(index, num int) int {
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

// Join 将多个字符串连接
func (s Strings) Join(sep string) string {
	return strings.Join(s, sep)
}
