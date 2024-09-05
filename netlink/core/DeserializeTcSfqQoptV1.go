package core

import "unsafe"

func DeserializeTcSfqQoptV1(b []byte) *TcSfqQoptV1 {
	return (*TcSfqQoptV1)(unsafe.Pointer(&b[0:SizeofTcSfqQoptV1][0]))
}
