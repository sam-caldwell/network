package core

import (
	"unsafe"
)

func DeserializeXfrmSelector(b []byte) *XfrmSelector {
	return (*XfrmSelector)(unsafe.Pointer(&b[0:SizeofXfrmSelector][0]))
}
