package core

// TcMirred - Represents the tc_mirred structure used in Linux Traffic Control for mirroring and redirection actions.
// The tc_mirred action can mirror or redirect packets to another interface.
// It is typically used to monitor network traffic or modify packet flow.
//
// Fields:
// - `TcGen`: The general traffic control structure containing common parameters (e.g., action, refcnt).
// - `Eaction`: Specifies the action to be taken (e.g., ingress or egress mirroring or redirection).
// - `Ifindex`: The interface index where packets should be mirrored or redirected.
//
// See:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_mirred.h
// - https://man7.org/linux/man-pages/man8/tc-mirred.8.html
type TcMirred struct {
	// TcGen - The general traffic control structure containing common parameters (e.g., action, refcnt).
	TcGen
	// Eaction - Defines the mirroring or redirection action, which can be ingress or egress mirror/redirect.
	Eaction int32
	// Ifindex - The index of the network interface to which the packets are redirected or mirrored.
	Ifindex uint32
}
