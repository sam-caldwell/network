package core

import (
	"bytes"
	"net"
)

// ToIP converts the XfrmAddress to a net.IP structure.
//
// This function checks if the address represents an IPv4 address by examining if the
// last 12 bytes (index 4-16) are empty. If so, it constructs an IPv4-mapped IPv6 address
// (with the 96-bit prefix `::ffff:` followed by the IPv4 address).
//
// If it's an IPv6 address, the function directly copies the 16 bytes into a net.IP structure.
//
// Returns:
// - `net.IP`: The corresponding IP address (IPv4 or IPv6).
//
// References:
// - Linux Kernel: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
func (x *XfrmAddress) ToIP() net.IP {
	// Empty array to check for IPv4 addresses stored in the IPv6 format.
	var empty = [12]byte{}

	// Initialize a net.IP slice to store the IPv6 address.
	ip := make(net.IP, net.IPv6len)

	// If the last 12 bytes of the XfrmAddress are zeroed, it's an IPv4 address.
	if bytes.Equal(x[4:16], empty[:]) {
		// Create an IPv4-mapped IPv6 address (::ffff:<IPv4>).
		ip[10] = 0xff
		ip[11] = 0xff
		copy(ip[12:16], x[0:4])
	} else {
		// Otherwise, it's a native IPv6 address, so copy the entire address.
		copy(ip[:], x[:])
	}

	return ip
}
