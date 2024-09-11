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
		{
			name:     "IPv4 address",
			ip:       net.ParseIP("192.168.1.1"),
			expected: AfInet,
		},
		{
			name:     "IPv6 address",
			ip:       net.ParseIP("2001:db8::68"),
			expected: AfInet6,
		},
		{
			name:     "Unspecified IP (IPv4)",
			ip:       net.IPv4zero,
			expected: AfInet,
		},
		{
			name:     "Unspecified IP (IPv6)",
			ip:       net.IPv6zero,
			expected: AfInet6,
		},
		{
			name:     "Empty IP",
			ip:       net.IP{},
			expected: AfUnspec,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			family := GetIPFamily(tt.ip)
			if family != tt.expected {
				t.Errorf("GetIPFamily(%v) = %d; want %d", tt.ip, family, tt.expected)
			}
		})
	}
}
