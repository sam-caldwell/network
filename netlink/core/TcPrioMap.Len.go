package core

// Len returns the size of the TcPriorityMap object in bytes.
//
// This function calculates the size of the TcPriorityMap object based on its structure
// to facilitate operations like serialization or validation.
//
// TcPriorityMap is used for priority queueing options in the Linux traffic control system
// to map priorities to traffic classes.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
func (msg *TcPriorityMap) Len() int {
	return SizeOfTcPriorityMap
}
