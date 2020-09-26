package conv

import (
	"fmt"
	"strconv"
	"time"
)

// ToInt 将interface{}转换为int，失败则返回defVal
func ToInt(v interface{}, defVal int) int {
	return int(ToInt64(v, int64(defVal)))
}

// ToInt8 将interface{}转换为int8，失败则返回defVal
func ToInt8(v interface{}, defVal int8) int8 {
	return int8(ToInt64(v, int64(defVal)))
}

// ToInt16 将interface{}转换为int16，失败则返回defVal
func ToInt16(v interface{}, defVal int16) int16 {
	return int16(ToInt64(v, int64(defVal)))
}

// ToInt32 将interface{}转换为int32，失败则返回defVal
func ToInt32(v interface{}, defVal int32) int32 {
	return int32(ToInt64(v, int64(defVal)))
}

// ToInt64 将interface{}转换为int64，失败则返回defVal
func ToInt64(v interface{}, defVal int64) int64 {
	switch val := v.(type) {
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
	case time.Duration:
		return int64(val)
	}
	return defVal
}

// ToUint 将interface{}转换为uint，失败则返回defVal
func ToUint(v interface{}, defVal uint) uint {
	return uint(ToUint64(v, uint64(defVal)))
}

// ToUint8 将interface{}转换为uint8，失败则返回defVal
func ToUint8(v interface{}, defVal uint8) uint8 {
	return uint8(ToUint64(v, uint64(defVal)))
}

// ToUint16 将interface{}转换为uint16，失败则返回defVal
func ToUint16(v interface{}, defVal uint16) uint16 {
	return uint16(ToUint64(v, uint64(defVal)))
}

// ToUint32 将interface{}转换为uint32，失败则返回defVal
func ToUint32(v interface{}, defVal uint32) uint32 {
	return uint32(ToUint64(v, uint64(defVal)))
}

// ToUint64 将interface{}转换为uint64，失败则返回defVal
func ToUint64(v interface{}, defVal uint64) uint64 {
	switch val := v.(type) {
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
		return val
	case float32:
		return uint64(val)
	case float64:
		return uint64(val)
	case time.Duration:
		return uint64(val)
	}
	return defVal
}

// ToFloat32 将interface{}转换为float32，失败则返回defVal
func ToFloat32(v interface{}, defVal float32) float32 {
	return float32(ToFloat64(v, float64(defVal)))
}

// ToFloat64 将interface{}转换为float64，失败则返回defVal
func ToFloat64(v interface{}, defVal float64) float64 {
	switch val := v.(type) {
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
		return val
	case time.Duration:
		return float64(val)
	}
	return defVal
}

// ToBool 将interface{}转换为bool，失败则返回defVal
func ToBool(v interface{}, defVal bool) bool {
	switch val := v.(type) {
	case bool:
		return val
	}
	return defVal
}

// ToDuration 将interface{}转换为time.Duration，失败则返回defVal
func ToDuration(v interface{}, defVal time.Duration) time.Duration {
	switch val := v.(type) {
	case time.Duration:
		return val
	case string:
		d, err := time.ParseDuration(val)
		if err != nil {
			return defVal
		}
		return d
	case int:
		return time.Duration(val)
	case int8:
		return time.Duration(val)
	case int16:
		return time.Duration(val)
	case int32:
		return time.Duration(val)
	case int64:
		return time.Duration(val)
	case uint:
		return time.Duration(val)
	case uint8:
		return time.Duration(val)
	case uint16:
		return time.Duration(val)
	case uint32:
		return time.Duration(val)
	case uint64:
		return time.Duration(val)
	case float32:
		return time.Duration(val)
	case float64:
		return time.Duration(val)
	}
	return defVal
}

// Int 将interface{}转换为int
func Int(v interface{}) (int, error) {
	val, err := Int64(v)
	if err != nil {
		return 0, nil
	}
	return int(val), nil
}

// Int8 将interface{}转换为int8
func Int8(v interface{}) (int8, error) {
	val, err := Int64(v)
	if err != nil {
		return 0, nil
	}
	return int8(val), nil
}

// Int16 将interface{}转换为int16
func Int16(v interface{}) (int16, error) {
	val, err := Int64(v)
	if err != nil {
		return 0, nil
	}
	return int16(val), nil
}

// Int32 将interface{}转换为int32
func Int32(v interface{}) (int32, error) {
	val, err := Int64(v)
	if err != nil {
		return 0, nil
	}
	return int32(val), nil
}

// Int64 将interface{}转换为int64
func Int64(v interface{}) (int64, error) {
	switch val := v.(type) {
	case int:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case int64:
		return int64(val), nil
	case uint:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint64:
		return int64(val), nil
	case float32:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case time.Duration:
		return int64(val), nil
	}
	return 0, fmt.Errorf("fail to convert %v to int", v)
}

// Uint 将interface{}转换为uint
func Uint(v interface{}) (uint, error) {
	val, err := Uint64(v)
	if err != nil {
		return 0, nil
	}
	return uint(val), nil
}

// Uint8 将interface{}转换为uint8
func Uint8(v interface{}) (uint8, error) {
	val, err := Uint64(v)
	if err != nil {
		return 0, nil
	}
	return uint8(val), nil
}

// Uint16 将interface{}转换为uint16
func Uint16(v interface{}) (uint16, error) {
	val, err := Uint64(v)
	if err != nil {
		return 0, nil
	}
	return uint16(val), nil
}

// Uint32 将interface{}转换为uint32
func Uint32(v interface{}) (uint32, error) {
	val, err := Uint64(v)
	if err != nil {
		return 0, nil
	}
	return uint32(val), nil
}

// Uint64 将interface{}转换为uint64
func Uint64(v interface{}) (uint64, error) {
	switch val := v.(type) {
	case int:
		return uint64(val), nil
	case int8:
		return uint64(val), nil
	case int16:
		return uint64(val), nil
	case int32:
		return uint64(val), nil
	case int64:
		return uint64(val), nil
	case uint:
		return uint64(val), nil
	case uint8:
		return uint64(val), nil
	case uint16:
		return uint64(val), nil
	case uint32:
		return uint64(val), nil
	case uint64:
		return val, nil
	case float32:
		return uint64(val), nil
	case float64:
		return uint64(val), nil
	case time.Duration:
		return uint64(val), nil
	}
	return 0, fmt.Errorf("fail to convert %v to uint", v)
}

// Float32 将interface{}转换为float32
func Float32(v interface{}) (float32, error) {
	val, err := Float64(v)
	if err != nil {
		return 0, err
	}
	return float32(val), nil
}

// Float64 将interface{}转换为float64
func Float64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case int:
		return float64(val), nil
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case uint:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case float32:
		return float64(val), nil
	case float64:
		return val, nil
	case time.Duration:
		return float64(val), nil
	}
	return 0, fmt.Errorf("fail to convert %v to float", v)
}

// Bool 将interface{}转换为bool
func Bool(v interface{}) (bool, error) {
	switch val := v.(type) {
	case bool:
		return val, nil
	}
	return false, fmt.Errorf("fail to convert %v to bool", v)
}

// String 将interface{}转换为string
func String(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case error:
		return val.Error()
	case int:
		return strconv.Itoa(val)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float64:
		return FormatFloat64(val, -1)
	case float32:
		return FormatFloat32(val, -1)
	default:
		return fmt.Sprint(val)
	}
}

// ToString deprcated ! Use String
func ToString(v interface{}) string {
	return String(v)
}

// Duration 将interface{}转换为time.Duration
func Duration(v interface{}) (time.Duration, error) {
	switch val := v.(type) {
	case time.Duration:
		return val, nil
	case string:
		d, err := time.ParseDuration(val)
		if err != nil {
			return d, err
		}
		return d, nil
	case int:
		return time.Duration(val), nil
	case int8:
		return time.Duration(val), nil
	case int16:
		return time.Duration(val), nil
	case int32:
		return time.Duration(val), nil
	case int64:
		return time.Duration(val), nil
	case uint:
		return time.Duration(val), nil
	case uint8:
		return time.Duration(val), nil
	case uint16:
		return time.Duration(val), nil
	case uint32:
		return time.Duration(val), nil
	case uint64:
		return time.Duration(val), nil
	case float32:
		return time.Duration(val), nil
	case float64:
		return time.Duration(val), nil
	}
	return 0, fmt.Errorf("fail to convert %v to duration", v)
}
