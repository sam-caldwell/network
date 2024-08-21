package core

import (
	"golang.org/x/sys/unix"
	"unsafe"
)

func (msg *IfAddressMessage) Serialize() []byte {
	return (*(*[unix.SizeofIfAddrmsg]byte)(unsafe.Pointer(msg)))[:]
}
