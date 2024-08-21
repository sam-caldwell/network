package core

import (
	"net"
)

// GetIPFamily returns the family type of net.IP (v4 or v6) or unspecified.
//
// Note that we only support IPv4 and IPv6 in this because that is
// what I need right now.  If you need more, please send me a PR to
// add it.
func GetIPFamily(ip net.IP) int {
	if len(ip) <= net.IPv4len || ip.To4() != nil {
		return AfInet
	}
	if len(ip) <= net.IPv6len || ip.To16() != nil {
		return AfInet6
	}
	return AfUnspec
}
