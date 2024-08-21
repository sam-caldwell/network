package core

import (
	"unsafe"
)

// Serialize - Serialize the IfInfoMsg structure
// ToDo: let's do this in a safer way. This has the potential for bad things.
func (msg *IfInfoMsg) Serialize() []byte {
	return (*(*[SizeofIfInfoMsg]byte)(unsafe.Pointer(msg)))[:]
}
