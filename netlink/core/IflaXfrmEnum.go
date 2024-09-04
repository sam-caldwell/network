package core

// IflaXfrmEnum - Enumeration for IPsec (XFRM) interface attributes.
//
// These attributes are used to configure and manage IPsec (XFRM) interfaces in the Linux kernel.
// They include parameters related to IPsec link associations and interface identifiers.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaXfrmEnum uint8

const (
	// IflaXfrmUnspec - IFLA_XFRM_UNSPEC -
	// Unspecified attribute, used as a placeholder for unknown or default values.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXfrmUnspec IflaXfrmEnum = iota

	// IflaXfrmLink - IFLA_XFRM_LINK -
	// Specifies the link layer (underlying interface) associated with the XFRM interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXfrmLink

	// IflaXfrmIfId - IFLA_XFRM_IF_ID -
	// Specifies the interface identifier for the XFRM interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXfrmIfId

	// IflaXfrmMax - IFLA_XFRM_MAX -
	// The maximum value for XFRM attributes, used for validation purposes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXfrmMax = iota - 1
)
