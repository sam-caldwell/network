package core

// XfrmUserTmpl represents the xfrm_user_tmpl structure used in the Linux kernel's XFRM subsystem for IPsec.
//
// The xfrm_user_tmpl structure defines a template for creating a Security Association (SA).
// It specifies various attributes of the SA such as source and destination addresses, the family (IPv4/IPv6),
// encryption and authentication algorithms, and the mode of operation.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/xfrm.h
type XfrmUserTmpl struct {
	// XfrmId represents the unique identifier for the Security Association (SA), including
	// the destination IP address, SPI (Security Parameter Index), and protocol (e.g., AH or ESP).
	// This field corresponds to the `xfrm_id` structure in the kernel.
	XfrmId XfrmId

	// Family defines the address family of the SA (e.g., AF_INET for IPv4 or AF_INET6 for IPv6).
	Family uint16

	// Pad1 is a 2-byte padding field to ensure correct alignment in the structure.
	Pad1 [2]byte

	// Saddr is the source IP address for the SA. It is of type XfrmAddress, which can hold either an
	// IPv4 or IPv6 address.
	Saddr XfrmAddress

	// Reqid is a unique identifier used to match policies to Security Associations.
	// This field corresponds to the `reqid` field in the kernel.
	Reqid uint32

	// Mode defines the mode of the Security Association (e.g., transport mode or tunnel mode).
	// This field corresponds to the `xfrm_mode` field in the kernel (e.g., XFRM_MODE_TRANSPORT or XFRM_MODE_TUNNEL).
	Mode uint8

	// Share defines the sharing mode of the SA (e.g., how the SA is shared between different users).
	// This field corresponds to the `xfrm_share` field in the kernel.
	Share uint8

	// Optional indicates whether the SA is optional. If set, the kernel may ignore this template when
	// setting up an SA if a suitable match cannot be found.
	Optional uint8

	// Pad2 is a single-byte padding field for alignment purposes.
	Pad2 byte

	// Aalgos is a bitmask representing the supported authentication algorithms for this SA.
	// This field corresponds to the `aalgos` field in the kernel.
	Aalgos uint32

	// Ealgos is a bitmask representing the supported encryption algorithms for this SA.
	// This field corresponds to the `ealgos` field in the kernel.
	Ealgos uint32

	// Calgos is a bitmask representing the supported compression algorithms for this SA.
	// This field corresponds to the `calgos` field in the kernel.
	Calgos uint32
}
