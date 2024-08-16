package network

import (
	"fmt"
	"net"
	"strings"
)

// StringToIPMask converts a subnet mask string (dotted-decimal or CIDR) to a net.IPMask type.
func StringToIPMask(maskStr string) (net.IPMask, error) {
	if strings.Contains(maskStr, "/") {
		// CIDR notation
		_, ipNet, err := net.ParseCIDR(maskStr)
		if err != nil {
			return nil, fmt.Errorf("invalid CIDR string: %v", err)
		}
		return ipNet.Mask, nil
	}

	// Dotted-decimal notation for IPv4
	ip := net.ParseIP(maskStr)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address: %s", maskStr)
	}

	if ip.To4() != nil {
		mask := net.ParseIP(maskStr).To4()
		if mask == nil {
			return nil, fmt.Errorf("invalid IPv4 address: %s", maskStr)
		}
		return MaskToIPMask(mask), nil
	}

	// IPv6 address handling
	if ip.To16() != nil {
		// For IPv6, assume a full /128 mask
		return net.CIDRMask(128, 128), nil
	}

	return nil, fmt.Errorf("invalid subnet mask: %s", maskStr)
}
