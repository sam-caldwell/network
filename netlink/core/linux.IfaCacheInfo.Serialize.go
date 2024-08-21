package core

import (
	"unsafe"
)

// Serialize - serialize IfaCacheInfo to byte slice
func (msg *IfaCacheInfo) Serialize() []byte {
	return (*(*[SizeOfIfaCacheinfo]byte)(unsafe.Pointer(msg)))[:]
}
