package core

// CntlMsgTypes - cntl_msg_types - All the following constants are coming from:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CntlMsgTypes uint8

const (

	// IpCtnlMsgCtNew - IPCTNL_MSG_CT_NEW
	IpCtnlMsgCtNew CntlMsgTypes = 0

	// IpCtnlMsgCtGet - IPCTNL_MSG_CT_GET
	IpCtnlMsgCtGet CntlMsgTypes = 1

	// IpCtnlMsgCtDelete - IPCTNL_MSG_CT_DELETE
	IpCtnlMsgCtDelete CntlMsgTypes = 2

	// IpCtnlMsgCtGetCtrZero - IPCTNL_MSG_CT_GET_CTRZERO
	IpCtnlMsgCtGetCtrZero CntlMsgTypes = 3

	// IpCtnlMsgCtGetStatsCpu - IPCTNL_MSG_CT_GET_STATS_CPU
	IpCtnlMsgCtGetStatsCpu CntlMsgTypes = 4

	// IpCtnlMsgCtGetStats - IPCTNL_MSG_CT_GET_STATS
	IpCtnlMsgCtGetStats CntlMsgTypes = 5

	// IpCtnlMsgCtGetDying - IPCTNL_MSG_CT_GET_DYING
	IpCtnlMsgCtGetDying CntlMsgTypes = 6

	// IpCtnlMsgCtGetUnconfirmed - IPCTNL_MSG_CT_GET_UNCONFIRMED
	IpCtnlMsgCtGetUnconfirmed CntlMsgTypes = 7

	// IpCtnlMsgMax - IPCTNL_MSG_MAX
	IpCtnlMsgMax CntlMsgTypes = iota - 1
)
