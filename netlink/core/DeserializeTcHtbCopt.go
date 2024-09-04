package core

import "unsafe"

func DeserializeTcHtbCopt(b []byte) *TcHtbCopt {
	return (*TcHtbCopt)(unsafe.Pointer(&b[0:SizeofTcHtbCopt][0]))
}
