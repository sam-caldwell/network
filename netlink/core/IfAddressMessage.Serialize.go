package core

import (
	"golang.org/x/sys/unix"
	"unsafe"
)

// Serialize - serialize an interface address message as a byte slice.
func (msg *IfAddressMessage) Serialize() []byte {
	return (*(*[unix.SizeofIfAddrmsg]byte)(unsafe.Pointer(msg)))[:]
}
