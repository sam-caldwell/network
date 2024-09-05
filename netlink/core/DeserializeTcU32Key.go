package core

import "unsafe"

func DeserializeTcU32Key(b []byte) *TcU32Key {
	return (*TcU32Key)(unsafe.Pointer(&b[0:SizeofTcU32Key][0]))
}
