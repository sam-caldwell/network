package core

// TcNetemCorr - Struct to represent correlation factors for network emulation.
// This structure defines various correlation factors for delay, loss, and duplication of packets.
// These correlations are used in the Linux Traffic Control system (NetEm) to simulate real-world network behaviors
// by adding a probabilistic correlation between subsequent packets for delay, loss, and duplication.
//
// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcNetemCorr struct {

	//DelayCorr - delay correlation - percentage correlation between consecutive packets experiencing delay
	DelayCorr uint32

	//LossCorr - packet loss correlation - percentage correlation between consecutive packets being lost
	LossCorr uint32

	//DupCorr - duplicate correlation - percentage correlation between consecutive packets being duplicated
	DupCorr uint32
}
