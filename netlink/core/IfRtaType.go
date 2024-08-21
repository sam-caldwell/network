package core

import "golang.org/x/sys/unix"

// IfRtaType - rtnetlink attribute option type.
//
//		See https://man7.org/linux/man-pages/man7/rtnetlink.7.html.
//
//	 Some rtnetlink messages have optional attributes after the
//	 initial header:
//
//			struct rtattr {
//	     	unsigned short rta_len;    /* Length of option */
//	         unsigned short rta_type;   /* Type of option */
//	         /* Data follows */
//	 	};
type IfRtaType int

const (
	//IflaUnspec - IFLA_UNSPEC - unspecified
	IflaUnspec IfRtaType = unix.IFLA_UNSPEC

	//IflaAddress - IFLA_ADDRESS - hardware address   interface L2 address
	IflaAddress IfRtaType = unix.IFLA_ADDRESS

	//IflaBroadcast - IFLA_BROADCAST - hardware address   L2 broadcast address
	IflaBroadcast IfRtaType = unix.IFLA_BROADCAST

	//IflaIfname - IFLA_IFNAME - asciiz string      Device name
	IflaIfname IfRtaType = unix.IFLA_IFNAME

	//IflaMtu - IFLA_MTU - unsigned int-  MTU of the device
	IflaMtu IfRtaType = unix.IFLA_MTU

	//IflaLink - IFLA_LINK - int- Link type
	IflaLink IfRtaType = unix.IFLA_LINK

	//IflaQdisc - IFLA_QDISC - asciiz string      Queueing discipline
	IflaQdisc IfRtaType = unix.IFLA_QDISC

	//IflaStats - IFLA_STATS - see below - Interface Statistics
	IflaStats IfRtaType = unix.IFLA_STATS

	//IflaPermAddress IFLA_PERM_ADDRESS- hardware address - hardware address provided by device (since Linux 5.5)
	IflaPermAddress IfRtaType = unix.IFLA_PERM_ADDRESS
)
