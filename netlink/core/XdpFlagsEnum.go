package core

// XdpFlagsEnum - Enumeration for XDP (eXpress Data Path) flags
//
// This enumeration defines various flags used in the configuration and operation of XDP (eXpress Data Path) in Linux.
// XDP is a high-performance packet processing path that allows for fast packet handling at the lowest level of the network stack.
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type XdpFlagsEnum uint8

const (
	// XDPFlagsUpdateIfNoExist - XDP_FLAGS_UPDATE_IF_NOEXIST -
	// This flag ensures that the XDP program is only attached to the network interface if no XDP program is currently attached.
	// It prevents overwriting an existing XDP program on the interface.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XDPFlagsUpdateIfNoExist XdpFlagsEnum = iota << 0

	// XDPFlagsSkbMode - XDP_FLAGS_SKB_MODE -
	// This flag forces the XDP program to run in SKB (socket buffer) mode.
	// SKB mode is compatible with more features but is generally slower than running in native mode.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XDPFlagsSkbMode XdpFlagsEnum = iota << 1

	// XDPFlagsDrvMode - XDP_FLAGS_DRV_MODE -
	// This flag forces the XDP program to run in driver mode.
	// Driver mode is faster and more efficient than SKB mode as it operates at a lower level in the network stack.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XDPFlagsDrvMode XdpFlagsEnum = iota << 2

	// XdpFlagsHwMode - XDP_FLAGS_HW_MODE -
	// This flag forces the XDP program to run in hardware mode, offloading the processing to the network hardware (e.g., NIC).
	// Hardware mode provides the highest performance by leveraging specialized network hardware.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XdpFlagsHwMode XdpFlagsEnum = iota << 3

	// XdpFlagsReplace - XDP_FLAGS_REPLACE -
	// This flag allows replacing an existing XDP program on a network interface.
	// It ensures that the existing program is replaced by the new one, useful for updates or changes in configurations.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XdpFlagsReplace XdpFlagsEnum = iota << 4

	// XdpFlagsModes - XDP_FLAGS_MODES -
	// This constant represents a bitmask of all mode-related XDP flags (SKB mode, driver mode, and hardware mode).
	// It is used to check or apply multiple mode flags simultaneously.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XdpFlagsModes = XDPFlagsSkbMode | XDPFlagsDrvMode | XdpFlagsHwMode

	// XDPFlagsMask - XDP_FLAGS_MASK -
	// This constant represents a bitmask of all valid XDP flags.
	// It is used to validate or check for the presence of any XDP-related flags in a configuration.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	XDPFlagsMask = XDPFlagsUpdateIfNoExist | XDPFlagsSkbMode | XDPFlagsDrvMode
)
