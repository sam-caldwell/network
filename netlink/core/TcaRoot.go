package core

// TcaRootEnum - Enumeration for TCA root attributes.
//
// These attributes are used for managing Traffic Control Actions (TCA) in the Linux rtnetlink system.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
type TcaRootEnum uint8

const (
	// TcaRootUnspec - TCA_ROOT_UNSPEC - unspecified root attribute.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaRootUnspec TcaRootEnum = iota

	// TcaRootTab - TCA_ROOT_TAB - root tab attribute.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaRootTab TcaRootEnum = iota

	// TcaRootFlags - TCA_ROOT_FLAGS - root flags attribute.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaRootFlags TcaRootEnum = iota

	// TcaRootCount - TCA_ROOT_COUNT - root count attribute.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaRootCount TcaRootEnum = iota

	// TcaRootTimeDelta - TCA_ROOT_TIME_DELTA - time delta in milliseconds.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaRootTimeDelta TcaRootEnum = iota

	// TcaRootExtWarnMsg - TCA_ROOT_EXT_WARN_MSG - external warning message.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaRootExtWarnMsg TcaRootEnum = iota

	// TcaRootMax - TCA_ROOT_MAX - maximum value for root attributes.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaRootMax = iota - 1

	/*

	 */

	// TcaActTab - TCA_ACT_TAB -
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActTab = TcaRootTab

	// TcaaMax - TCAA_MAX -
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaaMax = TcaRootTab
)
