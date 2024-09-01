package core

// BridgeFlag - Bridge Flags : https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
type BridgeFlag int

const (
	//BridgeFlagsMaster -BRIDGE_FLAGS_MASTER - Bridge command to/from master
	BridgeFlagsMaster BridgeFlag = 1

	// BridgeFlagsSelf - BRIDGE_FLAGS_SELF - Bridge command to/from lowerdev
	BridgeFlagsSelf BridgeFlag = 2
)
