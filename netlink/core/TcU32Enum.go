package core

// TC_U32 flags for packet classification in traffic control (tc)
//
// These constants define various behaviors for the u32 classifier in Linux tc.
// The `u32` classifier allows matching on specific 32-bit values in packet headers.
// The constants are used as flags to control how the classifier interprets and processes
// the matching rules.
//
// See:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
// - https://man7.org/linux/man-pages/man8/tc-u32.8.html
const (
	// TcU32Terminal - TC_U32_TERMINAL - Stop processing further rules once this rule matches.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcU32Terminal = 1 << iota

	// TcU32Offset - TC_U32_OFFSET - Allow offsets in the u32 matching process.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcU32Offset

	// TcU32Varoffset - TC_U32_VAROFFSET - Enable the use of variable offsets.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcU32Varoffset

	// TcU32Eat - TC_U32_EAT - Consumes the packet once matched.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcU32Eat
)
