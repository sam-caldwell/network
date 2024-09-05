package core

// Len - Returns the size of the TcSfqQopt structure in bytes.
// This is useful for allocating memory or determining how much space is needed for serialization.
// The size is derived from the SizeofTcSfqQopt constant, which corresponds to the size of the TcSfqQopt structure.
//
// See more about SFQ (Stochastic Fairness Queueing) at:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
func (msg *TcSfqQopt) Len() int {
	return SizeofTcSfqQopt
}
