package core

// XdpAttachedEnum - XDP program attach mode
//
// This enumeration defines the different modes in which an XDP (eXpress Data Path) program can be attached to a network interface.
// These values are used as dump values for `IFLA_XDP_ATTACHED` to indicate the current attachment mode of the XDP program.
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type XdpAttachedEnum uint8

const (
	// XdpAttachedNone - XDP_ATTACHED_NONE -
	// Indicates that no XDP program is currently attached to the network interface.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XdpAttachedNone XdpAttachedEnum = 0

	// XdpAttachedDrv - XDP_ATTACHED_DRV -
	// Indicates that an XDP program is attached in driver mode.
	// Driver mode runs the XDP program at a low level, directly within the network driver, providing high performance.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XdpAttachedDrv XdpAttachedEnum = 1

	// XdpAttachedSkb - XDP_ATTACHED_SKB -
	// Indicates that an XDP program is attached in SKB (socket buffer) mode.
	// SKB mode runs the XDP program at a higher level in the network stack, offering more compatibility with existing features.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XdpAttachedSkb XdpAttachedEnum = 2

	// XdpAttachedHw - XDP_ATTACHED_HW -
	// Indicates that an XDP program is attached in hardware mode.
	// Hardware mode offloads the XDP processing to network hardware (e.g., NIC), providing the highest performance.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XdpAttachedHw XdpAttachedEnum = 3

	// XdpAttachedMulti - XDP_ATTACHED_MULTI -
	// Indicates that multiple XDP programs are attached in different modes.
	// This is used when an interface has multiple XDP programs attached, potentially across different modes (e.g., driver and hardware).
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XdpAttachedMulti XdpAttachedEnum = 4
)
