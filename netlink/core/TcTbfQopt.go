package core

// TcTbfQopt - Configuration options for the Token Bucket Filter (TBF) queuing discipline.
//
// This struct represents the options for configuring a TBF in the Linux kernel's
// traffic control system. TBF is used to limit the rate of traffic on a network interface
// by using a token bucket algorithm.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcTbfQopt struct {
	// Rate - Specifies the average rate at which tokens are replenished.
	// This defines the sustained bandwidth for the TBF.
	Rate TcRateSpec

	// Peakrate - Specifies the maximum rate at which tokens can be consumed.
	// This defines the burst bandwidth for the TBF.
	Peakrate TcRateSpec

	// Limit - Specifies the maximum amount of bytes that can be queued.
	Limit uint32

	// Buffer - Specifies the buffer size in bytes, used to store tokens.
	Buffer uint32

	// Mtu - Specifies the Maximum Transmission Unit (MTU) for the interface.
	Mtu uint32
}
