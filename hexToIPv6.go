package network

import (
	"encoding/hex"
	"net"
)

// hexToIPv6 converts a hexadecimal string to a standard IPv6 address.
func hexToIPv6(hexStr string) string {
	ip := make(net.IP, 16)
	for i := 0; i < 16; i++ {
		value, _ := hex.DecodeString(hexStr[2*i : 2*i+2])
		ip[i] = value[0]
	}
	return ip.String()
}
