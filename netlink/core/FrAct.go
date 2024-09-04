package core

// FrActions defines rule actions in the Linux routing system.
// These actions dictate how packets are treated when they match certain rules.
// Examples include routing to a fixed table, dropping packets, or returning errors.
//
// For more details, refer to: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
type FrActions uint8

// Rule action types for IP rules
const (
	// FrActUnspec - FR_ACT_UNSPEC - Unspecified action
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FrActUnspec FrActions = iota

	// FrActToTbl - FR_ACT_TO_TBL - Pass packets to a fixed table
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FrActToTbl

	// FrActGoto - FR_ACT_GOTO - Jump to another rule
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FrActGoto

	// FrActNop - FR_ACT_NOP - No operation (rule is a no-op)
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FrActNop

	// FrActRes3 - FR_ACT_RES3 - Reserved
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FrActRes3

	// FrActRes4 - FR_ACT_RES4 - Reserved
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FrActRes4

	// FrActBlackhole - FR_ACT_BLACKHOLE - Drop packets without notification
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FrActBlackhole

	// FrActUnreachable - FR_ACT_UNREACHABLE - Drop packets and return ENETUNREACH error
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FrActUnreachable

	// FrActProhibit - FR_ACT_PROHIBIT - Drop packets and return EACCES error
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FrActProhibit
)
