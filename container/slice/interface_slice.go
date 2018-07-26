package slice

import "math/rand"

// Interfaces 定义方便操作[]interface{}的类型
type Interfaces []interface{}

// Len 返回slice长度
func (s Interfaces) Len() int {
	return len(s)
}

// Swap 交换两个位置上的数据
func (s Interfaces) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Cap 返回slice容量
func (s Interfaces) Cap() int {
	return cap(s)
}

// Copy 拷贝一份数据
func (s Interfaces) Copy() Interfaces {
	cp := make([]interface{}, len(s))
	copy(cp, s)
	return cp
}

// Shuffle 随机打乱slice数据
func (s Interfaces) Shuffle() Interfaces {
	ln := len(s)
	for i := 0; i < ln; i++ {
		j := rand.Intn(i + 1)
		s[j], s[i] = s[i], s[j]
	}
	return s
}

// Append 追加数据
func (s *Interfaces) Append(data ...interface{}) {
	*s = append(*s, data...)
}

// Prepend 向前添加数据
func (s *Interfaces) Prepend(data ...interface{}) {
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
func (s *Interfaces) Insert(index int, data ...interface{}) {
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
				slice = make([]interface{}, total)
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
func (s *Interfaces) Remove(index, num int) int {
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
