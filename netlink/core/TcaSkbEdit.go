package core

// TcSkbEdit - A structure representing a SKB (Socket Buffer) Edit action.
// This struct is used in Linux Traffic Control (TC) to modify or manipulate fields of SKB packets in transit.
//
// TcSkbEdit wraps around the base TcGen struct and provides general traffic control action capabilities.
//
// The TcGen struct contains common fields like index, capabilities, and reference counts.
// TcSkbEdit inherits these fields and extends the functionality with the ability to edit packet contents.
//
// References:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
type TcSkbEdit struct {
	// Inherits the TcGen fields such as index, action, and capabilities.
	TcGen
}
