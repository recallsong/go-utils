package size

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/recallsong/go-utils/reflectx"
)

// Bytes .
type Bytes int64

const (
	// B ... Byte Unit
	B = int64(1) << (10 * iota)
	KB
	MB
	GB
	TB
	PB
	EB
)

func (b Bytes) String() string {
	return FormatBytes(int64(b))
}

// Parse parse val to bytes
func (b *Bytes) Parse(val string) error {
	n, err := ParseBytes(val)
	if err != nil {
		return err
	}
	*((*int64)(b)) = n
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (b Bytes) MarshalJSON() ([]byte, error) {
	return []byte("\"" + FormatBytes(int64(b)) + "\""), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (b *Bytes) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) <= 0 {
		return fmt.Errorf("unmarshal empty bytes")
	}
	str := reflectx.BytesToString(data)
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		val, err := ParseBytes(str[1 : len(data)-1])
		if err != nil {
			return fmt.Errorf("fail to unmarshal bytes : %v", err)
		}
		*(*int64)(b) = val
	} else if str == "null" {
		b = nil
	} else {
		val, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return fmt.Errorf("fail to unmarshal bytes : %v", err)
		}
		*(*int64)(b) = val
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
func (b Bytes) MarshalText() ([]byte, error) {
	return []byte(FormatBytes(int64(b))), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (b *Bytes) UnmarshalText(data []byte) error {
	data = bytes.TrimSpace(data)
	val, err := ParseBytes(reflectx.BytesToString(data))
	if err != nil {
		return fmt.Errorf("fail to unmarshal bytes : %v", err)
	}
	*(*int64)(b) = val
	return nil
}

// ParseBytes parse string to bytes number
func ParseBytes(val string) (int64, error) {
	idx := len(val) - 1
	if idx < 0 {
		return 0, fmt.Errorf("error parsing empty value")
	}
	for ; idx > 0; idx-- {
		if ('0' <= val[idx] && val[idx] <= '9') || val[idx] == '.' || val[idx] == ' ' {
			break
		}
	}
	unit := val[idx+1:]
	value := strings.TrimSpace(val[:idx+1])
	bytes, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing value=%s", val)
	}
	unit = strings.ToUpper(unit)
	switch unit {
	case "", "B":
		return int64(bytes), nil
	case "K", "KB":
		return int64(bytes * float64(KB)), nil
	case "M", "MB":
		return int64(bytes * float64(MB)), nil
	case "G", "GB":
		return int64(bytes * float64(GB)), nil
	case "T", "TB":
		return int64(bytes * float64(TB)), nil
	case "P", "PB":
		return int64(bytes * float64(PB)), nil
	case "E", "EB":
		return int64(bytes * float64(EB)), nil
	default:
		return 0, fmt.Errorf("not support size unit %s in %s", unit, val)
	}
}

// FormatBytes convert b to string
func FormatBytes(b int64) string {
	var unit string
	value := float64(b)
	switch {
	case b >= EB:
		value /= float64(EB)
		unit = "EB"
	case b >= PB:
		value /= float64(PB)
		unit = "PB"
	case b >= TB:
		value /= float64(TB)
		unit = "TB"
	case b >= GB:
		value /= float64(GB)
		unit = "GB"
	case b >= MB:
		value /= float64(MB)
		unit = "MB"
	case b >= KB:
		value /= float64(KB)
		unit = "KB"
	case b == 0:
		return "0"
	default:
		return strconv.FormatInt(int64(b), 10) + "B"
	}
	text := strconv.FormatFloat(value, 'f', 2, 64)
	idx := len(text)
	i := idx - 1
	for ; i >= 0; i-- {
		if text[i] == '.' {
			break
		}
		if text[i] == '0' {
			idx--
		} else {
			break
		}
	}
	if i >= 0 {
		if text[idx-1] == '.' {
			idx--
		}
		text = text[0:idx]
	}
	// text := fmt.Sprintf("%f%s", value, unit)
	return text + unit
}
