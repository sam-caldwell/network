//go:build linux

package core

import (
	"unsafe"
)

func (b *BridgeVlanInfo) Serialize() []byte {
	return (*(*[SizeofBridgeVlanInfo]byte)(unsafe.Pointer(b)))[:]
}
