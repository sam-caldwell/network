package core

// IpSetAdtSpecificAttributes - An RtAttrType ADT specific attributes
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
type IpSetAdtSpecificAttributes int

const (
	// IpsetAttrEther - IPSET_ATTR_ETHER -
	IpsetAttrEther = IpSetAdtSpecificAttributes(IpsetAttrCadtMax) + 1

	// IpsetAttrName - IPSET_ATTR_NAME -
	IpsetAttrName IpSetAdtSpecificAttributes = IpsetAttrEther + 1

	// IpsetAttrNameref - IPSET_ATTR_NAMEREF -
	IpsetAttrNameref IpSetAdtSpecificAttributes = IpsetAttrEther + 2

	// IpsetAttrIp2 - IPSET_ATTR_IP2 -
	IpsetAttrIp2 IpSetAdtSpecificAttributes = IpsetAttrEther + 3

	// IpsetAttrCidr2 - IPSET_ATTR_CIDR2 -
	IpsetAttrCidr2 IpSetAdtSpecificAttributes = IpsetAttrEther + 4

	// IpsetAttrIp2To - IPSET_ATTR_IP2_TO -
	IpsetAttrIp2To IpSetAdtSpecificAttributes = IpsetAttrEther + 5

	// IpsetAttrIface - IPSET_ATTR_IFACE -
	IpsetAttrIface IpSetAdtSpecificAttributes = IpsetAttrEther + 6

	// IpsetAttrBytes - IPSET_ATTR_BYTES -
	IpsetAttrBytes IpSetAdtSpecificAttributes = IpsetAttrEther + 7

	// IpsetAttrPackets - IPSET_ATTR_PACKETS -
	IpsetAttrPackets IpSetAdtSpecificAttributes = IpsetAttrEther + 8

	// IpsetAttrComment - IPSET_ATTR_COMMENT -
	IpsetAttrComment IpSetAdtSpecificAttributes = IpsetAttrEther + 9

	// IpsetAttrSkbmark - IPSET_ATTR_SKBMARK -
	IpsetAttrSkbmark IpSetAdtSpecificAttributes = IpsetAttrEther + 10

	// IpsetAttrSkbprio - IPSET_ATTR_SKBPRIO -
	IpsetAttrSkbprio IpSetAdtSpecificAttributes = IpsetAttrEther + 11

	// IpsetAttrSkbqueue - IPSET_ATTR_SKBQUEUE -
	IpsetAttrSkbqueue IpSetAdtSpecificAttributes = IpsetAttrEther + 12

	// IpsetAttrPad - IPSET_ATTR_PAD -
	IpsetAttrPad IpSetAdtSpecificAttributes = IpsetAttrEther + 13
)

const (
	// IpsetAttrIpaddrIpv4 - IPSET_ATTR_IPADDR_IPV4 -
	IpsetAttrIpaddrIpv4 IpSetAdtSpecificAttributes = IpSetAdtSpecificAttributes(IpsetAttrUnspec + 1)

	// IpsetAttrIpaddrIpv6 - IPSET_ATTR_IPADDR_IPV6 -
	IpsetAttrIpaddrIpv6 IpSetAdtSpecificAttributes = IpSetAdtSpecificAttributes(IpsetAttrUnspec + 2)
)
