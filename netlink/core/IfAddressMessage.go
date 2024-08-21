//go:build linux

package core

import (
	"golang.org/x/sys/unix"
)

// IfAddressMessage - wrapper around unix.IfAddrmsg in case golang will
// eventually want to make changes here, and I need to minimize impact.
//
// The C implementation...
//
//	struct ifaddrmsg {
//	  __u8    ifa_family;
//	  __u8    ifa_prefixlen;  /* The prefix length    */
//	  __u8    ifa_flags;  /* Flags      */
//	  __u8    ifa_scope;  /* Address scope    */
//	  __u32   ifa_index;  /* Link index     */
//	};
//
// unix.IfAddrmsg looks like this...
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
//	We have some stricter types in this package (e.g. IfFamily) to help
//	developers quickly identify relevant constants without having to resort
//	to Linux man pages.
type IfAddressMessage struct {
	unix.IfAddrmsg
}
