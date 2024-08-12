package network

import "net"

// maskToIPMask converts a 4-byte IP address mask to a net.IPMask and calculates prefix length.
func maskToIPMask(mask net.IP) net.IPMask {
	maskBytes := mask.To4()
	if maskBytes == nil {
		return nil
	}

	onesCount := 0
	for _, b := range maskBytes {
		for b > 0 {
			if b&1 == 1 {
				onesCount++
			}
			b >>= 1
		}
	}
	return net.CIDRMask(onesCount, 32)
}
