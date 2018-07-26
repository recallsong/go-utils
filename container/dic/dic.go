/*
	基于map[string]interface{}定义的Dic类型，定义一些扩展的方法
*/
package dic

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/recallsong/go-utils/conv"
)

var (
	ErrTypeCast = errors.New("failed to type cast")
)

// Dic 基于map[string]interface{}定义，更方便使用
type Dic map[string]interface{}

// New 创建Dic
func New() Dic {
	return make(Dic)
}

// FromMap 将map[string]interface{}转换成Dic类型
func FromMap(mp map[string]interface{}) Dic {
	return Dic(mp)
}

// FromJSON 将JSON字符串解析到Dic
func FromJSON(jsonStr string) (Dic, error) {
	d := Dic{}
	err := json.Unmarshal([]byte(jsonStr), &d)
	if err != nil {
		return nil, err
	}
	return d, err
}

// LoadJSON 将json字符串解析到Dic，失败将返回错误
func (d Dic) LoadJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), &d)
}

// ToMap 将Dic转换成 map[string]interface{}
func (d Dic) ToMap() map[string]interface{} {
	return map[string]interface{}(d)
}

// JSON 将Dic转换成JSON字符串，失败则返回空字符串
func (d Dic) JSON() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// JSON 将Dic转换成JSON字符串，失败则返回空字符串
func (d Dic) String() string {
	bytes, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprint(map[string]interface{}(d))
	}
	return string(bytes)
}

// Contains 判断Dic是否包含该key
func (d Dic) Contains(key string) bool {
	_, ok := d[key]
	return ok
}

// Equals 对两个Dic进行深度比较，返回是否相等
func (d Dic) Equals(obj interface{}) bool {
	return reflect.DeepEqual(d, obj)
}

// Copy 对Dic进行浅拷贝
func (d Dic) Copy() Dic {
	cp := Dic{}
	for k, v := range d {
		cp[k] = v
	}
	return cp
}

// Delete 删除key，并返回将要被删除的值
func (d Dic) Delete(key string) interface{} {
	v, ok := d[key]
	if ok {
		delete(d, key)
		return v
	}
	return nil
}

// Clear 删除所有的健值对
func (d Dic) Clear() {
	for k := range d {
		delete(d, k)
	}
}

// Keys 获取所有的key
func (d Dic) Keys() []string {
	keys := make([]string, len(d))
	i := 0
	for k := range d {
		keys[i] = k
		i = i + 1
	}
	return keys
}

// Values 获取所有的值
func (d Dic) Values() []interface{} {
	vals := make([]interface{}, len(d))
	i := 0
	for _, v := range d {
		vals[i] = v
		i = i + 1
	}
	return vals
}

// Size 返回Dic的元素数量
func (d Dic) Size() int {
	return len(d)
}

// Set 设置健值对
func (d Dic) Set(key string, value interface{}) Dic {
	d[key] = value
	return d
}

// Put 设置健值对
func (d Dic) Put(key string, value interface{}) {
	d[key] = value
}

// Get 根据key获取value
func (d Dic) Get(key string) interface{} {
	return d[key]
}

// GetStrng 根据key获取string类型的值，如果值不是string类型，将返回defVal
func (d Dic) GetString(key string, defVal string) string {
	switch val := d[key].(type) {
	case string:
		return val
	}
	return defVal
}

// GetBool 根据key获取bool类型的值，如果值不是int类型，将返回defVal
func (d Dic) GetBool(key string, defVal bool) bool {
	return conv.ToBool(d[key], defVal)
}

// GetInt 根据key获取int类型的值，如果值不是int类型，将返回defVal
func (d Dic) GetInt(key string, defVal int) int {
	return conv.ToInt(d[key], defVal)
}

// GetInt8 根据key获取int8类型的值，如果值不是int8类型，将返回defVal
func (d Dic) GetInt8(key string, defVal int8) int8 {
	return conv.ToInt8(d[key], defVal)
}

// GetInt16 根据key获取int16类型的值，如果值不是int16类型，将返回defVal
func (d Dic) GetInt16(key string, defVal int16) int16 {
	return conv.ToInt16(d[key], defVal)
}

// GetInt32 根据key获取int32类型的值，如果值不是int32类型，将返回defVal
func (d Dic) GetInt32(key string, defVal int32) int32 {
	return conv.ToInt32(d[key], defVal)
}

// GetInt64 根据key获取int64类型的值，如果值不是int64类型，将返回defVal
func (d Dic) GetInt64(key string, defVal int64) int64 {
	return conv.ToInt64(d[key], defVal)
}

// GetUint 根据key获取uint类型的值，如果值不是uint类型，将返回defVal
func (d Dic) GetUint(key string, defVal uint) uint {
	return conv.ToUint(d[key], defVal)
}

// GetUint8 根据key获取uint8类型的值，如果值不是uint8类型，将返回defVal
func (d Dic) GetUint8(key string, defVal uint8) uint8 {
	return conv.ToUint8(d[key], defVal)
}

// GetUint16 根据key获取uint16类型的值，如果值不是uint16类型，将返回defVal
func (d Dic) GetUint16(key string, defVal uint16) uint16 {
	return conv.ToUint16(d[key], defVal)
}

// GetUint32 根据key获取uint32类型的值，如果值不是uint32类型，将返回defVal
func (d Dic) GetUint32(key string, defVal uint32) uint32 {
	return conv.ToUint32(d[key], defVal)
}

// GetUint64 根据key获取uint64类型的值，如果值不是uint64类型，将返回defVal
func (d Dic) GetUint64(key string, defVal uint64) uint64 {
	return conv.ToUint64(d[key], defVal)
}

// GetFloat32 根据key获取float32类型的值，如果值不是float32类型，将返回defVal
func (d Dic) GetFloat32(key string, defVal float32) float32 {
	return conv.ToFloat32(d[key], defVal)
}

// GetFloat64 根据key获取float64类型的值，如果值不是float64类型，将返回defVal
func (d Dic) GetFloat64(key string, defVal float64) float64 {
	return conv.ToFloat64(d[key], defVal)
}

// GetDic 根据key获取Dic类型的值，如果失败则返回错误
func (d Dic) GetDic(key string) (Dic, error) {
	v, ok := d[key]
	if ok {
		switch val := v.(type) {
		case map[string]interface{}:
			return FromMap(val), nil
		case Dic:
			return val, nil
		}
		return nil, ErrTypeCast
	}
	return nil, nil
}

// GetDic 根据key获取map[string]interface{}类型的值，如果失败则返回错误
func (d Dic) GetMap(key string) (map[string]interface{}, error) {
	val, err := d.GetDic(key)
	if err != nil {
		return nil, err
	}
	return val.ToMap(), err
}
