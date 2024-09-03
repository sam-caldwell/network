package core

import "golang.org/x/sys/unix"

// IfRtaEnum - rtnetlink attribute option type.
//
//		See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
//
//	 Some rtnetlink messages have optional attributes after the
//	 initial header:
type IfRtaEnum uint8

const (
	//IflaUnspec - IFLA_UNSPEC - unspecified
	//
	//		See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaUnspec IfRtaEnum = unix.IFLA_UNSPEC

	//IflaAddress - IFLA_ADDRESS - hardware address   interface L2 address
	//
	//		See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaAddress IfRtaEnum = unix.IFLA_ADDRESS

	//IflaBroadcast - IFLA_BROADCAST - hardware address - L2 broadcast address
	//
	//		See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBroadcast IfRtaEnum = unix.IFLA_BROADCAST

	//IflaIfname - IFLA_IFNAME - asciiz string - Device name
	//
	//		See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIfname IfRtaEnum = unix.IFLA_IFNAME

	//IflaMtu - IFLA_MTU - unsigned int - MTU of the device
	//
	//		See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMtu IfRtaEnum = unix.IFLA_MTU

	//IflaLink - IFLA_LINK - int- Link type
	//
	//		See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaLink IfRtaEnum = unix.IFLA_LINK

	//IflaQdisc - IFLA_QDISC - asciiz string - Queueing discipline
	//
	//		See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaQdisc IfRtaEnum = unix.IFLA_QDISC

	//IflaStats - IFLA_STATS - see below - Interface Statistics
	//
	//		See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaStats IfRtaEnum = unix.IFLA_STATS

	//IflaPermAddress IFLA_PERM_ADDRESS - hardware address - hardware address provided by device (since Linux 5.5)
	//
	//		See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaPermAddress IfRtaEnum = unix.IFLA_PERM_ADDRESS
)
