package core

// IpSetCmdAttrEnum - Enumeration for IP Set command-level attributes
//
// This enumeration defines the various attributes associated with IP Set commands in the Linux Netfilter framework.
// These attributes are used to provide additional data or options for specific IP Set commands.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
type IpSetCmdAttrEnum uint8

const (
	// IpsetAttrUnspec - IPSET_ATTR_UNSPEC - unspecified IP set attribute.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrUnspec IpSetCmdAttrEnum = 0

	// IpsetAttrProtocol - IPSET_ATTR_PROTOCOL - Protocol version
	// This attribute specifies the protocol version used by the IP Set infrastructure.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrProtocol IpSetCmdAttrEnum = 1

	// IpsetAttrSetname - IPSET_ATTR_SETNAME - Name of the set
	// This attribute specifies the name of an IP Set.
	// It is used when creating, modifying, or referencing a specific set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrSetname IpSetCmdAttrEnum = 2

	// IpsetAttrTypename - IPSET_ATTR_TYPENAME - Typename
	// This attribute specifies the type name of an IP Set.
	// The type name defines what kind of elements the set can contain (e.g., IP addresses, networks).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrTypename IpSetCmdAttrEnum = 3

	// IpsetAttrSetname2 - IPSET_ATTR_SETNAME2 - Setname at rename/swap
	// This attribute specifies the second set name during rename or swap operations.
	// It is used to reference the target set in these commands.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrSetname2 IpSetCmdAttrEnum = IpsetAttrTypename

	// IpsetAttrRevision - IPSET_ATTR_REVISION - Settype revision
	// This attribute specifies the revision of the IP Set type.
	// Revisions indicate different versions or enhancements of a particular set type.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrRevision IpSetCmdAttrEnum = 4

	// IpsetAttrFamily - IPSET_ATTR_FAMILY - Settype family
	// This attribute specifies the address family of the IP Set.
	// The family can be IPv4 or IPv6, indicating the type of addresses stored in the set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrFamily IpSetCmdAttrEnum = 5

	// IpsetAttrFlags - IPSET_ATTR_FLAGS - Flags at command level
	// This attribute specifies command-level flags for IP Set operations.
	// Flags control various options or behaviors of the IP Set commands.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrFlags IpSetCmdAttrEnum = 6

	// IpsetAttrData - IPSET_ATTR_DATA - Nested attributes
	// This attribute is a container for nested attributes.
	// It is used to encapsulate additional data required for specific commands.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrData IpSetCmdAttrEnum = 7

	// IpsetAttrAdt - IPSET_ATTR_ADT - Multiple data containers
	// This attribute specifies multiple data containers for IP Set operations.
	// It is typically used in commands that require handling multiple elements or sets simultaneously.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrAdt IpSetCmdAttrEnum = 8

	// IpsetAttrLineno - IPSET_ATTR_LINENO - Restore line number
	// This attribute specifies the line number for restoring IP Sets.
	// It is used to keep track of line numbers during restore operations.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrLineno IpSetCmdAttrEnum = 9

	// IpsetAttrProtocolMin - IPSET_ATTR_PROTOCOL_MIN - Minimal supported version number
	// This attribute specifies the minimal protocol version supported by the IP Set infrastructure.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrProtocolMin IpSetCmdAttrEnum = 10

	// IpsetAttrRevisionMin - IPSET_ATTR_REVISION_MIN - Type revision minimum
	// This attribute specifies the minimal revision supported for a specific IP Set type.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrRevisionMin IpSetCmdAttrEnum = IpsetAttrProtocolMin

	// IpsetAttrIndex - IPSET_ATTR_INDEX - Kernel index of set
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrIndex IpSetCmdAttrEnum = 11

	// IpsetAttrCmdMax - IPSET_ATTR_CMD_MAX - Maximum value of set
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetAttrCmdMax = IpsetAttrIndex
)
