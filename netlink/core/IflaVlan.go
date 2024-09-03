package core

// IflaVlan - Enumeration for VLAN attributes in Linux kernel -
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaVlan uint8

const (

	// IflaVlanUnspec - IFLA_VLAN_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	// It represents an undefined or unused attribute.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVlanUnspec IflaVlan = 0

	// IflaVlanId - IFLA_VLAN_ID -
	// This attribute represents the VLAN ID, a 12-bit identifier that specifies
	// the VLAN to which the frame belongs. VLAN IDs are typically in the range of 1 to 4094.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVlanId IflaVlan = 1

	// IflaVlanFlags - IFLA_VLAN_FLAGS -
	// This attribute represents various flags associated with the VLAN interface.
	// These flags can control specific behaviors or properties of the VLAN.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVlanFlags IflaVlan = 2

	// IflaVlanEgressQos - IFLA_VLAN_EGRESS_QOS -
	// This attribute represents the Quality of Service (QoS) settings for outgoing (egress) traffic
	// on the VLAN. QoS settings can prioritize certain types of traffic to ensure better performance.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVlanEgressQos IflaVlan = 3

	// IflaVlanIngressQos - IFLA_VLAN_INGRESS_QOS -
	// This attribute represents the Quality of Service (QoS) settings for incoming (ingress) traffic
	// on the VLAN. Similar to egress QoS, this helps manage and prioritize traffic entering the VLAN.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVlanIngressQos IflaVlan = 4

	// IflaVlanProtocol - IFLA_VLAN_PROTOCOL -
	// This attribute represents the protocol used for the VLAN. Typically, this is either
	// 802.1Q (the standard protocol for VLAN tagging) or another protocol as specified by the implementation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVlanProtocol IflaVlan = 5

	// IflaVlanMax - IFLA_VLAN_MAX -
	// This constant represents the maximum value for VLAN attributes.
	// It is used as a boundary marker to validate or limit the range of VLAN attributes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVlanMax = IflaVlanProtocol
)
