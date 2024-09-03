package core

// IpSetCmd Message type / commands
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
type IpSetCmd uint8

const (
	_ IpSetCmd = iota

	// IpsetCmdProtocol - IPSET_CMD_PROTOCOL - 1: Return protocol version
	IpsetCmdProtocol
	// IpsetCmdCreate - IPSET_CMD_CREATE - 2: Create a new (empty) set
	IpsetCmdCreate
	// IpsetCmdDestroy - IPSET_CMD_DESTROY - 3: Destroy a (empty) set
	IpsetCmdDestroy
	// IpsetCmdFlush - IPSET_CMD_FLUSH - 4: Remove all elements from a set
	IpsetCmdFlush
	// IpsetCmdRename - IPSET_CMD_RENAME - 5: Rename a set
	IpsetCmdRename
	// IpsetCmdSwap - IPSET_CMD_SWAP- 6: Swap two sets
	IpsetCmdSwap
	// IpsetCmdList - IPSET_CMD_LIST - 7: List sets
	IpsetCmdList
	// IpsetCmdSave - IPSET_CMD_SAVE - 8: Save sets
	IpsetCmdSave
	// IpsetCmdAdd  - IPSET_CMD_ADD - 9: Add an element to a set
	IpsetCmdAdd
	// IpsetCmdDel  - IPSET_CMD_DEL - 10: Delete an element from a set
	IpsetCmdDel
	// IpsetCmdTest  - IPSET_CMD_TEST - 11: Test an element in a set
	IpsetCmdTest
	// IpsetCmdHeader - IPSET_CMD_HEADER - 12: Get set header data only
	IpsetCmdHeader
	// IpsetCmdType  - IPSET_CMD_TYPE - 13: Get set type
	IpsetCmdType
)
