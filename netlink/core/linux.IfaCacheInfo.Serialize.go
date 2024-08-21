package core

import (
	"golang.org/x/sys/unix"
	"unsafe"
)

// Serialize - serialize IfaCacheInfo to byte slice
func (msg *IfaCacheInfo) Serialize() []byte {
	return (*(*[unix.SizeofIfaCacheinfo]byte)(unsafe.Pointer(msg)))[:]
}
