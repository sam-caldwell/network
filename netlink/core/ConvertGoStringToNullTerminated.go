//go:build linux

package core

// ConvertGoStringToNullTerminated - converts a Go string into a null-terminated byte slice.
//
// This function takes a string `s` and returns a byte slice that contains the
// characters of the string followed by a null byte (`0`). Null-terminated strings
// are commonly used in C and other low-level programming languages, especially
// in system calls or when interfacing with C libraries.
//
// Example:
//
//	s := "test"
//	result := ConvertGoStringToNullTerminated(s)
//	// result will be []byte{'t', 'e', 's', 't', 0}
func ConvertGoStringToNullTerminated(s string) []byte {

	const Null byte = 0x00

	out := make([]byte, len(s)+1)

	raw := []byte(s)

	copy(out, raw)

	out[len(raw)] = Null

	return out
}
