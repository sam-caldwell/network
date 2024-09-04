package core

// LinkLayer - defines link layer types as an enumerated type.
//
// This represents various link layer types defined in the Linux kernel, as seen in pkt_sched.h.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type LinkLayer uint8

const (
	// LinkLayerUnspec - LINKLAYER_UNSPEC - unspecified link layer type.
	// This is used when no specific link layer is indicated.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	LinkLayerUnspec = iota

	// LinkLayerEthernet - LINKLAYER_ETHERNET - represents Ethernet link layer.
	// Used for devices communicating over Ethernet.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	LinkLayerEthernet

	// LinkLayerAtm - LINKLAYER_ATM - represents Asynchronous Transfer Mode link layer.
	// Used for devices communicating over ATM (Asynchronous Transfer Mode).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	LinkLayerAtm
)

// TcLinklayerMask - TC_LINKLAYER_MASK - limit use to lower 4 bits
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
const TcLinklayerMask = 0x0F
