package core

// IpSetCadtAttributeFlags - Flags at CADT attribute level, upper half of command attributes.
//
// These flags are used in the context of IP sets within the Netfilter framework, particularly at the level of
// Create/Append/Delete/Test (CADT) attributes. The flags define various options and behaviors for handling IP sets.
//
// References:
// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
type IpSetCadtAttributeFlags uint8

const (

	// IpsetFlagBitBefore - IPSET_FLAG_BIT_BEFORE - Position before other rules (bit position 0).
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagBitBefore IpSetCadtAttributeFlags = 0

	// IpsetFlagBefore - IPSET_FLAG_BEFORE - Set when the IP set entry should be positioned before other rules.
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagBefore = 1 << IpsetFlagBitBefore

	// IpsetFlagBitPhysdev - IPSET_FLAG_BIT_PHYSDEV - Match on physical device (bit position 1).
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagBitPhysdev IpSetCadtAttributeFlags = 1

	// IpsetFlagPhysdev - IPSET_FLAG_PHYSDEV - Set when matching on the physical device (bridging).
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagPhysdev = 1 << IpsetFlagBitPhysdev

	// IpsetFlagBitNomatch - IPSET_FLAG_BIT_NOMATCH - Invert match (bit position 2).
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagBitNomatch IpSetCadtAttributeFlags = 2

	// IpsetFlagNomatch - IPSET_FLAG_NOMATCH - Set when the match should be inverted (i.e., non-matching entries).
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagNomatch = 1 << IpsetFlagBitNomatch

	// IpsetFlagBitWithCounters - IPSET_FLAG_BIT_WITH_COUNTERS - Include counters (bit position 3).
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagBitWithCounters IpSetCadtAttributeFlags = 3

	// IpsetFlagWithCounters - IPSET_FLAG_WITH_COUNTERS - Set when counters are included with the IP set entries.
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagWithCounters = 1 << IpsetFlagBitWithCounters

	// IpsetFlagBitWithComment - IPSET_FLAG_BIT_WITH_COMMENT - Include comments (bit position 4).
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagBitWithComment IpSetCadtAttributeFlags = 4

	// IpsetFlagWithComment - IPSET_FLAG_WITH_COMMENT - Set when comments are included with the IP set entries.
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagWithComment = 1 << IpsetFlagBitWithComment

	// IpsetFlagBitWithForceadd - IPSET_FLAG_BIT_WITH_FORCEADD - Force add entry (bit position 5).
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagBitWithForceadd IpSetCadtAttributeFlags = 5

	// IpsetFlagWithForceadd - IPSET_FLAG_WITH_FORCEADD - Set to force the addition of an entry, regardless of
	// existence.
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagWithForceadd = 1 << IpsetFlagBitWithForceadd

	// IpsetFlagBitWithSkbinfo - IPSET_FLAG_BIT_WITH_SKBINFO - Include skbinfo (bit position 6).
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagBitWithSkbinfo IpSetCadtAttributeFlags = 6

	// IpsetFlagWithSkbinfo - IPSET_FLAG_WITH_SKBINFO - Set when skbinfo (socket buffer info) is included with the
	// IP set entries.
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagWithSkbinfo = 1 << IpsetFlagBitWithSkbinfo

	// IpsetFlagCadtMax - IPSET_FLAG_CADT_MAX - Maximum value for CADT flags.
	//
	// - IPSet CADT Flags: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
	IpsetFlagCadtMax = 15
)
