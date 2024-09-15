package core

// Netlink Message Constants
//
// This file defines constants used in Netlink communication between user space and
// the Linux kernel. These constants correspond to message types, flags, and other
// parameters defined in the Linux kernel headers, particularly in [linux/netlink.h>`
// and [linux/rtnetlink.h>`.
//
// References:
// - Netlink definitions: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
// - Routing Netlink definitions: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h

const (
	// Alignment and Header Length Constants

	// NlmsgAlignTo - NLMSG_ALIGNTO - represents the alignment boundary for Netlink messages.
	// Netlink messages are aligned to NLMSG_ALIGNTO bytes (4 bytes).
	// This ensures proper alignment of messages in memory.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L60)
	NlmsgAlignTo = 0x4

	// NlmsgHdrLen represents the length of the Netlink message header.
	// The `nlmsghdr` structure is 16 bytes in length.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h)
	NlmsgHdrLen = 0x10 // 16 bytes

	// NlmsgMinType - NLMSG_MIN_TYPE- defines the minimum value for application-defined message types.
	// Message types below this value are reserved for system use.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L57)
	NlmsgMinType = 0x10 // 16
)

const (
	// Netlink Message Types (from [linux/netlink.h>`)

	// NlmsgNoop - NLMSG_NOOP - is a Netlink message type used for a no-operation message.
	// It is ignored by the kernel and can be used for debugging or synchronization.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L53)
	NlmsgNoop = 0x1

	// NlmsgError - NLMSG_ERROR - is a Netlink message type indicating an error.
	// It contains an error code and can include the original Netlink message that caused the error.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L54)
	NlmsgError = 0x2

	// NlmsgDone - NLMSG_DONE - is a Netlink message type indicating the end of a multipart message.
	// It is used as a terminator for multipart messages.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L55)
	NlmsgDone = 0x3

	// NlmsgOverrun - NLMSG_OVERRUN - is a Netlink message type indicating data loss due to buffer overrun.
	// This message type is used when the kernel fails to deliver all data due to buffer limitations.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L56)
	NlmsgOverrun = 0x4
)

const (
	// Netlink Message Flags (from [linux/netlink.h>`)

	// NlmFRequest - NLM_F_REQUEST - indicates that this is a request message.
	// All requests from user space to kernel space should set this flag.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L67)
	NlmFRequest = 0x1

	// NlmFMulti - NLM_F_MULTI - indicates that the message is part of a multipart message terminated by NLMSG_DONE.
	// This flag is set by the kernel when sending multipart messages.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L68)
	NlmFMulti = 0x2

	// NlmFAck - NLM_F_ACK - requests an acknowledgment from the receiver of the message.
	// When set in a request, the kernel will send an acknowledgment message back.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L69)
	NlmFAck = 0x4

	// NlmFEcho - NLM_F_ECHO - requests that the receiver echoes this request back.
	// This can be used for synchronization purposes.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L70)
	NlmFEcho = 0x8

	// NlmFDumpIntr - NLM_F_DUMP_INTR - indicates that a dump was inconsistent due to a sequence change.
	// This flag can be set in a response to indicate that the data may be inconsistent.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L75)
	NlmFDumpIntr = 0x10

	// NlmFDumpFiltered - NLM_F_DUMP_FILTERED - indicates that the dump was filtered.
	// Used in dump responses to indicate that the data was filtered according to criteria.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L76)
	NlmFDumpFiltered = 0x20
)

const (
	// Netlink Flags for Get and Dump Requests (from [linux/netlink.h>`)

	// NlmFRoot - NLM_F_ROOT - specifies that the operation should be performed from the root.
	// Often used in dump requests to indicate starting from the root of a table.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L73)
	NlmFRoot = 0x100

	// NlmFMatch - NLM_F_MATCH - specifies that the dump should only include entries matching the criteria.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L74)
	NlmFMatch = 0x200

	// NlmFAtomic - NLM_F_ATOMIC - requests an atomic dump.
	// The operation should be performed atomically to provide a consistent snapshot.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L71)
	NlmFAtomic = 0x400

	// NlmFDump - NLM_F_DUMP - is a combination of NLM_F_ROOT and NLM_F_MATCH, used to request a complete dump.
	// NLM_F_DUMP = NLM_F_ROOT | NLM_F_MATCH
	NlmFDump = NlmFRoot | NlmFMatch
)

const (
	// Netlink Flags for New Requests (from [linux/rtnetlink.h>`)

	// NlmFReplace - NLM_F_REPLACE - requests that the kernel replaces an existing object with the one supplied.
	//
	// Defined in:
	// - [<linux/rtnetlink.h>](https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h#L87)
	NlmFReplace = 0x100

	// NlmFExcl - NLM_F_EXCL - requests that the kernel does not replace the object if it already exists.
	// The operation should fail if the object exists.
	//
	// Defined in:
	// - [<linux/rtnetlink.h>](https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h#L88)
	NlmFExcl = 0x200

	// NlmFCreate - NLM_F_CREATE - requests that the kernel creates a new object if it does not already exist.
	//
	// Defined in:
	// - [<linux/rtnetlink.h>](https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h#L89)
	NlmFCreate = 0x400

	// NlmFAppend - NLM_F_APPEND - requests that the kernel adds the object to the end of the object list.
	//
	// Defined in:
	// - [<linux/rtnetlink.h>](https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h#L90)
	NlmFAppend = 0x800
)

const (
	// Additional Netlink Flags

	// NlmFNonrec - NLM_F_NONREC - indicates a non-recursive dump request.
	// Used to request that the kernel does not recursively dump linked objects.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L72)
	NlmFNonrec = 0x100

	// NlmFCapped - NLM_F_CAPPED - indicates that the data was truncated due to the buffer size limit.
	// Used in responses to indicate that the data was capped.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L77)
	NlmFCapped = 0x100

	// NlmFAckTlvs - NLM_F_ACK_TLVS - requests that the kernel include extended ACK information in the response.
	// This includes additional TLVs (Type-Length-Value) data in the acknowledgment.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L78)
	NlmFAckTlvs = 0x200

	// NlmFBulk - NLM_F_BULK - requests a bulk operation.
	// Used to indicate that the message contains multiple operations to be processed together.
	//
	// Defined in:
	// - [linux/netlink.h](https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h#L79)
	NlmFBulk = 0x200
)

// Routing Netlink Message Types and Flags
//
// This file defines constants used for routing Netlink messages
// between user space and the Linux kernel. These constants correspond
// to message types, flags, and parameters defined in the Linux kernel,
// particularly in `<linux/rtnetlink.h>`.
//
// References:
// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h

const (
	// RtmBase corresponds to RTM_BASE.
	// It defines the base value for Netlink routing messages.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmBase = 0x10

	// RtmDeleteAction corresponds to RTM_DELACTION.
	// Used to delete an action from the kernel action table.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteAction = 0x31

	// RtmDeleteAddr corresponds to RTM_DELADDR.
	// Used to delete an IP address from an interface.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteAddr = 0x15

	// RtmDeleteAddrLabel corresponds to RTM_DELADDRLABEL.
	// Used to delete address labels for IPv6 addresses.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteAddrLabel = 0x49

	// RtmDeleteChain corresponds to RTM_DELCHAIN.
	// Used to delete a rule chain.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteChain = 0x65

	// RtmDeleteLink corresponds to RTM_DELLINK.
	// Used to delete a network interface.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteLink = 0x11

	// RtmDeleteLinkProp corresponds to RTM_DELLINKPROP.
	// Used to delete properties of a link.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteLinkProp = 0x6d

	// RtmDeleteMdb corresponds to RTM_DELMDB.
	// Used to delete a multicast database entry.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteMdb = 0x55

	// RtmDeleteNeigh corresponds to RTM_DELNEIGH.
	// Used to delete a neighbor cache entry.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteNeigh = 0x1d

	// RtmDeleteNetConf corresponds to RTM_DELNETCONF.
	// Used to delete network configuration.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteNetConf = 0x51

	// RtmDeleteNextHop corresponds to RTM_DELNEXTHOP.
	// Used to delete a nexthop entry.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteNextHop = 0x69

	// RtmDeleteNextHopBucket corresponds to RTM_DELNEXTHOPBUCKET.
	// Used to delete a nexthop bucket.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteNextHopBucket = 0x75

	// RtmDeleteNsId corresponds to RTM_DELNSID.
	// Used to delete a network namespace identifier.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteNsId = 0x59

	// RtmDelQdisc corresponds to RTM_DELQDISC.
	// Used to delete a queuing discipline.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDelQdisc = 0x25

	// RtmDeleteRoute corresponds to RTM_DELROUTE.
	// Used to delete a routing table entry.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteRoute = 0x19

	// RtmDeleteRule corresponds to RTM_DELRULE.
	// Used to delete a routing rule.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteRule = 0x21

	// RtmDeleteTClass corresponds to RTM_DELTCLASS.
	// Used to delete a traffic control class.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteTClass = 0x29

	// RtmDeleteTFilter corresponds to RTM_DELTFILTER.
	// Used to delete a traffic control filter.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteTFilter = 0x2d

	// RtmDeleteTunnel corresponds to RTM_DELTUNNEL.
	// Used to delete a tunnel interface.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteTunnel = 0x79

	// RtmDeleteVlan corresponds to RTM_DELVLAN.
	// Used to delete a VLAN interface.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmDeleteVlan = 0x71

	// RtmFCloned corresponds to RTM_F_CLONED.
	// Flag indicating a cloned route.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmFCloned = 0x200

	// RtmFEqualize corresponds to RTM_F_EQUALIZE.
	// Flag indicating multipath equalizer.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmFEqualize = 0x400

	// RtmFFibMatch corresponds to RTM_F_FIB_MATCH.
	// Flag indicating route matches FIB table.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmFFibMatch = 0x2000

	// RtmFLookupTable corresponds to RTM_F_LOOKUP_TABLE.
	// Flag to request a specific routing table.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmFLookupTable = 0x1000

	// RtmFNotify corresponds to RTM_F_NOTIFY.
	// Flag indicating notifications requested.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmFNotify = 0x100

	// RtmFOffload corresponds to RTM_F_OFFLOAD.
	// Flag indicating the route is offloaded.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmFOffload = 0x4000

	// RtmFOffloadFailed corresponds to RTM_F_OFFLOAD_FAILED.
	// Flag indicating route offload failed.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmFOffloadFailed = 0x20000000

	// RtmFPrefix corresponds to RTM_F_PREFIX.
	// Flag indicating prefix address.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmFPrefix = 0x800

	// RtmFTrap corresponds to RTM_F_TRAP.
	// Flag indicating a trap route.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmFTrap = 0x8000

	// RtmGetAction corresponds to RTM_GETACTION.
	// Used to retrieve actions from the kernel action table.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetAction = 0x32

	// RtmGetAddr corresponds to RTM_GETADDR.
	// Used to retrieve IP addresses from interfaces.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetAddr = 0x16

	// RtmGetAddrLabel corresponds to RTM_GETADDRLABEL.
	// Used to retrieve address labels for IPv6 addresses.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetAddrLabel = 0x4a

	// RtmGetAnycast corresponds to RTM_GETANYCAST.
	// Used to retrieve anycast addresses.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetAnycast = 0x3e

	// RtmGetChain corresponds to RTM_GETCHAIN.
	// Used to get rule chains.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetChain = 0x66

	// RtmGetdcb corresponds to RTM_GETDCB.
	// Used to retrieve DCB (Data Center Bridging) configuration.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetdcb = 0x4e

	// RtmGetLink corresponds to RTM_GETLINK.
	// Used to retrieve network interface information.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetLink = 0x12

	// RtmGetLinkProp corresponds to RTM_GETLINKPROP.
	// Used to retrieve link properties.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetLinkProp = 0x6e

	// RtmGetMdb corresponds to RTM_GETMDB.
	// Used to retrieve multicast database entries.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetMdb = 0x56

	// RtmGetMulticast corresponds to RTM_GETMULTICAST.
	// Used to retrieve multicast addresses.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetMulticast = 0x3a

	// RtmGetNeigh corresponds to RTM_GETNEIGH.
	// Used to retrieve neighbor cache entries.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetNeigh = 0x1e

	// RtmGetNeighTbl corresponds to RTM_GETNEIGHTBL.
	// Used to retrieve neighbor tables.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetNeighTbl = 0x42

	// RtmGetNetConf corresponds to RTM_GETNETCONF.
	// Used to retrieve network configuration.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetNetConf = 0x52

	// RtmGetNextHop corresponds to RTM_GETNEXTHOP.
	// Used to retrieve nexthop entries.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetNextHop = 0x6a

	// RtmGetNextHopBucket corresponds to RTM_GETNEXTHOPBUCKET.
	// Used to retrieve nexthop buckets.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetNextHopBucket = 0x76

	// RtmGetNsId corresponds to RTM_GETNSID.
	// Used to retrieve network namespace identifiers.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetNsId = 0x5a

	// RtmGetQdisc corresponds to RTM_GETQDISC.
	// Used to retrieve queuing disciplines.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetQdisc = 0x26

	// RtmGetRoute corresponds to RTM_GETROUTE.
	// Used to retrieve routing table entries.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetRoute = 0x1a

	// RtmGetRule corresponds to RTM_GETRULE.
	// Used to retrieve routing rules.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetRule = 0x22

	// RtmGetStats corresponds to RTM_GETSTATS.
	// Used to retrieve statistics.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetStats = 0x5e

	// RtmGetTClass corresponds to RTM_GETTCLASS.
	// Used to retrieve traffic control classes.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetTClass = 0x2a

	// RtmGetTFilter corresponds to RTM_GETTFILTER.
	// Used to retrieve traffic control filters.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetTFilter = 0x2e

	// RtmGetTunnel corresponds to RTM_GETTUNNEL.
	// Used to retrieve tunnel interfaces.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetTunnel = 0x7a

	// RtmGetVlan corresponds to RTM_GETVLAN.
	// Used to retrieve VLAN interfaces.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmGetVlan = 0x72

	// RtmMax corresponds to RTM_MAX.
	// Defines the maximum message type.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmMax = 0x7b

	// RtmNewAction corresponds to RTM_NEWACTION.
	// Used to add a new action to the kernel action table.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewAction = 0x30

	// RtmNewAddr corresponds to RTM_NEWADDR.
	// Used to add a new IP address to an interface.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewAddr = 0x14

	// RtmNewAddrLabel corresponds to RTM_NEWADDRLABEL.
	// Used to add new address labels for IPv6 addresses.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewAddrLabel = 0x48

	// RtmNewCacheReport corresponds to RTM_NEWCACHEREPORT.
	// Used for cache report messages.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewCacheReport = 0x60

	// RtmNewChain corresponds to RTM_NEWCHAIN.
	// Used to add a new rule chain.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewChain = 0x64

	// RtmNewLink corresponds to RTM_NEWLINK.
	// Used to add or change a network interface.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewLink = 0x10

	// RtmNewLinkProp corresponds to RTM_NEWLINKPROP.
	// Used to set properties of a link.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewLinkProp = 0x6c

	// RtmNewMdb corresponds to RTM_NEWMDB.
	// Used to add a new multicast database entry.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewMdb = 0x54

	// RtmNewNdUserOpt corresponds to RTM_NEWNDUSEROPT.
	// Used to add new neighbor discovery user options.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewNdUserOpt = 0x44

	// RtmNewNeigh corresponds to RTM_NEWNEIGH.
	// Used to add a new neighbor cache entry.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewNeigh = 0x1c

	// RtmNewNeighTbl corresponds to RTM_NEWNEIGHTBL.
	// Used to add a new neighbor table.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewNeighTbl = 0x40

	// RtmNewNetConf corresponds to RTM_NEWNETCONF.
	// Used to add new network configuration.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewNetConf = 0x50

	// RtmNewNextHop corresponds to RTM_NEWNEXTHOP.
	// Used to add a new nexthop entry.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewNextHop = 0x68

	// RtmNewNextHopBucket corresponds to RTM_NEWNEXTHOPBUCKET.
	// Used to add a new nexthop bucket.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewNextHopBucket = 0x74

	// RtmNewNsId corresponds to RTM_NEWNSID.
	// Used to add a new network namespace identifier.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewNsId = 0x58

	// RtmNewNVlan corresponds to RTM_NEWNVLAN.
	// Used to add a new VLAN interface.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewNVlan = 0x70

	// RtmNewPrefix corresponds to RTM_NEWPREFIX.
	// Used to add a new prefix route.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewPrefix = 0x34

	// RtmNewQdisc corresponds to RTM_NEWQDISC.
	// Used to add a new queuing discipline.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewQdisc = 0x24

	// RtmNewRoute corresponds to RTM_NEWROUTE.
	// Used to add a new routing table entry.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewRoute = 0x18

	// RtmNewRule corresponds to RTM_NEWRULE.
	// Used to add a new routing rule.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewRule = 0x20

	// RtmNewStats corresponds to RTM_NEWSTATS.
	// Used to send new statistics.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewStats = 0x5c

	// RtmNewTClass corresponds to RTM_NEWTCLASS.
	// Used to add a new traffic control class.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewTClass = 0x28

	// RtmNewTFilter corresponds to RTM_NEWTFILTER.
	// Used to add a new traffic control filter.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewTFilter = 0x2c

	// RtmNewTunnel corresponds to RTM_NEWTUNNEL.
	// Used to add a new tunnel interface.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNewTunnel = 0x78

	// RtmNrFamilies corresponds to RTM_NR_FAMILIES.
	// Number of address families.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNrFamilies = 0x1b

	// RtmNrMsgTypes corresponds to RTM_NR_MSGTYPES.
	// Number of message types.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmNrMsgTypes = 0x6c

	// RtmSetDcb corresponds to RTM_SETDCB.
	// Used to set DCB configuration.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmSetDcb = 0x4f

	// RtmSetLink corresponds to RTM_SETLINK.
	// Used to set the state of a network interface.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmSetLink = 0x13

	// RtmSetNeighTbl corresponds to RTM_SETNEIGHTBL.
	// Used to set neighbor table parameters.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmSetNeighTbl = 0x43

	// RtmSetStats corresponds to RTM_SETSTATS.
	// Used to set statistics.
	//
	// References:
	// - Linux kernel source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
	RtmSetStats = 0x5f
)
