package core

// IflaGtpEnum - Enumeration for GTP (GPRS Tunneling Protocol) interface attributes.
//
// These attributes are used to configure and retrieve settings for GTP interfaces in the Linux kernel.
// GTP is primarily used in mobile networks for transmitting data over GPRS, LTE, and 5G.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
type IflaGtpEnum uint8

const (
	// IflaGtpUnspec - IFLA_GTP_UNSPEC - Unspecified attribute, used as a placeholder.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGtpUnspec IflaGtpEnum = iota

	// IflaGtpFd0 - IFLA_GTP_FD0 - Specifies the file descriptor for the GTP-U socket (GTPv1-U).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGtpFd0

	// IflaGtpFd1 - IFLA_GTP_FD1 - Specifies the file descriptor for the GTP-C socket (GTPv1-C).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGtpFd1

	// IflaGtpPdpHashsize - IFLA_GTP_PDP_HASHSIZE - Defines the size of the hash table used for storing
	// PDP (Packet Data Protocol) contexts.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGtpPdpHashsize

	// IflaGtpRole - IFLA_GTP_ROLE - Defines the role of the GTP instance, such as GGSN or SGSN in a mobile network.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_tunnel.h
	IflaGtpRole
)
