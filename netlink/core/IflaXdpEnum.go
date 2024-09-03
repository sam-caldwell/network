package core

// IflaXdpEnum - Enumeration for XDP (eXpress Data Path) attributes
//
// This enumeration defines various attributes related to XDP (eXpress Data Path) in Linux.
//
// XDP is a high-performance, programmable network data path in the Linux kernel, allowing for packet processing
// at the lowest levels of the network stack.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaXdpEnum uint8

const (
	// IflaXdpUnspec - IFLA_XDP_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	// It represents an undefined or unused attribute within the context of XDP.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXdpUnspec IflaXdpEnum = 0

	// IflaXdpFd - IFLA_XDP_FD -
	// This attribute specifies the file descriptor for the XDP program.
	// The file descriptor is used to load or attach an XDP program to a network interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXdpFd IflaXdpEnum = 1

	// IflaXdpAttached - IFLA_XDP_ATTACHED -
	// This attribute indicates whether an XDP program is currently attached to a network interface.
	// It can be used to check the status of the XDP program on the interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXdpAttached IflaXdpEnum = 2

	// IflaXdpFlags - IFLA_XDP_FLAGS -
	// This attribute specifies various flags for the XDP program.
	// Flags control how the XDP program operates, such as whether it runs in hardware or software mode.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXdpFlags IflaXdpEnum = 3

	// IflaXdpProgId - IFLA_XDP_PROG_ID -
	// This attribute provides the ID of the loaded XDP program.
	// The program ID is a unique identifier used to reference the XDP program within the kernel.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXdpProgId IflaXdpEnum = 4

	// IflaXdpDrvProgId - IFLA_XDP_DRV_PROG_ID -
	// This attribute provides the ID of the XDP program running in driver mode.
	// Driver mode allows the XDP program to be run in the context of the network driver for high performance.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXdpDrvProgId IflaXdpEnum = 5

	// IflaXdpSkbProgId - IFLA_XDP_SKB_PROG_ID -
	// This attribute provides the ID of the XDP program running in SKB (socket buffer) mode.
	// SKB mode processes packets at a higher level in the kernel stack, offering compatibility with more features.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXdpSkbProgId IflaXdpEnum = 6

	// IflaXdpHwProgId - IFLA_XDP_HW_PROG_ID -
	// This attribute provides the ID of the XDP program running in hardware mode.
	// Hardware mode offloads the XDP processing to network hardware, such as a NIC, for better performance.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXdpHwProgId IflaXdpEnum = 7

	// IflaXdpExpectedFd - IFLA_XDP_EXPECTED_FD -
	// This attribute specifies the expected file descriptor for an XDP program.
	// It is used during program replacement to ensure that the correct program is being replaced or updated.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXdpExpectedFd IflaXdpEnum = 8

	// IflaXdpMax - IFLA_XDP_MAX -
	// This constant represents the maximum valid value for XDP attributes.
	// It is used as a boundary marker to validate or limit the range of XDP attributes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaXdpMax IflaXdpEnum = iota - 1
)
