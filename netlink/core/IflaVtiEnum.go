package core

// IflaVtiEnum - Enumeration for Virtual Tunnel Interface (VTI) attributes.
//
// These attributes correspond to the definitions in the Linux kernel for managing VTI interfaces,
// particularly for configuring tunnel links, keys, and IP addresses.
//
// References:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
type IflaVtiEnum uint8

const (
	// IflaVtiUnspec - IFLA_VTI_UNSPEC - This is an unspecified attribute, used as a placeholder for unknown or
	// default values.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaVtiUnspec IflaVtiEnum = iota

	// IflaVtiLink - IFLA_VTI_LINK - This attribute specifies the link layer (underlying interface) for the VTI.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaVtiLink

	// IflaVtiIkey - IFLA_VTI_IKEY - This attribute specifies the inbound key for IPsec encapsulation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaVtiIkey

	// IflaVtiOkey - IFLA_VTI_OKEY - This attribute specifies the outbound key for IPsec encapsulation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaVtiOkey

	// IflaVtiLocal - IFLA_VTI_LOCAL - This attribute specifies the local IP address for the VTI.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaVtiLocal

	// IflaVtiRemote - IFLA_VTI_REMOTE - This attribute specifies the remote IP address for the VTI.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaVtiRemote

	// IflaVtiMax - IFLA_VTI_MAX - This constant represents the maximum value for VTI attributes, used for validation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaVtiMax = iota - 1
)
