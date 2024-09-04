package core

// RdmaNlDevCmdEnum - Enumerated commands for RDMA device operations in Netlink.
//
// This enumeration defines various commands for interacting with RDMA devices via Netlink. These commands allow
// for getting, setting, and managing RDMA device states, and are important for RDMA resources like ports, queues,
// and memory regions.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
type RdmaNlDevCmdEnum uint8

const (
	// RdmaNldevCmdUnspec - RDMA_NLDEV_CMD_UNSPEC - Unspecified command (typically unused).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdUnspec RdmaNlDevCmdEnum = 0

	// RdmaNldevCmdGet - RDMA_NLDEV_CMD_GET - Get RDMA device attributes.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdGet RdmaNlDevCmdEnum = 1

	// RdmaNldevCmdSet - RDMA_NLDEV_CMD_SET - Set RDMA device attributes.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdSet RdmaNlDevCmdEnum = 2

	// RdmaNldevCmdNewlink - RDMA_NLDEV_CMD_NEWLINK - Create a new RDMA link.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdNewlink RdmaNlDevCmdEnum = 3

	// RdmaNldevCmdDellink - RDMA_NLDEV_CMD_DELLINK - Delete an RDMA link.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdDellink RdmaNlDevCmdEnum = 4

	// RdmaNldevCmdPortGet - RDMA_NLDEV_CMD_PORT_GET - Get information about RDMA ports (can dump).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdPortGet RdmaNlDevCmdEnum = 5

	// RdmaNldevCmdSysGet - RDMA_NLDEV_CMD_SYS_GET - Get RDMA system attributes.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdSysGet RdmaNlDevCmdEnum = 6

	// RdmaNldevCmdSysSet - RDMA_NLDEV_CMD_SYS_SET - Set RDMA system attributes.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdSysSet RdmaNlDevCmdEnum = 7

	// Placeholder for unused value
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	_ RdmaNlDevCmdEnum = 8

	// RdmaNldevCmdResGet - RDMA_NLDEV_CMD_RES_GET - Get RDMA resource information (can dump).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResGet RdmaNlDevCmdEnum = 9

	// RdmaNldevCmdResQpGet - RDMA_NLDEV_CMD_RES_QP_GET - Get RDMA Queue Pair (QP) resource information (can dump).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResQpGet RdmaNlDevCmdEnum = 10

	// RdmaNldevCmdResCmIdGet - RDMA_NLDEV_CMD_RES_CM_ID_GET - Get RDMA Communication Manager (CM) ID resources (can dump).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResCmIdGet RdmaNlDevCmdEnum = 11

	// RdmaNldevCmdResCqGet - RDMA_NLDEV_CMD_RES_CQ_GET - Get RDMA Completion Queue (CQ) resource information (can dump).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResCqGet RdmaNlDevCmdEnum = 12

	// RdmaNldevCmdResMrGet - RDMA_NLDEV_CMD_RES_MR_GET - Get RDMA Memory Region (MR) resources (can dump).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResMrGet RdmaNlDevCmdEnum = 13

	// RdmaNldevCmdResPdGet - RDMA_NLDEV_CMD_RES_PD_GET - Get RDMA Protection Domain (PD) resource information (can dump).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResPdGet RdmaNlDevCmdEnum = 14

	// RdmaNldevCmdGetChardev - RDMA_NLDEV_CMD_GET_CHARDEV - Get RDMA character device information.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdGetChardev RdmaNlDevCmdEnum = 15

	// RdmaNldevCmdStatSet - RDMA_NLDEV_CMD_STAT_SET - Set RDMA device statistics.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdStatSet RdmaNlDevCmdEnum = 16

	// RdmaNldevCmdStatGet - RDMA_NLDEV_CMD_STAT_GET - Get RDMA device statistics (can dump).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdStatGet RdmaNlDevCmdEnum = 17

	// RdmaNldevCmdStatDel - RDMA_NLDEV_CMD_STAT_DEL - Delete RDMA device statistics.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdStatDel RdmaNlDevCmdEnum = 18

	// RdmaNldevCmdResQpGetRaw - RDMA_NLDEV_CMD_RES_QP_GET_RAW - Get raw RDMA Queue Pair (QP) resource information.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResQpGetRaw RdmaNlDevCmdEnum = 19

	// RdmaNldevCmdResCqGetRaw - RDMA_NLDEV_CMD_RES_CQ_GET_RAW - Get raw RDMA Completion Queue (CQ) resource information.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResCqGetRaw RdmaNlDevCmdEnum = 20

	// RdmaNldevCmdResMrGetRaw - RDMA_NLDEV_CMD_RES_MR_GET_RAW - Get raw RDMA Memory Region (MR) resource information.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResMrGetRaw RdmaNlDevCmdEnum = 21

	// RdmaNldevCmdResCtxGet - RDMA_NLDEV_CMD_RES_CTX_GET - Get RDMA context resource information (can dump).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResCtxGet RdmaNlDevCmdEnum = 22

	// RdmaNldevCmdResSrqGet - RDMA_NLDEV_CMD_RES_SRQ_GET - Get RDMA Shared Receive Queue (SRQ) resource information (can dump).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResSrqGet RdmaNlDevCmdEnum = 23

	// RdmaNldevCmdStatGetStatus - RDMA_NLDEV_CMD_STAT_GET_STATUS - Get the status of RDMA device statistics.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdStatGetStatus RdmaNlDevCmdEnum = 24

	// RdmaNldevCmdResSrqGetRaw - RDMA_NLDEV_CMD_RES_SRQ_GET_RAW - Get raw RDMA Shared Receive Queue (SRQ) resource information.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdResSrqGetRaw RdmaNlDevCmdEnum = 25

	// RdmaNldevCmdNewdev - RDMA_NLDEV_CMD_NEWDEV - Create a new RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdNewdev RdmaNlDevCmdEnum = 26

	// RdmaNldevCmdDeldev - RDMA_NLDEV_CMD_DELDEV - Delete an RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevCmdDeldev RdmaNlDevCmdEnum = 27

	// RdmaNldevNumOps - RDMA_NLDEV_NUM_OPS - Total number of RDMA Netlink operations.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevNumOps RdmaNlDevCmdEnum = 28
)
