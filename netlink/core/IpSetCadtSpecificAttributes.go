package core

// IpSetCadtSpecificAttributes - An RtAttrType CADT specific attributes
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
type IpSetCadtSpecificAttributes int

const (

	// IPSET_ATTR_UNSPEC - unspecified IP set attribute.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IPSET_ATTR_UNSPEC = iota

	/*
		General attributes
	*/

	// IpsetAttrIp - IPSET_ATTR_IP - represents a single IP address attribute.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrIp IpSetCadtSpecificAttributes = 1

	// IpsetAttrIpFrom - IPSET_ATTR_IP_FROM - Alias for IpsetAttrIp representing starting ip address in range.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrIpFrom IpSetCadtSpecificAttributes = IpsetAttrIp

	// IpsetAttrIpTo - IPSET_ATTR_IP_TO - Alias for IpsetAttrIp representing ending ip address in range.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrIpTo IpSetCadtSpecificAttributes = 2

	// IpsetAttrCidr - IPSET_ATTR_CIDR - Represents a CIDR (Classless Inter-Domain Routing) block attribute.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrCidr IpSetCadtSpecificAttributes = 3

	// IpsetAttrPort - IPSET_ATTR_PORT - Represents a single port attribute.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrPort IpSetCadtSpecificAttributes = 4

	// IpsetAttrPortFrom - IPSET_ATTR_PORT_FROM - Represents a single port range start attribute.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrPortFrom IpSetCadtSpecificAttributes = 4

	// IpsetAttrPortTo - IPSET_ATTR_PORT_TO - Represents a single port range ending attribute.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrPortTo IpSetCadtSpecificAttributes = 5

	// IpsetAttrTimeout - IPSET_ATTR_TIMEOUT - Represents the timeout attribute for an IP set entry.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrTimeout IpSetCadtSpecificAttributes = 6

	// IpsetAttrProto - IPSET_ATTR_PROTO - Represents the protocol (e.g., TCP, UDP) associated with an IP set entry.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrProto IpSetCadtSpecificAttributes = 7

	// IpsetAttrCadtFlags - IPSET_ATTR_CADT_FLAGS -
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrCadtFlags IpSetCadtSpecificAttributes = 8

	// IpsetAttrCadtLineno - IPSET_ATTR_CADT_LINENO - Alias of IpSetCmdAttr.IpsetAttrLineno
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrCadtLineno IpSetCadtSpecificAttributes = IpSetCadtSpecificAttributes(IpsetAttrLineno)

	// IpsetAttrMark - IPSET_ATTR_MARK - Represents a mark attribute, often used for marking packets.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrMark IpSetCadtSpecificAttributes = 10

	// IpsetAttrMarkmask - IPSET_ATTR_MARKMASK - Represents a mask applied to the mark attribute.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrMarkmask IpSetCadtSpecificAttributes = 11

	// IpsetAttrCadtMax - IPSET_ATTR_CADT_MAX - Reserve empty slots
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrCadtMax IpSetCadtSpecificAttributes = 16

	/*
	  Create-only specific attributes
	*/

	// IpsetAttrInitval - IPSET_ATTR_INITVAL - Initial value for a hash set, was previously unused
	// (was IPSET_ATTR_GC).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrInitval IpSetCadtSpecificAttributes = 17

	// IpsetAttrHashsize -IPSET_ATTR_HASHSIZE - Represents the size of the hash table in a hash set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrHashsize IpSetCadtSpecificAttributes = 18

	// IpsetAttrMaxelem -IPSET_ATTR_MAXELEM - Represents the maximum number of elements in a set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrMaxelem IpSetCadtSpecificAttributes = 19

	// IpsetAttrNetmask -IPSET_ATTR_NETMASK - Represents the netmask attribute, used in network set types.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrNetmask IpSetCadtSpecificAttributes = 20

	// IpsetAttrBucketsize -IPSET_ATTR_BUCKETSIZE - Represents the bucket size in a hash set, was previously
	// unused (was IPSET_ATTR_PROBES).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrBucketsize IpSetCadtSpecificAttributes = 21

	// IpsetAttrResize -IPSET_ATTR_RESIZE - Represents the resize attribute, controlling automatic resizing
	// of the set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrResize IpSetCadtSpecificAttributes = 22

	// IpsetAttrSize -IPSET_ATTR_SIZE - Represents the initial size or capacity of the set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrSize IpSetCadtSpecificAttributes = 23

	/*
		Kernel-only attributes
	*/

	// IpsetAttrElements - IpsetAttrElements - Kernel-only attribute. Represents the current number of elements
	// in the set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrElements IpSetCadtSpecificAttributes = 24

	// IpsetAttrReferences - IpsetAttrReferences - Kernel-only attribute. Represents the number of references
	// to the set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrReferences IpSetCadtSpecificAttributes = 25

	// IpsetAttrMemsize - IPSET_ATTR_MEMSIZE - Kernel-only attribute. Represents the memory size used by the set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrMemsize IpSetCadtSpecificAttributes = 26

	// SetAttrCreateMax - SET_ATTR_CREATE_MAX - Kernel-only attribute. Maximum attribute index used during set
	// creation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	SetAttrCreateMax IpSetCadtSpecificAttributes = 26
)
