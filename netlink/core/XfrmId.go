package core

// XfrmId represents the unique identifier for an IPsec Security Association (SA).
// It corresponds to the xfrm_id structure in the Linux kernel, which identifies a specific
// SA by its destination address, SPI (Security Parameter Index), and protocol.
//
// The structure is used in IPsec processing to match inbound or outbound traffic with a specific SA.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h

type XfrmId struct {
	// Daddr represents the destination IP address (IPv4 or IPv6) of the IPsec SA.
	Daddr XfrmAddress

	// Spi (Security Parameter Index) is a unique identifier for the SA.
	// It is used to distinguish different SAs between the same pair of peers.
	// The SPI is stored in big-endian format.
	Spi uint32 // big-endian

	// Proto specifies the IPsec protocol (e.g., ESP or AH) associated with the SA.
	// The protocol is represented as an 8-bit value.
	Proto uint8

	// Pad is a 3-byte padding used to align the structure to a 4-byte boundary.
	// This ensures proper alignment and memory handling on different platforms.
	Pad [3]byte
}
