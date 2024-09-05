package core

// TcaAct - Enumeration for Traffic Control Action (TCA) attributes.
//
// These constants represent various attributes related to traffic control actions in the Linux kernel.
//
// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
type TcaAct uint8

const (
	// TcaActUnspec - TCA_ACT_UNSPEC - Unspecified action type.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActUnspec TcaAct = iota

	// TcaActKind - TCA_ACT_KIND - Describes the type of action.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActKind TcaAct = iota

	// TcaActOptions - TCA_ACT_OPTIONS - Options for the action.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActOptions TcaAct = iota

	// TcaActIndex - TCA_ACT_INDEX - Index of the action.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActIndex TcaAct = iota

	// TcaActStats - TCA_ACT_STATS - Statistics for the action.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActStats TcaAct = iota

	// TcaActPad - TCA_ACT_PAD - Padding for alignment.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActPad TcaAct = iota

	// TcaActCookie - TCA_ACT_COOKIE - Cookie for the action.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActCookie TcaAct = iota

	// TcaActFlags - TCA_ACT_FLAGS - Flags for the action.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActFlags TcaAct = iota

	// TcaActHwStats - TCA_ACT_HW_STATS - Hardware statistics for the action.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActHwStats TcaAct = iota

	// TcaActUsedHwStats - TCA_ACT_USED_HW_STATS - Indicates usage of hardware statistics.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActUsedHwStats TcaAct = iota

	// TcaActInHwCount - TCA_ACT_IN_HW_COUNT - Number of times the action was used in hardware.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActInHwCount TcaAct = iota

	// TcaActMax - TCA_ACT_MAX - Maximum valid value for TCA actions.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActMax TcaAct = iota - 1
)

const (
	// TcaActGact - TCA_ACT_GACT
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	TcaActGact = 5
)
