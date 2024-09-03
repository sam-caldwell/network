package core

// IfLaBridgeEnum - Bridge management nested attributes
//
// Bridge management nested attributes
//
//	[IFLA_AF_SPEC] = {
//	   [IflaBridgeFlags]
//	   [IFLA_BRIDGE_MODE]
//	   [IFLA_BRIDGE_VLAN_INFO]
//	}
type IfLaBridgeEnum int

const (
	// IflaBridgeFlags - IFLA_BRIDGE_FLAGS
	IflaBridgeFlags IfLaBridgeEnum = 0

	// IflaBridgeMode - IFLA_BRIDGE_MODE
	IflaBridgeMode IfLaBridgeEnum = 1

	// IflaBridgeVlanInfo - IFLA_BRIDGE_VLAN_INFO
	IflaBridgeVlanInfo IfLaBridgeEnum = 2

	// IflaBridgeVlanTunnelInfo - IFLA_BRIDGE_VLAN_TUNNEL_INFO
	IflaBridgeVlanTunnelInfo IfLaBridgeEnum = 3

	// IflaBridgeMrp - IFLA_BRIDGE_MRP
	IflaBridgeMrp IfLaBridgeEnum = 4

	// IflaBridgeCfm - IFLA_BRIDGE_CFM
	IflaBridgeCfm IfLaBridgeEnum = 5

	// IflaBridgeMst - IFLA_BRIDGE_MST
	IflaBridgeMst IfLaBridgeEnum = 6

	// IflaBridgeMax - IFLA_BRIDGE_MAX (maximum value)
	IflaBridgeMax IfLaBridgeEnum = IflaBridgeMst
)
