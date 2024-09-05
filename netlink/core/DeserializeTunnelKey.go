package core

import "unsafe"

func DeserializeTunnelKey(b []byte) *TcTunnelKey {
	return (*TcTunnelKey)(unsafe.Pointer(&b[0:SizeofTcTunnelKey][0]))
}
