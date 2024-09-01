package core

// BridgeVlanInfoEnum - https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_bridge.h
type BridgeVlanInfoEnum uint16

const (
	// BridgeVlanInfoMaster - BRIDGE_VLAN_INFO_MASTER - Operate on Bridge device as well
	BridgeVlanInfoMaster BridgeVlanInfoEnum = 1

	// BridgeVlanInfoPvid - BRIDGE_VLAN_INFO_PVID -  VLAN is PVID, ingress untagged
	BridgeVlanInfoPvid BridgeVlanInfoEnum = 2

	// BridgeVlanInfoUntagged - BRIDGE_VLAN_INFO_UNTAGGED - VLAN egresses untagged
	BridgeVlanInfoUntagged BridgeVlanInfoEnum = 4

	// BridgeVlanInfoRangeBegin - BRIDGE_VLAN_INFO_RANGE_BEGIN - VLAN is start of vlan range
	BridgeVlanInfoRangeBegin BridgeVlanInfoEnum = 8

	// BridgeVlanInfoRangeEnd - BRIDGE_VLAN_INFO_RANGE_END - VLAN is end of vlan range
	BridgeVlanInfoRangeEnd BridgeVlanInfoEnum = 16

	// BridgeVlanInfoBrentry - BRIDGE_VLAN_INFO_BRENTRY - Global bridge VLAN entry
	BridgeVlanInfoBrentry BridgeVlanInfoEnum = 32

	// BridgeVlanInfoOnlyOpts - BRIDGE_VLAN_INFO_ONLY_OPTS - Skip create/delete/flags
	BridgeVlanInfoOnlyOpts BridgeVlanInfoEnum = 64
)
