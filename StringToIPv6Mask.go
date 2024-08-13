package network

import (
	"fmt"
	"net"
	"strings"
)

// StringToIPv6Mask returns the IPv6 mask for a given IPv6 address in CIDR notation.
func StringToIPv6Mask(cidr string) (net.IPMask, error) {
	// Remove square brackets if present
	parts := strings.Split(cidr, "/")
	cidr = strings.TrimSuffix(strings.TrimPrefix(parts[0], "["), "]")
	if len(parts) > 1 {
		cidr = fmt.Sprintf("%s/%s", cidr, parts[1])
	}
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	return ipNet.Mask, nil
}
