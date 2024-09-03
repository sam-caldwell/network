package core

// BridgeModeEnum - Bridge mode enum definition
//
//	Controls whether traffic may be sent back out of the port on which it
//	was received. This option is also called reflective relay mode, and is
//	used to support basic VEPA (Virtual Ethernet Port Aggregator)
//	capabilities. By default, this flag is turned off and the bridge will
//	not forward traffic back out of the receiving port.
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type BridgeModeEnum uint8

const (
	// BridgeModeUnspec - BRIDGE_MODE_UNSPEC -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	BridgeModeUnspec BridgeModeEnum = 0

	// BridgeModeHairpin - BRIDGE_MODE_HAIRPIN -
	//
	//	Controls whether traffic may be sent back out of the port on which it
	//	was received. This option is also called reflective relay mode, and is
	//	used to support basic VEPA (Virtual Ethernet Port Aggregator)
	//	capabilities. By default, this flag is turned off and the bridge will
	//	not forward traffic back out of the receiving port.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	BridgeModeHairpin BridgeModeEnum = 1
)
