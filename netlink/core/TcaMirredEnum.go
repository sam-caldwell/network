package core

// Constants for the mirred action in the Linux traffic control system.
// These are used for redirecting or mirroring packets to different network interfaces.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
const (
	// TcaActMirred - TCA_ACT_MIRRED - Action ID for the mirred action.
	// The value 8 corresponds to the mirred action, used for redirecting or mirroring packets.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
	TcaActMirred = 8
)

// TcaMirredEnum - Enumeration for mirred action attributes.
//
// These constants represent different attributes for the mirred action,
// which is used in traffic control for redirecting or mirroring packets.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
type TcaMirredEnum uint8

const (
	// TcaMirredUnspec - TCA_MIRRED_UNSPEC - Unspecified action.
	// Placeholder for unspecified or unknown mirred action.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
	TcaMirredUnspec TcaMirredEnum = iota

	// TcaMirredTm - TCA_MIRRED_TM - Time management for mirred action.
	// This is used to manage timing for mirred actions.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
	TcaMirredTm

	// TcaMirredParms - TCA_MIRRED_PARMS - Parameters for the mirred action.
	// Used to specify parameters for redirecting or mirroring actions.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
	TcaMirredParms

	// TcaMirredMax - TCA_MIRRED_MAX - Maximum attribute for mirred actions.
	// Defines the maximum attribute in this enumeration for validation purposes.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
	TcaMirredMax = TcaMirredParms
)
