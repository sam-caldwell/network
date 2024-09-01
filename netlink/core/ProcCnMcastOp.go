package core

// ProcCnMcastOp - proc_cn_mcast_op - Userspace sends this enum to register with the kernel that it is
//
//	listening for events on the connector.
type ProcCnMcastOp int

const (

	// ProcCnMcastListen - PROC_CN_MCAST_LISTEN - enum ProcCnMcastOp (proc_cn_mcast_op)
	ProcCnMcastListen ProcCnMcastOp = 1

	// ProcCnMcastIgnore - PROC_CN_MCAST_IGNORE - enum ProcCnMcastOp (proc_cn_mcast_op)
	ProcCnMcastIgnore ProcCnMcastOp = 2
)
