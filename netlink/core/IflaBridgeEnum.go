package core

// IflaBridgeEnum - Bridge management nested attributes
//
// Bridge management nested attributes
//
//	[IFLA_AF_SPEC] = {
//	   [IflaBridgeFlags]
//	   [IFLA_BRIDGE_MODE]
//	   [IFLA_BRIDGE_VLAN_INFO]
//	}
type IflaBridgeEnum int

const (
	// IflaBridgeFlags - IFLA_BRIDGE_FLAGS
	IflaBridgeFlags IflaBridgeEnum = 0

	// IflaBridgeMode - IFLA_BRIDGE_MODE
	IflaBridgeMode IflaBridgeEnum = 1

	// IflaBridgeVlanInfo - IFLA_BRIDGE_VLAN_INFO
	IflaBridgeVlanInfo IflaBridgeEnum = 2

	// IflaBridgeVlanTunnelInfo - IFLA_BRIDGE_VLAN_TUNNEL_INFO
	IflaBridgeVlanTunnelInfo IflaBridgeEnum = 3

	// IflaBridgeMrp - IFLA_BRIDGE_MRP
	IflaBridgeMrp IflaBridgeEnum = 4

	// IflaBridgeCfm - IFLA_BRIDGE_CFM
	IflaBridgeCfm IflaBridgeEnum = 5

	// IflaBridgeMst - IFLA_BRIDGE_MST
	IflaBridgeMst IflaBridgeEnum = 6

	// IflaBridgeMax - IFLA_BRIDGE_MAX (maximum value)
	IflaBridgeMax IflaBridgeEnum = IflaBridgeMst
)
