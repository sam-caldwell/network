package core

import "golang.org/x/sys/unix"

// InterfaceFamily - interface family used in IfAddressMessage (unix.IfAddrmsg/ifaddrmsg)
//
//		 See https://man7.org/linux/man-pages/man7/rtnetlink.7.html and
//	     https://github.com/torvalds/linux/blob/master/include/linux/socket.h
type InterfaceFamily uint8

const (
	// AfUnspec - AF_UNSPEC - unspecified
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfUnspec = unix.AF_UNSPEC

	// AfUnix AF_UNIX - Unix domain sockets (unix.AF_UNIX)
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfUnix = unix.AF_UNIX

	// AfInet - AF_INET - Ipv4
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfInet = unix.AF_INET

	// AfAx25 - AF_AX25 - Amateur Radio AX.25
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfAx25 = unix.AF_X25

	// AfIpx - AF_IPX - Novell IPX
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfIpx = unix.AF_IPX

	// AfAppleTalk - AfAppletalk - Appletalk DDP
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfAppleTalk = unix.AF_APPLETALK

	// AfNetRom - AF_NETROM - Amateur radio NetROM
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfNetRom = unix.AF_NETROM

	// AfBridge - AF_BRIDGE - Multi-protocol bridge
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfBridge = unix.AF_BRIDGE

	// AfAal5 - AF_AAL5 - Reserved for Werner's ATM
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfAal5 = 8 //Note Not defined in unix package

	// AfX25 - AF_X25 - Reserved for X.25 project
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfX25 = unix.AF_X25

	// AfInet6 - AF_INET6 - IPv6
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfInet6 = unix.AF_INET6

	//AfMax - alias for unix.AF_MAX (maximum value...for now)
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfMax = unix.AF_MAX

	//AfNetlink - AF_NETLINK - a protocol family in the Linux kernel used for communication between the kernel and
	//  user-space processes. It facilitates a messaging system that allows user-space applications to interact with
	//  the kernel, primarily for networking-related tasks.
	//
	// Purpose of AF_NETLINK
	//   - Kernel-User Space Communication:
	//  	Provides a socket-based interface for processes to communicate with the kernel, allowing the exchange
	//      of information related to networking, routing, and other kernel subsystems.
	//
	//   - Networking Configuration:
	//  	AF_NETLINK is extensively used for configuring and managing network interfaces, routing tables, firewall
	// 		rules, and other network-related settings. For example, utilities like iproute2 use AF_NETLINK to
	//		communicate with the kernel for configuring network parameters.
	//
	//	 - Asynchronous Messaging:
	//		It supports asynchronous message exchange, meaning messages can be sent and received without blocking,
	//		making it suitable for monitoring and reacting to changes in kernel state.
	//
	// Common Netlink Protocols
	// - NETLINK_ROUTE: Used for networking-related operations, such as managing routing tables and network interfaces.
	// - NETLINK_INET_DIAG: Provides socket monitoring and diagnostics.
	// - NETLINK_NETFILTER: Used for packet filtering and NAT (Network Address Translation) configurations.
	//
	// Example Use Cases
	//  - Network Interface Management: Applications can use AF_NETLINK to add or remove network interfaces, set IP addresses, and configure interface parameters.
	//  - Routing Table Management: Modify or query routing tables in the kernel.
	//  - Firewall Configuration: Set up or modify firewall rules using tools like iptables or nftables.
	//
	// See https://github.com/torvalds/linux/blob/master/include/linux/socket.h
	AfNetlink = 0x10
)

func (i InterfaceFamily) Int() int {
	return int(i)
}
func (i InterfaceFamily) Uint() uint {
	return uint(i)
}
func (i InterfaceFamily) Uint8() uint8 {
	return uint8(i)
}
func (i InterfaceFamily) Uint16() uint16 {
	return uint16(i)
}
func (i InterfaceFamily) Uint32() uint32 {
	return uint32(i)
}
func (i InterfaceFamily) Uint64() uint64 {
	return uint64(i)
}
