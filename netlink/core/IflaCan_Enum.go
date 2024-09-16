package core

// IflaCan_Enum - Enumeration for CAN (Controller Area Network) interface attributes.
//
// These attributes are used to configure and manage CAN interfaces in the Linux kernel.
// CAN is a robust vehicle bus standard designed to allow microcontrollers and devices to communicate with each
// other in applications without a host computer.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
type IflaCan_Enum uint8

const (
	// IflaCanUnspec - IFLA_CAN_UNSPEC - Unspecified attribute, used as a placeholder for unknown or default values.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanUnspec IflaCan_Enum = 0

	// IflaCanBittiming - IFLA_CAN_BITTIMING - Specifies the bit timing parameters for the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanBittiming IflaCan_Enum = 1

	// IflaCanBittimingConst - IFLA_CAN_BITTIMING_CONST - Specifies the bit timing constants for the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanBittimingConst IflaCan_Enum = 2

	// IflaCanClock - IFLA_CAN_CLOCK - Specifies the clock frequency for the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanClock IflaCan_Enum = 3

	// IflaCanState - IFLA_CAN_STATE - Specifies the state of the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanState IflaCan_Enum = 4

	// IflaCanCtrlmode - IFLA_CAN_CTRLMODE - Specifies the control mode of the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanCtrlmode IflaCan_Enum = 5

	// IflaCanRestartMs - IFLA_CAN_RESTART_MS - Specifies the restart delay in milliseconds for the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanRestartMs IflaCan_Enum = 6

	// IflaCanRestart - IFLA_CAN_RESTART - Requests a restart of the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanRestart IflaCan_Enum = 7

	// IflaCanBerrCounter - IFLA_CAN_BERR_COUNTER - Specifies the error counter for the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanBerrCounter IflaCan_Enum = 8

	// IflaCanDataBittiming - IFLA_CAN_DATA_BITTIMING - Specifies the data bit timing parameters for CAN FD
	// (Flexible Data-Rate) frames.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanDataBittiming IflaCan_Enum = 9

	// IflaCanDataBittimingConst - IFLA_CAN_DATA_BITTIMING_CONST - Specifies the data bit timing constants for
	// CAN FD frames.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanDataBittimingConst IflaCan_Enum = 10

	// IflaCanTermination - IFLA_CAN_TERMINATION - Specifies the termination resistor value for the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanTermination IflaCan_Enum = 11

	// IflaCanTerminationConst - IFLA_CAN_TERMINATION_CONST - Specifies the termination resistor constants for the
	// CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanTerminationConst IflaCan_Enum = 12

	// IflaCanBitrateConst - IFLA_CAN_BITRATE_CONST - Specifies the constant bitrates supported by the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanBitrateConst IflaCan_Enum = 13

	// IflaCanDataBitrateConst - IFLA_CAN_DATA_BITRATE_CONST - Specifies the constant data bitrates supported by the
	// CAN FD interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanDataBitrateConst IflaCan_Enum = 14

	// IflaCanBitrateMax - IFLA_CAN_BITRATE_MAX - Specifies the maximum bitrate supported by the CAN interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanBitrateMax IflaCan_Enum = 15

	// IflaCanMax - IFLA_CAN_MAX - The maximum value for CAN attributes, used for validation purposes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/can/netlink.h
	IflaCanMax = iota - 1
)
