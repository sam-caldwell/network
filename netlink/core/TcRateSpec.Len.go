package core

// Len - Returns the size of the TcRateSpec structure in bytes.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
func (msg *TcRateSpec) Len() int {
	return SizeOfTcRateSpec
}
