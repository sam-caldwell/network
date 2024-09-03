package core

import (
	"strconv"
)

// ErrorString - Return the string error for a given IpSetError
func (e IPSetError) ErrorString() string {
	switch int(e) {
	case IPSET_ERR_PRIVATE:
		return "private"
	case IPSET_ERR_PROTOCOL:
		return "invalid protocol"
	case IPSET_ERR_FIND_TYPE:
		return "invalid type"
	case IPSET_ERR_MAX_SETS:
		return "max sets reached"
	case IPSET_ERR_BUSY:
		return "busy"
	case IPSET_ERR_EXIST_SETNAME2:
		return "exist_setname2"
	case IPSET_ERR_TYPE_MISMATCH:
		return "type mismatch"
	case IPSET_ERR_EXIST:
		return "exist"
	case IPSET_ERR_INVALID_CIDR:
		return "invalid cidr"
	case IPSET_ERR_INVALID_NETMASK:
		return "invalid netmask"
	case IPSET_ERR_INVALID_FAMILY:
		return "invalid family"
	case IPSET_ERR_TIMEOUT:
		return "timeout"
	case IPSET_ERR_REFERENCED:
		return "referenced"
	case IPSET_ERR_IPADDR_IPV4:
		return "invalid ipv4 address"
	case IPSET_ERR_IPADDR_IPV6:
		return "invalid ipv6 address"
	case IPSET_ERR_COUNTER:
		return "invalid counter"
	case IPSET_ERR_COMMENT:
		return "invalid comment"
	case IPSET_ERR_INVALID_MARKMASK:
		return "invalid markmask"
	case IPSET_ERR_SKBINFO:
		return "skbinfo"
	default:
		return "errno " + strconv.Itoa(int(e))
	}
}
