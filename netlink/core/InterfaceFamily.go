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

	// AfAppleTalk - AF_APPLETALK - Appletalk DDP
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
)
