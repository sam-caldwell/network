package network

import "net"

// CompareIPMask compares two net.IPMask values for equality.
func CompareIPMask(a, b net.IPMask) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
