package network

import "net"

// StringToIP - Given a string, parse to an IP or return empty on error
func StringToIP(ipStr string) net.IP {
	return net.ParseIP(ipStr)
}
