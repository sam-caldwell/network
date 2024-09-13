package core

import "unsafe"

const (
	// IfAddressMessageSize - bytes as derived from unix.IfAddressMessageSize
	IfAddressMessageSize = int(unsafe.Sizeof(IfAddressMessage{}))
)
