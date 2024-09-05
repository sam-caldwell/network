package core

import "unsafe"

func (msg *TcSkbEdit) Serialize() []byte {
	return (*(*[SizeofTcSkbEdit]byte)(unsafe.Pointer(x)))[:]
}
