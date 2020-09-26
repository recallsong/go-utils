package conv

import (
	"strconv"
)

// FormatFloat64 .
func FormatFloat64(value float64, precision int) string {
	text := strconv.FormatFloat(value, 'f', precision, 64)
	return CleanFloat(text)
}

// FormatFloat32 .
func FormatFloat32(value float32, precision int) string {
	text := strconv.FormatFloat(float64(value), 'f', precision, 32)
	return CleanFloat(text)
}

// CleanFloat .
func CleanFloat(value string) string {
	idx := len(value)
	i := idx - 1
	for ; i >= 0; i-- {
		if value[i] == '.' {
			break
		}
		if value[i] == '0' {
			idx--
		} else {
			break
		}
	}

	if i >= 0 {
		if value[idx-1] == '.' {
			idx--
		}
		value = value[0:idx]
	}
	return value
}
