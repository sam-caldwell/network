package core

// XfrmEncapTmpl represents the `xfrm_encap_tmpl` structure used in the Linux kernel's IPsec subsystem.
// This structure is used to define an encapsulation template for IPsec tunnels, typically in the case of
// encapsulating Security Payload (ESP) with User Datagram Protocol (UDP) encapsulation to traverse NAT devices.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmEncapTmpl struct {
	// EncapType defines the type of encapsulation, such as UDP encapsulation for NAT traversal.
	// Common values are:
	// - 1: UDP encapsulation
	EncapType uint16

	// EncapSport represents the source port for encapsulation in big-endian format.
	// This port is used for encapsulated traffic. For example, when using UDP encapsulation,
	// it would define the source UDP port.
	EncapSport uint16 // big endian

	// EncapDport represents the destination port for encapsulation in big-endian format.
	// This port is used for encapsulated traffic. For example, when using UDP encapsulation,
	// it would define the destination UDP port.
	EncapDport uint16 // big endian

	// Pad is reserved padding to align the structure.
	Pad [2]byte

	// EncapOa (Original Address) represents the outer address for the encapsulation.
	// This field stores the outer IP address to be used for the encapsulated traffic.
	// It can be either IPv4 or IPv6, depending on the address family used.
	EncapOa XfrmAddress
}
