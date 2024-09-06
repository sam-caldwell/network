package core

import (
	"golang.org/x/sys/unix"
	"net"
)

// FromIP converts a net.IP object into the XfrmAddress format.
// This function checks whether the IP is IPv4 or IPv6, then populates
// the XfrmAddress struct accordingly. For IPv4, it pads the extra
// bytes with zeroes. For IPv6, the full 16-byte address is copied.
//
// In the Linux kernel, the xfrm_address_t union can hold either an IPv4
// or IPv6 address, so this function implements similar logic.
//
// Reference:
// - https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/xfrm.h#L29
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
func (x *XfrmAddress) FromIP(ip net.IP) {
	var empty = [16]byte{} // empty array to pad unused bytes
	if len(ip) < net.IPv4len {
		// Invalid IP length, zero out the extra bytes
		copy(x[4:16], empty[:])
	} else if GetIPFamily(ip) == unix.AF_INET {
		// If it's an IPv4 address, copy the first 4 bytes
		copy(x[0:4], ip.To4()[0:4])
		// Pad the remaining bytes with zeroes
		copy(x[4:16], empty[:12])
	} else {
		// If it's an IPv6 address, copy the entire 16 bytes
		copy(x[0:16], ip.To16()[0:16])
	}
}
