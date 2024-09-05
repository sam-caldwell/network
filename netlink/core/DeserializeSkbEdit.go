package core

import "unsafe"

func DeserializeSkbEdit(b []byte) *TcSkbEdit {
	return (*TcSkbEdit)(unsafe.Pointer(&b[0:SizeofTcSkbEdit][0]))
}
