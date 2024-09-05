package core

// TcaU32Enum - Enumeration of Traffic Control Action (TCA) U32 filter attributes.
//
// The U32 filter is a highly flexible and programmable classifier used in Linux Traffic Control (TC).
// It allows for complex filtering based on the bitmask and comparison of arbitrary packet data.
//
// This enum maps to U32 filter attributes as defined in the Linux kernel source code.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
type TcaU32Enum uint8

const (
	// TcaU32Unspec - TCA_U32_UNSPEC - Unspecified U32 filter attribute.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Unspec TcaU32Enum = iota

	// TcaU32ClassID - TCA_U32_CLASSID - Class ID for matching packets.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32ClassID

	// TcaU32Hash - TCA_U32_HASH - Hash value used for packet classification.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Hash

	// TcaU32Link - TCA_U32_LINK - Link to another U32 filter.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Link

	// TcaU32Divisor - TCA_U32_DIVISOR - Divisor for hash table.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Divisor

	// TcaU32Sel - TCA_U32_SEL - Selector configuration for filtering specific fields.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Sel

	// TcaU32Police - TCA_U32_POLICE - Policing options for rate limiting.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Police

	// TcaU32Act - TCA_U32_ACT - Action configuration for U32 filter.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Act

	// TcaU32Indev - TCA_U32_INDEV - Input device for packet matching.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Indev

	// TcaU32Pcnt - TCA_U32_PCNT - Packet count for statistics.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Pcnt

	// TcaU32Mark - TCA_U32_MARK - Mark value for identifying packets.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Mark

	// TcaU32Max - TCA_U32_MAX - Maximum value of U32 filter attributes.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
	TcaU32Max = TcaU32Mark
)
