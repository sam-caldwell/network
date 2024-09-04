package core

import "unsafe"

// Serialize - Converts the TcNetemCorrupt object into a byte slice.
func (msg *TcNetemCorrupt) Serialize() []byte {
	return (*(*[SizeofTcNetemCorrupt]byte)(unsafe.Pointer(msg)))[:]
}
