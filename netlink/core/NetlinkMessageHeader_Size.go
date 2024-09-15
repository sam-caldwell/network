package core

import (
	"unsafe"
)

// NetlinkMessageHeaderSize - NLMSG_HDRLEN - Size of NlMsghdr struct//
// References:
// - Netlink definitions: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
// - Routing Netlink definitions: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
const NetlinkMessageHeaderSize = int(unsafe.Sizeof(NlMsghdr{}))
