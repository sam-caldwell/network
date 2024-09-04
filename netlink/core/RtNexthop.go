package core

import (
	"golang.org/x/sys/unix"
)

// RtNexthop - rtnexthop - wraps the unix.RtNexthop structure and adds a Children field, which is a slice of
// NetlinkRequestData. This is useful for representing a route's next-hop attributes in the Linux kernel's routing
// table and allows for the inclusion of additional netlink request data. This structure is often used to define
// routing rules where multiple next-hop destinations are possible.
//
// You would use Children to store additional netlink requests related to the specific next-hop, enabling more
// complex routing configurations.
//
// See https://man7.org/linux/man-pages/man7/rtnetlink.7.html,
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
type RtNexthop struct {
	unix.RtNexthop

	Children []NetlinkRequestData
}
