package core

import "golang.org/x/sys/unix"

// IfFamily - interface family used in IfAddressMessage (unix.IfAddrmsg/ifaddrmsg)
//
// See https://man7.org/linux/man-pages/man7/rtnetlink.7.html.
//
// "ifa_family is the address family type (currently AfInet or AfInet6)"
type IfFamily uint8

const (
	// AfInet - Ipv4 - alias for unix.AF_INET
	AfInet = unix.AF_INET

	// AfInet6 - IPv6 - alias for unix.AF_INET6
	AfInet6 = unix.AF_INET6
)
