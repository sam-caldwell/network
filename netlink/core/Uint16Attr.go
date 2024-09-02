//go:build linux

package core

// Uint16Attr - convert uint16 to []byte using system-native endianness
func Uint16Attr(v uint16) []byte {

	bytes := make([]byte, 2)
	nativeEndian.PutUint16(bytes, v)
	return bytes
}
