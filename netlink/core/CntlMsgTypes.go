package core

// CntlMsgTypes - All the following constants are coming from:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CntlMsgTypes uint8

const (

	// IpctnlMsgCtNew - IPCTNL_MSG_CT_NEW
	IpctnlMsgCtNew CntlMsgTypes = 0

	// IpctnlMsgCtGet - IPCTNL_MSG_CT_GET
	IpctnlMsgCtGet CntlMsgTypes = 1

	// IpctnlMsgCtDelete - IPCTNL_MSG_CT_DELETE
	IpctnlMsgCtDelete CntlMsgTypes = 2

	// IpctnlMsgCtGetCtrzero - IPCTNL_MSG_CT_GET_CTRZERO
	IpctnlMsgCtGetCtrzero CntlMsgTypes = 3

	// IpctnlMsgCtGetStatsCpu - IPCTNL_MSG_CT_GET_STATS_CPU
	IpctnlMsgCtGetStatsCpu CntlMsgTypes = 4

	// IpctnlMsgCtGetStats - IPCTNL_MSG_CT_GET_STATS
	IpctnlMsgCtGetStats CntlMsgTypes = 5

	// IpctnlMsgCtGetDying - IPCTNL_MSG_CT_GET_DYING
	IpctnlMsgCtGetDying CntlMsgTypes = 6

	// IpctnlMsgCtGetUnconfirmed - IPCTNL_MSG_CT_GET_UNCONFIRMED
	IpctnlMsgCtGetUnconfirmed CntlMsgTypes = 7

	// IpctnlMsgMax - IPCTNL_MSG_MAX
	IpctnlMsgMax CntlMsgTypes = IpctnlMsgCtGetUnconfirmed
)
