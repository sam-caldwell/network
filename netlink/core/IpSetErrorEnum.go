package core

import (
	"errors"
	"strconv"
)

// IpSetErrorEnum -  represents an error type for IP sets in the Netlink interface.
//
// In the context of Netlink communication, IP sets are used to manage sets of IP addresses, ports, and other elements in the kernel.
// Errors in managing these sets are represented using error codes that are often passed as pointers or special values.
//
// This type mirrors the error handling mechanisms in the Linux kernel for IP sets, particularly in the Netfilter subsystem.
//
// References:
// - Linux Kernel IPSet: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/ipset/ip_set.h
// - Netlink Error Handling: https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h
type IpSetErrorEnum uintptr

const (
	// IpsetErrPrivate - IPSET_ERR_PRIVATE
	IpsetErrPrivate = 4096 + iota

	// IpsetErrProtocol - IPSET_ERR_PROTOCOL
	IpsetErrProtocol

	// IpsetErrFindType - IPSET_ERR_FIND_TYPE
	IpsetErrFindType

	// IpsetErrMaxSets - IPSET_ERR_MAX_SETS
	IpsetErrMaxSets

	// IpsetErrBusy - IPSET_ERR_BUSY
	IpsetErrBusy

	// IpsetErrExistSetname2 - IPSET_ERR_EXIST_SETNAME2
	IpsetErrExistSetname2

	// IpsetErrTypeMismatch - IPSET_ERR_TYPE_MISMATCH
	IpsetErrTypeMismatch

	// IpsetErrExist - IPSET_ERR_EXIST
	IpsetErrExist

	// IpsetErrInvalidCidr - IPSET_ERR_INVALID_CIDR
	IpsetErrInvalidCidr

	// IpsetErrInvalidNetmask - IPSET_ERR_INVALID_NETMASK
	IpsetErrInvalidNetmask

	// IpsetErrInvalidFamily - IPSET_ERR_INVALID_FAMILY
	IpsetErrInvalidFamily

	// IpsetErrTimeout - IPSET_ERR_TIMEOUT
	IpsetErrTimeout

	// IpsetErrReferenced - IPSET_ERR_REFERENCED
	IpsetErrReferenced

	// IpsetErrIpaddrIpv4 - IPSET_ERR_IPADDR_IPV4
	IpsetErrIpaddrIpv4

	// IpsetErrIpaddrIpv6 - IPSET_ERR_IPADDR_IPV6
	IpsetErrIpaddrIpv6

	// IpsetErrCounter - IPSET_ERR_COUNTER
	IpsetErrCounter

	// IpsetErrComment - IPSET_ERR_COMMENT
	IpsetErrComment

	// IpsetErrInvalidMarkmask - IPSET_ERR_INVALID_MARKMASK
	IpsetErrInvalidMarkmask

	// IpsetErrSkbinfo - IPSET_ERR_SKBINFO
	IpsetErrSkbinfo

	//IpsetErrTypeSpecific - IPSET_ERR_TYPE_SPECIFIC -  Type specific error codes
	IpsetErrTypeSpecific = 4352
)

// Error - Return the error for a given IpSetError
func (e IpSetErrorEnum) Error() error {
	return errors.New(e.ErrorString())
}

// ErrorString - Return the string error for a given IpSetError
func (e IpSetErrorEnum) ErrorString() string {
	switch int(e) {
	case IpsetErrPrivate:
		return "private"
	case IpsetErrProtocol:
		return "invalid protocol"
	case IpsetErrFindType:
		return "invalid type"
	case IpsetErrMaxSets:
		return "max sets reached"
	case IpsetErrBusy:
		return "busy"
	case IpsetErrExistSetname2:
		return "exist_setname2"
	case IpsetErrTypeMismatch:
		return "type mismatch"
	case IpsetErrExist:
		return "exist"
	case IpsetErrInvalidCidr:
		return "invalid cidr"
	case IpsetErrInvalidNetmask:
		return "invalid netmask"
	case IpsetErrInvalidFamily:
		return "invalid family"
	case IpsetErrTimeout:
		return "timeout"
	case IpsetErrReferenced:
		return "referenced"
	case IpsetErrIpaddrIpv4:
		return "invalid ipv4 address"
	case IpsetErrIpaddrIpv6:
		return "invalid ipv6 address"
	case IpsetErrCounter:
		return "invalid counter"
	case IpsetErrComment:
		return "invalid comment"
	case IpsetErrInvalidMarkmask:
		return "invalid markmask"
	case IpsetErrSkbinfo:
		return "skbinfo"
	default:
		return "errno " + strconv.Itoa(int(e))
	}
}
