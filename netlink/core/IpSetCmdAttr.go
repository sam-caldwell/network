package core

// IpSetCmdAttr - An RtAttrType command-level attributes
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
type IpSetCmdAttr int

const (
	_ IpSetCmdAttr = 0

	// IpsetAttrProtocol - IPSET_ATTR_PROTOCOL- 1: Protocol version
	IpsetAttrProtocol IpSetCmdAttr = 1

	// IpsetAttrSetname - IPSET_ATTR_SETNAME - 2: Name of the set
	IpsetAttrSetname IpSetCmdAttr = 2

	// IpsetAttrTypename - IPSET_ATTR_TYPENAME- 3: Typename
	IpsetAttrTypename IpSetCmdAttr = 3

	// IpsetAttrRevision - IPSET_ATTR_REVISION- 4: Settype revision
	IpsetAttrRevision IpSetCmdAttr = 4

	// IpsetAttrFamily - IPSET_ATTR_FAMILY - 5: Settype family
	IpsetAttrFamily IpSetCmdAttr = 5

	// IpsetAttrFlags - IPSET_ATTR_FLAGS - 6: Flags at command level
	IpsetAttrFlags IpSetCmdAttr = 6

	// IpsetAttrData - IPSET_ATTR_DATA - 7: Nested attributes
	IpsetAttrData IpSetCmdAttr = 7

	// IpsetAttrAdt - IPSET_ATTR_ADT - 8: Multiple data containers
	IpsetAttrAdt IpSetCmdAttr = 8

	// IpsetAttrLineno - IPSET_ATTR_LINENO - 9: Restore line number
	IpsetAttrLineno IpSetCmdAttr = 9

	// IpsetAttrProtocolMin - IPSET_ATTR_PROTOCOL_MIN - 10: Minimal supported version number
	IpsetAttrProtocolMin IpSetCmdAttr = 10

	// IpsetAttrSetname2 - IPSET_ATTR_SETNAME2 - Setname at rename/swap
	IpsetAttrSetname2 IpSetCmdAttr = IpsetAttrTypename

	// IpsetAttrRevisionMin - IPSET_ATTR_REVISION_MIN  type rev min
	IpsetAttrRevisionMin IpSetCmdAttr = IpsetAttrProtocolMin
)
