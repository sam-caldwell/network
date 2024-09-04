package core

// TcaPriority - Enumeration for Traffic Control Action (TCA) priorities.
//
// These constants represent various priority settings for TCA in the Linux kernel.
//
// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcaPriority uint8

const (
	// TcaPrioUnspec - TCA_PRIO_UNSPEC - Unspecified priority setting.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaPrioUnspec TcaPriority = iota

	// TcaPrioMq - TCA_PRIO_MQ - Multi-queue priority setting.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaPrioMq

	// TcaPrioMax - TCA_PRIO_MAX - Maximum priority setting (equivalent to multi-queue).
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaPrioMax = TcaPrioMq
)
