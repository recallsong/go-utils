package slicebuf

// Buffer buffer
type Buffer struct {
	data []byte
	pos  int
}

// New create a clice buffer
func New(cap int) *Buffer {
	return &Buffer{data: make([]byte, cap, cap)}
}

func (b *Buffer) Require(n int) []byte {
	if n <= 0 {
		return nil
	}
	length := b.pos + n
	if len(b.data) < length {
		var cap int
		switch {
		case length <= 1024:
			cap = length * 2
		case length <= 1024*1024:
			cap = int(float64(length) * 1.5)
		case length <= 1024*1024*1024:
			cap = int(float64(length) * 1.2)
		default:
			cap = int(float64(length) * 1.1)
		}
		data := make([]byte, cap, cap)
		copy(data, b.data[0:b.pos])
		b.data = data
	}
	data := b.data[b.pos:length:length]
	b.pos = length
	return data
}

func (b *Buffer) Bytes() []byte {
	return b.data[0:b.pos]
}

func (b *Buffer) Reset() {
	b.pos = 0
}

func (b *Buffer) Len() int {
	return b.pos
}
