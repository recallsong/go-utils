package encoding

// RawBytes .
type RawBytes []byte

// Bytes convert bytes
func (bs RawBytes) Bytes() []byte {
	return bs
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (bs RawBytes) MarshalBinary() ([]byte, error) {
	return bs, nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (bs *RawBytes) UnmarshalBinary(data []byte) error {
	*bs = data
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (bs RawBytes) MarshalJSON() ([]byte, error) {
	return bs, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (bs *RawBytes) UnmarshalJSON(data []byte) error {
	*bs = data
	return nil
}

// GobEncode implements the gob.GobEncoder interface.
func (bs RawBytes) GobEncode() ([]byte, error) {
	return bs, nil
}

// GobDecode implements the gob.GobDecoder interface.
func (bs *RawBytes) GobDecode(data []byte) error {
	*bs = data
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
// The time is formatted in RFC 3339 format, with sub-second precision added if present.
func (bs RawBytes) MarshalText() ([]byte, error) {
	return bs, nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The time is expected to be in RFC 3339 format.
func (bs *RawBytes) UnmarshalText(data []byte) error {
	*bs = data
	return nil
}
