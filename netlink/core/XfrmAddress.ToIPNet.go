package core

import (
	"bytes"
	"golang.org/x/sys/unix"
	"net"
)

// ToIPNet converts the XfrmAddress to a *net.IPNet structure, representing an IP network.
//
// - `prefixlen` is the length of the network mask (in bits) for the IPNet structure.
// - `family` is used to determine if the IP is IPv4 or IPv6 when the address is empty (zeroed).
//
// The function handles two cases:
// 1. If the XfrmAddress and `prefixlen` are both 0, the address will be set based on the `family`:
//   - For `unix.AF_INET` (IPv4), it returns an IPNet with address "0.0.0.0".
//   - For `unix.AF_INET6` (IPv6), it returns an IPNet with address "::" (unspecified IPv6 address).
//
// 2. Otherwise, the XfrmAddress is converted to an IP, and the network mask is set based on the IP family:
//   - IPv4 addresses get a mask length of 32 bits.
//   - IPv6 addresses get a mask length of 128 bits.
//
// References:
// - Linux Kernel Source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
func (x *XfrmAddress) ToIPNet(prefixlen uint8, family uint16) *net.IPNet {
	empty := [SizeofXfrmAddress]byte{}

	// If the address is empty and prefixlen is 0, return based on the family type (IPv4/IPv6).
	if bytes.Equal(x[:], empty[:]) && prefixlen == 0 {
		if family == unix.AF_INET {
			// Return an IPv6 address "::" with a /128 prefix (default route)
			return &net.IPNet{IP: net.ParseIP("::"), Mask: net.CIDRMask(int(prefixlen), 128)}
		}
		// Return an IPv4 address "0.0.0.0" with a /32 prefix (default route)
		return &net.IPNet{IP: net.ParseIP("0.0.0.0"), Mask: net.CIDRMask(int(prefixlen), 32)}
	}

	// Convert XfrmAddress to net.IP and determine its family (IPv4 or IPv6)
	ip := x.ToIP()

	// For IPv4, return IPNet with a 32-bit mask (default route)
	if GetIPFamily(ip) == unix.AF_INET {
		return &net.IPNet{IP: ip, Mask: net.CIDRMask(int(prefixlen), 32)}
	}

	// For IPv6, return IPNet with a 128-bit mask (default route)
	return &net.IPNet{IP: ip, Mask: net.CIDRMask(int(prefixlen), 128)}
}
