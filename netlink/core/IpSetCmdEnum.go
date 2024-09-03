package core

// IpSetCmdEnum - Enumeration for IP Set commands
//
// This enumeration defines various commands related to IP sets in the Linux Netfilter framework.
// IP sets are used to store multiple IP addresses, networks, or port numbers and are typically used in firewall rules.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
type IpSetCmdEnum uint8

const (
	// IpsetCmdProtocol - IPSET_CMD_PROTOCOL - 1: Return protocol version
	// This command returns the protocol version used by the IP set infrastructure.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdProtocol IpSetCmdEnum = 1

	// IpsetCmdCreate - IPSET_CMD_CREATE - 2: Create a new (empty) set
	// This command creates a new, empty IP set. The set can then be populated with IP addresses, networks, or
	// port numbers.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdCreate IpSetCmdEnum = 2

	// IpsetCmdDestroy - IPSET_CMD_DESTROY - 3: Destroy a (empty) set
	// This command destroys an IP set. The set must be empty before it can be destroyed.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdDestroy IpSetCmdEnum = 3

	// IpsetCmdFlush - IPSET_CMD_FLUSH - 4: Remove all elements from a set
	// This command removes all elements from an IP set, effectively emptying it without deleting the set itself.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdFlush IpSetCmdEnum = 4

	// IpsetCmdRename - IPSET_CMD_RENAME - 5: Rename a set
	// This command renames an existing IP set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdRename IpSetCmdEnum = 5

	// IpsetCmdSwap - IPSET_CMD_SWAP - 6: Swap two sets
	// This command swaps the content and metadata of two IP sets. It is useful for atomic updates of IP sets.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdSwap IpSetCmdEnum = 6

	// IpsetCmdList - IPSET_CMD_LIST - 7: List sets
	// This command lists all existing IP sets, along with their elements and metadata.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdList IpSetCmdEnum = 7

	// IpsetCmdSave - IPSET_CMD_SAVE - 8: Save sets
	// This command saves the state of all IP sets, including their elements and metadata, for later restoration.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdSave IpSetCmdEnum = 8

	// IpsetCmdAdd - IPSET_CMD_ADD - 9: Add an element to a set
	// This command adds an element, such as an IP address or network, to an existing IP set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdAdd IpSetCmdEnum = 9

	// IpsetCmdDel - IPSET_CMD_DEL - 10: Delete an element from a set
	// This command deletes an element, such as an IP address or network, from an existing IP set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdDel IpSetCmdEnum = 10

	// IpsetCmdTest - IPSET_CMD_TEST - 11: Test an element in a set
	// This command tests whether a given element, such as an IP address or network, is a member of an IP set.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdTest IpSetCmdEnum = 11

	// IpsetCmdHeader - IPSET_CMD_HEADER - 12: Get set header data only
	// This command retrieves only the header data of an IP set, such as its name and type, without listing its
	// elements.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdHeader IpSetCmdEnum = 12

	// IpsetCmdType - IPSET_CMD_TYPE - 13: Get set type
	// This command retrieves the type of IP set, which defines what kind of elements it can contain (e.g.,
	// IP addresses, networks).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetCmdType IpSetCmdEnum = 13
)
