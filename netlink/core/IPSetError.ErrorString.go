package core

import (
	"strconv"
)

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
