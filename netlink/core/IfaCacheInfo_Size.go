package core

import "unsafe"

const (
	// IfaCacheInfoSize - 0x16 // bytes as derived from unix.IfaCacheInfoSize
	IfaCacheInfoSize = int(unsafe.Sizeof(IfaCacheInfo{}))
)
