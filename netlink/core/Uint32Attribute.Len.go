package core

import "unsafe"

// Len returns the length of the Uint32Attribute struct in bytes.
func (a *Uint32Attribute) Len() int {
	return int(unsafe.Sizeof(*a))
}
