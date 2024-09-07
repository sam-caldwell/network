//go:build linux

package core

import (
	"golang.org/x/sys/unix"
)

// IfAddressMessage - wrapper around unix.IfAddrmsg in case golang will
// eventually want to make changes here, and I need to minimize impact.
//
// // unix.IfAddrmsg looks like this...
//
//	type IfAddrmsg struct {
//		Family    uint8
//		Prefixlen uint8
//		Flags     uint8
//		Scope     uint8
//		Index     uint32
//	}
//
// Note:
//
//	We have some stricter types in this package (e.g. InterfaceFamily) to help
//	developers quickly identify relevant constants without having to resort
//	to Linux man pages.
type IfAddressMessage struct {
	unix.IfAddrmsg
}
