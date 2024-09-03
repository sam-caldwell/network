package core

// IflaVfEnum - Enumeration for Virtual Function (VF) attributes
//
// This enumeration defines various attributes related to Virtual Functions (VFs) in network interfaces.
// These attributes are used to query or configure settings for VFs, which are used in SR-IOV (Single Root I/O
// Virtualization) setups.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaVfEnum uint8

const (
	// IflaVfUnspec - IFLA_VF_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	// It represents an undefined or unused attribute within the context of Virtual Functions.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfUnspec IflaVfEnum = 0

	// IflaVfMac - IFLA_VF_MAC -
	// This attribute specifies the MAC address of the Virtual Function (VF).
	// It allows the configuration of a unique hardware address for each VF.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfMac IflaVfEnum = 1

	// IflaVfVlan - IFLA_VF_VLAN -
	// This attribute configures the VLAN ID and Quality of Service (QoS) settings for the VF.
	// It is used to assign a VF to a specific VLAN and control traffic prioritization.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfVlan IflaVfEnum = 2

	// IflaVfTxRate - IFLA_VF_TX_RATE -
	// This attribute sets the maximum transmit (TX) bandwidth allocation for the VF.
	// It limits the amount of bandwidth that the VF can use for outgoing traffic.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfTxRate IflaVfEnum = 3

	// IflaVfSpoofchk - IFLA_VF_SPOOFCHK -
	// This attribute enables or disables spoof checking on the VF.
	// Spoof checking prevents the VF from sending traffic with a MAC or IP address that it is not assigned.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfSpoofchk IflaVfEnum = 4

	// IflaVfLinkState - IFLA_VF_LINK_STATE -
	// This attribute controls the link state of the VF, with options to enable, disable, or set it to auto.
	// The link state determines whether the VF is considered to be up, down, or follows the state of the physical
	// link.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfLinkState IflaVfEnum = 5

	// IflaVfRate - IFLA_VF_RATE -
	// This attribute configures both the minimum and maximum TX bandwidth allocation for the VF.
	// It allows for finer control over the bandwidth usage of the VF.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfRate IflaVfEnum = 6

	// IflaVfRssQueryEn - IFLA_VF_RSS_QUERY_EN -
	// This attribute enables or disables the ability to query the RSS (Receive Side Scaling) redirection table and
	// hash key.
	//
	// RSS is used to distribute network traffic across multiple CPUs to improve performance.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfRssQueryEn IflaVfEnum = 7

	// IflaVfStats - IFLA_VF_STATS -
	// This attribute provides access to network device statistics specific to the VF.
	// It includes metrics such as packet counts, error counts, and bandwidth usage.
	IflaVfStats IflaVfEnum = 8

	// IflaVfTrust - IFLA_VF_TRUST -
	// This attribute indicates whether the VF is trusted.
	// A trusted VF may have additional capabilities or permissions, such as the ability to change its MAC address.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfTrust IflaVfEnum = 9

	// IflaVfIbNodeGuid - IFLA_VF_IB_NODE_GUID -
	// This attribute sets the Infiniband node GUID (Globally Unique Identifier) for the VF.
	// It is used in environments where the VF operates over Infiniband networks.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfIbNodeGuid IflaVfEnum = 10

	// IflaVfIbPortGuid - IFLA_VF_IB_PORT_GUID -
	// This attribute sets the Infiniband port GUID for the VF.
	// It is used to uniquely identify the VF's port in an Infiniband network.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfIbPortGuid IflaVfEnum = 11

	// IflaVfVlanList - IFLA_VF_VLAN_LIST -
	// This attribute provides a nested list of VLANs assigned to the VF, allowing for the configuration of multiple
	// VLANs. It can also be used for QinQ (802.1ad) configurations, where a VF is associated with a stacked VLAN
	// setup.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfVlanList IflaVfEnum = 12

	// IflaVfBroadcast - IFLA_VF_BROADCAST -
	// This attribute controls broadcast traffic settings for the VF.
	// It is used to manage how broadcast traffic is handled or transmitted by the VF.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfBroadcast IflaVfEnum = 13

	// IflaVfMax - IFLA_VF_MAX -
	// This constant represents the maximum valid value for VF attributes.
	// It is used as a boundary marker to validate or limit the range of VF attributes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfMax = iota - 1
)
