package core

// IpSetAdtSpecificAttributes - An RtAttrType ADT specific attributes
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
type IpSetAdtSpecificAttributes int

const (
	// IpsetAttrEther - IPSET_ATTR_ETHER -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrEther = IpSetAdtSpecificAttributes(IpsetAttrCadtMax) + 1

	// IpsetAttrName - IPSET_ATTR_NAME -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrName IpSetAdtSpecificAttributes = IpsetAttrEther + 1

	// IpsetAttrNameref - IPSET_ATTR_NAMEREF -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrNameref IpSetAdtSpecificAttributes = IpsetAttrEther + 2

	// IpsetAttrIp2 - IPSET_ATTR_IP2 -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrIp2 IpSetAdtSpecificAttributes = IpsetAttrEther + 3

	// IpsetAttrCidr2 - IPSET_ATTR_CIDR2 -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrCidr2 IpSetAdtSpecificAttributes = IpsetAttrEther + 4

	// IpsetAttrIp2To - IPSET_ATTR_IP2_TO -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrIp2To IpSetAdtSpecificAttributes = IpsetAttrEther + 5

	// IpsetAttrIface - IPSET_ATTR_IFACE -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrIface IpSetAdtSpecificAttributes = IpsetAttrEther + 6

	// IpsetAttrBytes - IPSET_ATTR_BYTES -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrBytes IpSetAdtSpecificAttributes = IpsetAttrEther + 7

	// IpsetAttrPackets - IPSET_ATTR_PACKETS -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrPackets IpSetAdtSpecificAttributes = IpsetAttrEther + 8

	// IpsetAttrComment - IPSET_ATTR_COMMENT -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrComment IpSetAdtSpecificAttributes = IpsetAttrEther + 9

	// IpsetAttrSkbmark - IPSET_ATTR_SKBMARK -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrSkbmark IpSetAdtSpecificAttributes = IpsetAttrEther + 10

	// IpsetAttrSkbprio - IPSET_ATTR_SKBPRIO -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrSkbprio IpSetAdtSpecificAttributes = IpsetAttrEther + 11

	// IpsetAttrSkbqueue - IPSET_ATTR_SKBQUEUE -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrSkbqueue IpSetAdtSpecificAttributes = IpsetAttrEther + 12

	// IpsetAttrPad - IPSET_ATTR_PAD -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrPad IpSetAdtSpecificAttributes = IpsetAttrEther + 13
)

const (
	// IpsetAttrIpaddrIpv4 - IPSET_ATTR_IPADDR_IPV4 -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrIpaddrIpv4 IpSetAdtSpecificAttributes = IpSetAdtSpecificAttributes(IpsetAttrUnspec + 1)

	// IpsetAttrIpaddrIpv6 - IPSET_ATTR_IPADDR_IPV6 -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrIpaddrIpv6 IpSetAdtSpecificAttributes = IpSetAdtSpecificAttributes(IpsetAttrUnspec + 2)
)
