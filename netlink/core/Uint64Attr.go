//go:build linux

package core

// Uint64Attr - convert uint64 to []byte using system-native endianness
func Uint64Attr(v uint64) []byte {

	bytes := make([]byte, 8)

	NativeEndian.PutUint64(bytes, v)

	return bytes

}
