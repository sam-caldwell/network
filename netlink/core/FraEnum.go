package core

// FraEnum represents rule attribute types used in routing rules in Linux.
// These constants define the attributes used for filtering and processing rules,
// such as destination, source, interface names, and more.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
type FraEnum uint8

const (
	// FraUnspec - FRA_UNSPEC - Unspecified rule attribute
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraUnspec FraEnum = iota

	// FraDst - FRA_DST - Destination address for rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraDst

	// FraSrc - FRA_SRC - Source address for rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraSrc

	// FraIifName - FRA_IIFNAME - Incoming interface name
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraIifName

	// FraGoto - FRA_GOTO - Target to jump to (FR_ACT_GOTO)
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraGoto

	// FraUnused2 - FRA_UNUSED2 - Reserved/unused attribute
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraUnused2

	// FraPriority - FRA_PRIORITY - Rule priority/preference
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraPriority

	// FraUnused3 - FRA_UNUSED3 - Reserved/unused attribute
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraUnused3

	// FraUnused4 - FRA_UNUSED4 - Reserved/unused attribute
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraUnused4

	// FraUnused5 - FRA_UNUSED5 - Reserved/unused attribute
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraUnused5

	// FraFwMark - FRA_FWMARK - Firewall mark for rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraFwMark

	// FraFlow - FRA_FLOW - Flow/class ID for rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraFlow

	// FraTunId - FRA_TUN_ID - Tunnel ID for rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraTunId

	// FraSuppressIfGroup - FRA_SUPPRESS_IFGROUP - Suppress based on interface group
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraSuppressIfGroup

	// FraSuppressPrefixLen - FRA_SUPPRESS_PREFIXLEN - Suppress based on prefix length
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraSuppressPrefixLen

	// FraTable - FRA_TABLE - Extended table ID for rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraTable

	// FraFwMask - FRA_FWMASK - Mask for firewall mark
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraFwMask

	// FraOifName - FRA_OIFNAME - Outgoing interface name
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraOifName

	// FraPad - FRA_PAD - Padding for alignment
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraPad

	// FraL3Mdev - FRA_L3MDEV - Rule for L3 master device
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraL3Mdev

	// FraUidRange - FRA_UID_RANGE - User ID range for rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraUidRange

	// FraProtocol - FRA_PROTOCOL - Originator protocol of the rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraProtocol

	// FraIpProto - FRA_IP_PROTO - IP protocol for rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraIpProto

	// FraSportRange - FRA_SPORT_RANGE - Source port range for rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraSportRange

	// FraDportRange - FRA_DPORT_RANGE - Destination port range for rule
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/fib_rules.h
	FraDportRange
)
