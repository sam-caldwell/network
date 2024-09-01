package core

// IfLaBridge - Bridge management nested attributes
//
// Bridge management nested attributes
//
//	[IFLA_AF_SPEC] = {
//	   [IflaBridgeFlags]
//	   [IFLA_BRIDGE_MODE]
//	   [IFLA_BRIDGE_VLAN_INFO]
//	}
type IfLaBridge int

const (
	// IflaBridgeFlags - IFLA_BRIDGE_FLAGS
	IflaBridgeFlags IfLaBridge = 0

	// IflaBridgeMode - IFLA_BRIDGE_MODE
	IflaBridgeMode IfLaBridge = 1

	// IflaBridgeVlanInfo - IFLA_BRIDGE_VLAN_INFO
	IflaBridgeVlanInfo IfLaBridge = 2

	// IflaBridgeVlanTunnelInfo - IFLA_BRIDGE_VLAN_TUNNEL_INFO
	IflaBridgeVlanTunnelInfo IfLaBridge = 3

	// IflaBridgeMrp - IFLA_BRIDGE_MRP
	IflaBridgeMrp IfLaBridge = 4

	// IflaBridgeCfm - IFLA_BRIDGE_CFM
	IflaBridgeCfm IfLaBridge = 5

	// IflaBridgeMst - IFLA_BRIDGE_MST
	IflaBridgeMst IfLaBridge = 6

	// IflaBridgeMax - IFLA_BRIDGE_MAX (maximum value)
	IflaBridgeMax IfLaBridge = IflaBridgeMst
)
