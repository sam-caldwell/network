//go:build linux

package core

// NonNullTerminated - convert Go string into a byte slice without a null terminator.
//
// This function takes a string `s` and returns a byte slice containing the characters
// of the string without adding a null byte (`0`) at the end. This is useful in scenarios
// where you need to pass raw bytes that represent a string, but a null terminator is
// either unnecessary or unwanted.
//
// Example:
//
//	s := "test"
//	result := NonZeroTerminated(s)
//	// result will be []byte{'t', 'e', 's', 't'}
func NonNullTerminated(s string) []byte {

	return []byte(s)

}
