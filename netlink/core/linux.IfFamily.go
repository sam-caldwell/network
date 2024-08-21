package core

import "golang.org/x/sys/unix"

// IfFamily - interface family used in IfAddressMessage (unix.IfAddrmsg/ifaddrmsg)
//
// See https://man7.org/linux/man-pages/man7/rtnetlink.7.html.
//
// "ifa_family is the address family type (currently AfInet or AfInet6)"
type IfFamily uint8

const (
	// AfUnspec - AF_UNSPEC - unspecified
	AfUnspec = unix.AF_UNSPEC

	// AfUnix AF_UNIX - Unix domain sockets (unix.AF_UNIX)
	AfUnix = unix.AF_UNIX

	// AfInet - AF_INET - Ipv4
	AfInet = unix.AF_INET

	// AfAx25 - AF_AX25 - Amateur Radio AX.25
	AfAx25 = unix.AF_X25

	// AfIpx - AF_IPX - Novell IPX
	AfIpx = unix.AF_IPX

	// AfAppleTalk - AF_APPLETALK - Appletalk DDP
	AfAppleTalk = unix.AF_APPLETALK

	// AfNetRom - AF_NETROM - Amateur radio NetROM
	AfNetRom = unix.AF_NETROM

	// AfBridge - AF_BRIDGE - Multi-protocol bridge
	AfBridge = unix.AF_BRIDGE

	// AfAal5 - AF_AAL5 - Reserved for Werner's ATM
	AfAal5 = 8 //Note Not defined in unix package

	// AfX25 - AF_X25 - Reserved for X.25 project
	AfX25 = unix.AF_X25

	// AfInet6 - AF_INET6 - IPv6
	AfInet6 = unix.AF_INET6

	//AfMax - alias for unix.AF_MAX (maximum value...for now)
	AfMax = unix.AF_MAX
)
