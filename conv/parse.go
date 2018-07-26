package conv

import "strconv"

// ParseInt 解析字符串为int类型的值，失败则返回defVal
func ParseInt(str string, defVal int) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defVal
	}
	return val
}

// ParseInt8 解析字符串为in8类型的值，失败则返回defVal
func ParseInt8(str string, defVal int8) int8 {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defVal
	}
	return int8(val)
}

// ParseInt16 解析字符串为in16类型的值，失败则返回defVal
func ParseInt16(str string, defVal int16) int16 {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defVal
	}
	return int16(val)
}

// ParseInt32 解析字符串为in32类型的值，失败则返回defVal
func ParseInt32(str string, defVal int32) int32 {
	val, err := strconv.Atoi(str)
	if err != nil {
		return defVal
	}
	return int32(val)
}

// ParseInt64 解析字符串为int64类型的值，失败则返回defVal
func ParseInt64(str string, defVal int64) int64 {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defVal
	}
	return val
}

// ParseUint 解析字符串为uint类型的值，失败则返回defVal
func ParseUint(str string, defVal uint) uint {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return defVal
	}
	return uint(val)
}

// ParseUint8 解析字符串为uint8类型的值，失败则返回defVal
func ParseUint8(str string, defVal uint8) uint8 {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return defVal
	}
	return uint8(val)
}

// ParseUint16 解析字符串为uint16类型的值，失败则返回defVal
func ParseUint16(str string, defVal uint16) uint16 {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return defVal
	}
	return uint16(val)
}

// ParseUint32 解析字符串为uint32类型的值，失败则返回defVal
func ParseUint32(str string, defVal uint32) uint32 {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return defVal
	}
	return uint32(val)
}

// ParseUint64 解析字符串为uint64类型的值，失败则返回defVal
func ParseUint64(str string, defVal uint64) uint64 {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return defVal
	}
	return val
}

// ParseFloat32 解析字符串为float32类型的值，失败则返回defVal
func ParseFloat32(str string, defVal float32) float32 {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return defVal
	}
	return float32(val)
}

// ParseFloat64 解析字符串为float64类型的值，失败则返回defVal
func ParseFloat64(str string, defVal float64) float64 {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return defVal
	}
	return float64(val)
}

// ParseBool 解析字符串为bool类型的值，失败则返回defVal
func ParseBool(str string, defVal bool) bool {
	val, err := strconv.ParseBool(str)
	if err != nil {
		return defVal
	}
	return val
}
