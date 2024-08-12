package network

import (
	"net"
	"strconv"
)

// hexToIPv4 converts a hexadecimal string to a standard IPv4 address.
func hexToIPv4(hexStr string) string {
	ip := make(net.IP, 4)
	for i := 0; i < 4; i++ {
		value, _ := strconv.ParseUint(hexStr[2*i:2*i+2], 16, 8)
		ip[i] = byte(value)
	}
	return ip.String()
}
