package core

// IflaBridge - Bridge management nested attributes
//
// Bridge management nested attributes
//
//	[IFLA_AF_SPEC] = {
//	   [IflaBridgeFlags]
//	   [IFLA_BRIDGE_MODE]
//	   [IFLA_BRIDGE_VLAN_INFO]
//	}
type IflaBridge int

const (
	// IflaBridgeFlags - IFLA_BRIDGE_FLAGS
	IflaBridgeFlags IflaBridge = 0

	// IflaBridgeMode - IFLA_BRIDGE_MODE
	IflaBridgeMode IflaBridge = 1

	// IflaBridgeVlanInfo - IFLA_BRIDGE_VLAN_INFO
	IflaBridgeVlanInfo IflaBridge = 2

	// IflaBridgeVlanTunnelInfo - IFLA_BRIDGE_VLAN_TUNNEL_INFO
	IflaBridgeVlanTunnelInfo IflaBridge = 3

	// IflaBridgeMrp - IFLA_BRIDGE_MRP
	IflaBridgeMrp IflaBridge = 4

	// IflaBridgeCfm - IFLA_BRIDGE_CFM
	IflaBridgeCfm IflaBridge = 5

	// IflaBridgeMst - IFLA_BRIDGE_MST
	IflaBridgeMst IflaBridge = 6

	// IflaBridgeMax - IFLA_BRIDGE_MAX (maximum value)
	IflaBridgeMax IflaBridge = IflaBridgeMst
)
