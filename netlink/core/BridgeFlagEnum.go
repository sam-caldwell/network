package core

// BridgeFlagEnum - Bridge Flags
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
type BridgeFlagEnum uint8

const (
	// BridgeFlagsUnspec - unspecified placeholder
	BridgeFlagsUnspec BridgeFlagEnum = 0

	//BridgeFlagsMaster -BRIDGE_FLAGS_MASTER - Bridge command to/from master
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
	BridgeFlagsMaster BridgeFlagEnum = 1

	// BridgeFlagsSelf - BRIDGE_FLAGS_SELF - Bridge command to/from lowerdev
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
	BridgeFlagsSelf BridgeFlagEnum = 2
)
