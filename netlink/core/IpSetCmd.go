package core

type IpSetCmd int

const (
	_ IpSetCmd = iota
	// IPSET_CMD_PROTOCOL - 1: Return protocol version
	IPSET_CMD_PROTOCOL
	// IPSET_CMD_CREATE - 2: Create a new (empty) set
	IPSET_CMD_CREATE
	// IPSET_CMD_DESTROY - 3: Destroy a (empty) set
	IPSET_CMD_DESTROY
	// IPSET_CMD_FLUSH - 4: Remove all elements from a set
	IPSET_CMD_FLUSH
	// IPSET_CMD_RENAME -  5: Rename a set
	IPSET_CMD_RENAME
	// IPSET_CMD_SWAP - 6: Swap two sets
	IPSET_CMD_SWAP
	// IPSET_CMD_LIST -  7: List sets
	IPSET_CMD_LIST
	// IPSET_CMD_SAVE - 8: Save sets
	IPSET_CMD_SAVE
	// IPSET_CMD_ADD  - 9: Add an element to a set
	IPSET_CMD_ADD
	// IPSET_CMD_DEL  - 10: Delete an element from a set
	IPSET_CMD_DEL
	// IPSET_CMD_TEST  - 11: Test an element in a set
	IPSET_CMD_TEST
	// IPSET_CMD_HEADER - 12: Get set header data only
	IPSET_CMD_HEADER
	// IPSET_CMD_TYPE  - 13: Get set type
	IPSET_CMD_TYPE
)
