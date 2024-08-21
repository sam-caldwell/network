package core

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

// DeserializeIfAddressMessage - deserialize the interface address message
func DeserializeIfAddressMessage(b []byte) *IfAddressMessage {

	return (*IfAddressMessage)(unsafe.Pointer(&b[0:unix.SizeofIfAddrmsg][0]))

}
