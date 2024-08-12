package network

import (
	"net"
)

// GetNetworkAddress takes a given CIDR string and returns the network address.
func GetNetworkAddress(cidr string) (string, error) {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}

	// Get the network address from the parsed CIDR
	networkAddr := ip.Mask(ipNet.Mask).String()
	return networkAddr, nil
}
