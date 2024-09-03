package core

import "golang.org/x/sys/unix"

// IfRtaEnum - rtnetlink attribute option type.
//
//		See https://man7.org/linux/man-pages/man7/rtnetlink.7.html.
//
//	 Some rtnetlink messages have optional attributes after the
//	 initial header:
type IfRtaEnum uint8

const (
	//IflaUnspec - IFLA_UNSPEC - unspecified
	IflaUnspec IfRtaEnum = unix.IFLA_UNSPEC

	//IflaAddress - IFLA_ADDRESS - hardware address   interface L2 address
	IflaAddress IfRtaEnum = unix.IFLA_ADDRESS

	//IflaBroadcast - IFLA_BROADCAST - hardware address   L2 broadcast address
	IflaBroadcast IfRtaEnum = unix.IFLA_BROADCAST

	//IflaIfname - IFLA_IFNAME - asciiz string      Device name
	IflaIfname IfRtaEnum = unix.IFLA_IFNAME

	//IflaMtu - IFLA_MTU - unsigned int-  MTU of the device
	IflaMtu IfRtaEnum = unix.IFLA_MTU

	//IflaLink - IFLA_LINK - int- Link type
	IflaLink IfRtaEnum = unix.IFLA_LINK

	//IflaQdisc - IFLA_QDISC - asciiz string      Queueing discipline
	IflaQdisc IfRtaEnum = unix.IFLA_QDISC

	//IflaStats - IFLA_STATS - see below - Interface Statistics
	IflaStats IfRtaEnum = unix.IFLA_STATS

	//IflaPermAddress IFLA_PERM_ADDRESS- hardware address - hardware address provided by device (since Linux 5.5)
	IflaPermAddress IfRtaEnum = unix.IFLA_PERM_ADDRESS
)
