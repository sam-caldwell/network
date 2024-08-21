//go:build linux

package core

import "unsafe"

func DeserializeBridgeVlanInfo(b []byte) *BridgeVlanInfo {
	return (*BridgeVlanInfo)(unsafe.Pointer(&b[0:SizeofBridgeVlanInfo][0]))
}
