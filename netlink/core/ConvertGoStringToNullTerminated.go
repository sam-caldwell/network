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
	bytes := make([]byte, len(s)+1)
	for i := 0; i < len(s); i++ {
		bytes[i] = s[i]
	}
	bytes[len(s)] = 0
	return bytes
}
