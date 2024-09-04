package core

// https://github.com/torvalds/linux/blob/master/include/uapi/rdma/rdma_netlink.h

const (
	// RdmaNlGetClientShift is a bit shift constant used in the RDMA (Remote Direct Memory Access) Netlink API
	// to identify client types in RDMA communication. Netlink is a protocol that facilitates communication between
	// the kernel and user space, and the RDMA Netlink subsystem handles RDMA-related messages.
	//
	// The purpose of RdmaNlGetClientShift is to create space in the message type field to uniquely identify
	// RDMA clients (e.g., InfiniBand, iWARP) by shifting bits to avoid overlap with other message types.
	RdmaNlGetClientShift = 10
)
