package core

// RdmaNlEnum - Rdma Netlink Enumerated type
//
// This enumeration defines various types used in the RDMA (Remote Direct Memory Access) Netlink API.
// RDMA Netlink is used for communication between the kernel and user space for RDMA operations.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
type RdmaNlEnum uint8

const (
	// RdmaNlIwcm - RDMA_NL_IWCM - Represents the RDMA iWARP Connection Manager.
	// iWARP is a protocol for RDMA over IP networks. This identifier is used in
	// managing RDMA connections for iWARP devices.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNlIwcm RdmaNlEnum = 2

	// RdmaNlRsvd - RDMA_NL_RSVD - Reserved for future use.
	// This identifier is reserved for potential future RDMA Netlink clients or features.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNlRsvd RdmaNlEnum = 3

	// RdmaNlLs - RDMA_NL_LS - Represents RDMA Local Services.
	// This is used for managing and interacting with local RDMA services such as
	// local RDMA communication services.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNlLs RdmaNlEnum = 4

	// RdmaNlNldev - RDMA_NL_NLDEV - Represents RDMA device management.
	// This identifier is used for managing RDMA device interfaces and configurations via Netlink.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNlNldev RdmaNlEnum = 5

	// RdmaNlNumClients - RDMA_NL_NUM_CLIENTS - Represents the number of RDMA device clients.
	// This is used in RDMA device interface management, indicating the number of clients
	// connected to or using the RDMA device.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h
	RdmaNlNumClients RdmaNlEnum = 6
)
