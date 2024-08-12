package network

import (
	"fmt"
	"net"
)

// StringToIPNet - Parse a string to an IPNet or error
func StringToIPNet(cidr string) (net.IPNet, error) {
	zero := net.IPNet{}
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return zero, fmt.Errorf("invalid CIDR string: %v", err)
	}
	return *ipNet, nil
}
