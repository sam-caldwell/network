package core

// XfrmUsersaId represents the `xfrm_usersa_id` structure used in the Linux kernel's XFRM subsystem.
// This structure is used to define the identifier for an IPsec Security Association (SA). It includes
// details such as the destination address (Daddr), Security Parameter Index (SPI), and the protocol.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmUsersaId struct {
	// Daddr represents the destination IP address for the Security Association (SA).
	// It can be an IPv4 or IPv6 address, depending on the family.
	Daddr XfrmAddress

	// Spi (Security Parameter Index) is a unique identifier for the Security Association (SA).
	// It is a 32-bit value in big-endian format, used to distinguish between multiple SAs.
	Spi uint32 // big endian

	// Family defines the address family of the Security Association (SA), which can be either AF_INET (IPv4) or AF_INET6 (IPv6).
	// This field helps the kernel identify the type of IP address (IPv4 or IPv6).
	Family uint16

	// Proto defines the protocol associated with the Security Association (SA). For example, it can be ESP (Encapsulating Security Payload)
	// or AH (Authentication Header). This is represented as an 8-bit unsigned integer.
	Proto uint8

	// Pad is a 1-byte padding field used for memory alignment purposes.
	Pad byte
}
