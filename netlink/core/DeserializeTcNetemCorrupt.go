package core

import "unsafe"

// DeserializeTcNetemCorrupt - Converts a byte slice into a TcNetemCorrupt object.
func DeserializeTcNetemCorrupt(b []byte) *TcNetemCorrupt {
	return (*TcNetemCorrupt)(unsafe.Pointer(&b[0:SizeofTcNetemCorrupt][0]))
}
