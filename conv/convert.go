package conv

import (
	"fmt"
	"strconv"
)

// ToInt 将interface{}转换为int，失败则返回defVal
func ToInt(obj interface{}, defVal int) int {
	return int(ToInt64(obj, int64(defVal)))
}

// ToInt8 将interface{}转换为int8，失败则返回defVal
func ToInt8(obj interface{}, defVal int8) int8 {
	return int8(ToInt64(obj, int64(defVal)))
}

// ToInt16 将interface{}转换为int16，失败则返回defVal
func ToInt16(obj interface{}, defVal int16) int16 {
	return int16(ToInt64(obj, int64(defVal)))
}

// ToInt32 将interface{}转换为int32，失败则返回defVal
func ToInt32(obj interface{}, defVal int32) int32 {
	return int32(ToInt64(obj, int64(defVal)))
}

// ToInt64 将interface{}转换为int64，失败则返回defVal
func ToInt64(obj interface{}, defVal int64) int64 {
	switch val := obj.(type) {
	case int:
		return int64(val)
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return int64(val)
	case uint:
		return int64(val)
	case uint8:
		return int64(val)
	case uint16:
		return int64(val)
	case uint32:
		return int64(val)
	case uint64:
		return int64(val)
	case float32:
		return int64(val)
	case float64:
		return int64(val)
	}
	return defVal
}

// ToUint 将interface{}转换为uint，失败则返回defVal
func ToUint(obj interface{}, defVal uint) uint {
	return uint(ToUint64(obj, uint64(defVal)))
}

// ToUint8 将interface{}转换为uint8，失败则返回defVal
func ToUint8(obj interface{}, defVal uint8) uint8 {
	return uint8(ToUint64(obj, uint64(defVal)))
}

// ToUint16 将interface{}转换为uint16，失败则返回defVal
func ToUint16(obj interface{}, defVal uint16) uint16 {
	return uint16(ToUint64(obj, uint64(defVal)))
}

// ToUint32 将interface{}转换为uint32，失败则返回defVal
func ToUint32(obj interface{}, defVal uint32) uint32 {
	return uint32(ToUint64(obj, uint64(defVal)))
}

// ToUint64 将interface{}转换为uint64，失败则返回defVal
func ToUint64(obj interface{}, defVal uint64) uint64 {
	switch val := obj.(type) {
	case int:
		return uint64(val)
	case int8:
		return uint64(val)
	case int16:
		return uint64(val)
	case int32:
		return uint64(val)
	case int64:
		return uint64(val)
	case uint:
		return uint64(val)
	case uint8:
		return uint64(val)
	case uint16:
		return uint64(val)
	case uint32:
		return uint64(val)
	case uint64:
		return uint64(val)
	case float32:
		return uint64(val)
	case float64:
		return uint64(val)
	}
	return defVal
}

// ToFloat32 将interface{}转换为float32，失败则返回defVal
func ToFloat32(obj interface{}, defVal float32) float32 {
	return float32(ToFloat64(obj, float64(defVal)))
}

// ToFloat64 将interface{}转换为float64，失败则返回defVal
func ToFloat64(obj interface{}, defVal float64) float64 {
	switch val := obj.(type) {
	case int:
		return float64(val)
	case int8:
		return float64(val)
	case int16:
		return float64(val)
	case int32:
		return float64(val)
	case int64:
		return float64(val)
	case uint:
		return float64(val)
	case uint8:
		return float64(val)
	case uint16:
		return float64(val)
	case uint32:
		return float64(val)
	case uint64:
		return float64(val)
	case float32:
		return float64(val)
	case float64:
		return float64(val)
	}
	return defVal
}

// ToBool 将interface{}转换为bool，失败则返回defVal
func ToBool(obj interface{}, defVal bool) bool {
	switch val := obj.(type) {
	case bool:
		return val
	}
	return defVal
}

// ToByte 将interface{}转换为string
func ToString(obj interface{}) string {
	switch val := obj.(type) {
	case string:
		return val
	case error:
		return val.Error()
	case int:
		return strconv.Itoa(val)
	default:
		return fmt.Sprint(val)
	}
}
