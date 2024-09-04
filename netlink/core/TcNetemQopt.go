package core

// TcNetemQopt represents the network emulation queuing discipline options in Linux.
// It allows simulation of various network conditions, such as latency, packet loss, and duplication.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcNetemQopt struct {
	Latency   uint32 /* added delay in microseconds (us) */
	Limit     uint32 /* limit on the number of packets in the queue (fifo limit) */
	Loss      uint32 /* probability of random packet loss (0 = none, ~0 = 100%) */
	Gap       uint32 /* re-ordering gap (0 = no reordering) */
	Duplicate uint32 /* probability of packet duplication (0 = none, ~0 = 100%) */
	Jitter    uint32 /* random jitter in latency, in microseconds (us) */
}
