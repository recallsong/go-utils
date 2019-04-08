package size

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesFormat(t *testing.T) {
	// B
	b := FormatBytes(0)
	assert.Equal(t, "0", b)
	// B
	b = FormatBytes(515)
	assert.Equal(t, "515B", b)

	// KB
	b = FormatBytes(31323)
	assert.Equal(t, "30.59KB", b)

	// MB
	b = FormatBytes(13231323)
	assert.Equal(t, "12.62MB", b)

	// GB
	b = FormatBytes(7323232398)
	assert.Equal(t, "6.82GB", b)

	// TB
	b = FormatBytes(7323232398434)
	assert.Equal(t, "6.66TB", b)

	// PB
	b = FormatBytes(9923232398434432)
	assert.Equal(t, "8.81PB", b)

	// EB
	b = FormatBytes(math.MaxInt64)
	assert.Equal(t, "8.00EB", b)
}

func TestBytesParseErrors(t *testing.T) {
	_, err := ParseBytes("B999")
	if assert.Error(t, err) {
		assert.EqualError(t, err, "error parsing value=B999")
	}
}

func TestFloats(t *testing.T) {
	// From string:
	str := "12.25KB"
	value, err := ParseBytes(str)
	assert.NoError(t, err)
	assert.Equal(t, int64(12544), value)

	str2 := FormatBytes(value)
	assert.Equal(t, str, str2)

	// To string:
	val := int64(13233029)
	str = FormatBytes(val)
	assert.Equal(t, "12.62MB", str)

	val2, err := ParseBytes(str)
	assert.NoError(t, err)
	assert.Equal(t, val, val2)
}

func TestBytesParse(t *testing.T) {
	// B
	b, err := ParseBytes("999")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(999), b)
	}
	b, err = ParseBytes("-100")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(-100), b)
	}
	b, err = ParseBytes("100.1")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(100), b)
	}
	b, err = ParseBytes("515B")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(515), b)
	}

	// B with space
	b, err = ParseBytes("515 B")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(515), b)
	}

	// KB
	b, err = ParseBytes("12.25KB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12544), b)
	}
	b, err = ParseBytes("12KB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12288), b)
	}
	b, err = ParseBytes("12K")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12288), b)
	}

	// KB with space
	b, err = ParseBytes("12.25 KB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12544), b)
	}
	b, err = ParseBytes("12 KB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12288), b)
	}
	b, err = ParseBytes("12 K")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(12288), b)
	}

	// MB
	b, err = ParseBytes("2MB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(2097152), b)
	}
	b, err = ParseBytes("2M")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(2097152), b)
	}

	// GB with space
	b, err = ParseBytes("6 GB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(6442450944), b)
	}
	b, err = ParseBytes("6 G")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(6442450944), b)
	}

	// GB
	b, err = ParseBytes("6GB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(6442450944), b)
	}
	b, err = ParseBytes("6G")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(6442450944), b)
	}

	// TB
	b, err = ParseBytes("5TB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(5497558138880), b)
	}
	b, err = ParseBytes("5T")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(5497558138880), b)
	}

	// TB with space
	b, err = ParseBytes("5 TB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(5497558138880), b)
	}
	b, err = ParseBytes("5 T")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(5497558138880), b)
	}

	// PB
	b, err = ParseBytes("9PB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(10133099161583616), b)
	}
	b, err = ParseBytes("9P")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(10133099161583616), b)
	}

	// PB with space
	b, err = ParseBytes("9 PB")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(10133099161583616), b)
	}
	b, err = ParseBytes("9 P")
	if assert.NoError(t, err) {
		assert.Equal(t, int64(10133099161583616), b)
	}

	// EB
	b, err = ParseBytes("8EB")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b-1)
	}
	b, err = ParseBytes("8E")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b-1)
	}

	// EB with spaces
	b, err = ParseBytes("8 EB")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b-1)
	}
	b, err = ParseBytes("8 E")
	if assert.NoError(t, err) {
		assert.True(t, math.MaxInt64 == b-1)
	}
}

func TestBytesJSON(t *testing.T) {
	var n Bytes = 2048
	var s = struct {
		Bytes Bytes `json:"bytes"`
	}{
		Bytes: n,
	}
	data, err := json.Marshal(s)
	if assert.NoError(t, err) {
		assert.Equal(t, `{"bytes":"2.00KB"}`, string(data))
	}
	s.Bytes = 0
	err = json.Unmarshal(data, &s)
	if assert.NoError(t, err) {
		assert.Equal(t, Bytes(2048), s.Bytes)
	}

	data = []byte(`{"bytes": "4KB"}`)
	s.Bytes = 0
	err = json.Unmarshal(data, &s)
	if assert.NoError(t, err) {
		assert.Equal(t, Bytes(4096), s.Bytes)
	}

	data = []byte(`{"bytes": 4096}`)
	s.Bytes = 0
	err = json.Unmarshal(data, &s)
	if assert.NoError(t, err) {
		assert.Equal(t, Bytes(4096), s.Bytes)
	}

	var sp = struct {
		Bytes *Bytes `json:"bytes"`
	}{
		Bytes: &n,
	}
	data = []byte(`{"bytes": null}`)
	err = json.Unmarshal(data, &sp)
	if assert.NoError(t, err) {
		assert.Equal(t, (*Bytes)(nil), sp.Bytes)
	}
	data = []byte(`{"bytes": "8KB"}`)
	err = json.Unmarshal(data, &sp)
	if assert.NoError(t, err) {
		assert.Equal(t, Bytes(8*1024), *sp.Bytes)
	}
	data = []byte(`{"bytes": "1024"}`)
	err = json.Unmarshal(data, &sp)
	if assert.NoError(t, err) {
		assert.Equal(t, Bytes(1024), *sp.Bytes)
	}
	data = []byte(`{"bytes": "4096B"}`)
	err = json.Unmarshal(data, &sp)
	if assert.NoError(t, err) {
		assert.Equal(t, Bytes(4096), *sp.Bytes)
	}
	*sp.Bytes = 8 * 1024
	data, err = json.Marshal(sp)
	if assert.NoError(t, err) {
		assert.Equal(t, `{"bytes":"8.00KB"}`, string(data))
	}
}

func TestBytesText(t *testing.T) {
	var n Bytes = 2048
	data, err := n.MarshalText()
	if assert.NoError(t, err) {
		assert.Equal(t, `2.00KB`, string(data))
	}
	err = n.UnmarshalText([]byte(`4.00KB`))
	if assert.NoError(t, err) {
		assert.Equal(t, Bytes(4096), n)
	}
}
