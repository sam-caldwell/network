package core

import (
	"golang.org/x/sys/unix"
)

// RtGenMsg - Wrapper for the unix.RtGenmsg structure used in rtnetlink communications.
//
// RtGenMsg represents a generic RTNetlink message used when sending or receiving routing-related information between user-space
// and the Linux kernel. It is primarily used to request or retrieve information about network routes, links, and addresses.
// The structure is small and only includes the family of the address it applies to.
//
// The C struct looks like:
//
//	struct rtgenmsg {
//	   unsigned char rtgen_family; /* AF_UNSPEC */
//	};
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
type RtGenMsg struct {
	unix.RtGenmsg
}
