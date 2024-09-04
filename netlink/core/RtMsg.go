package core

import (
	"golang.org/x/sys/unix"
)

// RtMsg - A wrapper around the unix.RtMsg structure.
//
// The RtMsg structure is used in conjunction with routing messages in the Linux kernel's Netlink API.
// This structure represents a route entry in the routing table and is used when performing operations
// like adding, deleting, or modifying routes.
//
// The primary fields include information on routing family, destination length, source length, and flags
// which control the behavior of route messages.
//
// For more details, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h,
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
type RtMsg struct {
	unix.RtMsg
}
