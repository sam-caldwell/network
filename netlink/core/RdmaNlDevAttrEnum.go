package core

// RdmaNlDevAttrEnum - Enumerated attributes for RDMA devices in Netlink.
//
// This enum defines various attributes associated with RDMA devices, such as device index,
// port information, firmware version, and network namespace details.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
type RdmaNlDevAttrEnum uint8

const (
	// RdmaNldevAttrDevIndex - RDMA_NLDEV_ATTR_DEV_INDEX - The index of the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrDevIndex RdmaNlDevAttrEnum = 1

	// RdmaNldevAttrDevName - RDMA_NLDEV_ATTR_DEV_NAME - The name of the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrDevName RdmaNlDevAttrEnum = 2

	// RdmaNldevAttrPortIndex - RDMA_NLDEV_ATTR_PORT_INDEX - The index of the RDMA port.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrPortIndex RdmaNlDevAttrEnum = 3

	// RdmaNldevAttrCapFlags - RDMA_NLDEV_ATTR_CAP_FLAGS - Capability flags for the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrCapFlags RdmaNlDevAttrEnum = 4

	// RdmaNldevAttrFwVersion - RDMA_NLDEV_ATTR_FW_VERSION - Firmware version of the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrFwVersion RdmaNlDevAttrEnum = 5

	// RdmaNldevAttrNodeGuid - RDMA_NLDEV_ATTR_NODE_GUID - Node GUID of the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrNodeGuid RdmaNlDevAttrEnum = 6

	// RdmaNldevAttrSysImageGuid - RDMA_NLDEV_ATTR_SYS_IMAGE_GUID - System image GUID for the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrSysImageGuid RdmaNlDevAttrEnum = 7

	// RdmaNldevAttrSubnetPrefix - RDMA_NLDEV_ATTR_SUBNET_PREFIX - Subnet prefix associated with the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrSubnetPrefix RdmaNlDevAttrEnum = 8

	// RdmaNldevAttrLid - RDMA_NLDEV_ATTR_LID - Local Identifier (LID) for the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrLid RdmaNlDevAttrEnum = 9

	// RdmaNldevAttrSmLid - RDMA_NLDEV_ATTR_SM_LID - Subnet Manager Local Identifier (SM LID) for the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrSmLid RdmaNlDevAttrEnum = 10

	// RdmaNldevAttrLmc - RDMA_NLDEV_ATTR_LMC - Link Multicast Count (LMC) for the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrLmc RdmaNlDevAttrEnum = 11

	// RdmaNldevAttrPortState - RDMA_NLDEV_ATTR_PORT_STATE - State of the RDMA port.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrPortState RdmaNlDevAttrEnum = 12

	// RdmaNldevAttrPortPhysState - RDMA_NLDEV_ATTR_PORT_PHYS_STATE - Physical state of the RDMA port.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrPortPhysState RdmaNlDevAttrEnum = 13

	// RdmaNldevAttrDevNodeType - RDMA_NLDEV_ATTR_DEV_NODE_TYPE - Type of node associated with the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrDevNodeType RdmaNlDevAttrEnum = 14

	// RdmaNldevAttrNdevName - RDMA_NLDEV_ATTR_NDEV_NAME - Name of the network device associated with RDMA.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrNdevName RdmaNlDevAttrEnum = 51

	// RdmaNldevAttrLinkType - RDMA_NLDEV_ATTR_LINK_TYPE - Type of link used by the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevAttrLinkType RdmaNlDevAttrEnum = 65

	// RdmaNldevSysAttrNetnsMode - RDMA_NLDEV_SYS_ATTR_NETNS_MODE - Network namespace mode for RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevSysAttrNetnsMode RdmaNlDevAttrEnum = 66

	// RdmaNldevNetNsFd - RDMA_NLDEV_NET_NS_FD - File descriptor for the network namespace.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNldevNetNsFd RdmaNlDevAttrEnum = 68
)
