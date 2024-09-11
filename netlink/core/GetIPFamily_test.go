package core

import (
	"net"
	"testing"
)

func TestGetIPFamily(t *testing.T) {
	tests := []struct {
		name     string
		ip       net.IP
		expected int
	}{
		//0
		{
			name:     "IPv4 address",
			ip:       net.ParseIP("192.168.1.1"),
			expected: AfInet,
		},
		//1
		{
			name:     "IPv6 address",
			ip:       net.ParseIP("2001:db8::68"),
			expected: AfInet6,
		},
		//2
		{
			name:     "Unspecified IP (IPv4)",
			ip:       net.IPv4zero,
			expected: AfInet,
		},
		//3
		{
			name:     "Unspecified IP (IPv6)",
			ip:       net.IPv6zero,
			expected: AfInet6,
		},
		//4
		{
			name:     "Empty IP",
			ip:       net.IP{},
			expected: AfInet,
		},
	}

	for lineNo, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			family := GetIPFamily(tt.ip)
			if family != tt.expected {
				t.Errorf("GetIPFamily() failed\n"+
					"         lineNo: %s (line %d)\n"+
					"             ip: %v\n"+
					"  actual family: %d\n"+
					"expected family: %d",
					tt.name, lineNo, tt.ip, family, tt.expected)
			}
		})
	}
}
