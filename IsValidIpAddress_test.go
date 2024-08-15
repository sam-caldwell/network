package network

import "testing"

func TestIsValidIpAddress(t *testing.T) {
	// Test valid IPv4 addresses
	ipv4Addresses := []string{
		"192.0.2.1",
		"10.0.0.1",
		"172.16.0.1",
	}
	for _, ip := range ipv4Addresses {
		if !IsValidIpAddress(ip) {
			t.Fatalf("Expected %s to be a valid IPv4 address", ip)
		}
	}

	// Test valid IPv6 addresses
	ipv6Addresses := map[string]bool{
		"2001:0db8:85a3:0000:0000:8a2e:0370:7334": true,
		"fe80::1":                      true,
		"::1":                          true,
		"fe80::1fa8:4cd:8210:5e87/128": false,
	}
	for ip, valid := range ipv6Addresses {
		if valid && !IsValidIpAddress(ip) || !valid && IsValidIpAddress(ip) {
			t.Fatalf("Expected %s to be a valid IPv6 address", ip)
		}
	}

	// Test invalid IP addresses
	invalidAddresses := []string{
		"notAnIp",
		"256.0.0.1",
		"2001:0db8:85a3::8a2e:0370::7334",
	}
	for _, ip := range invalidAddresses {
		if IsValidIpAddress(ip) {
			t.Fatalf("Expected %s to be an invalid IP address", ip)
		}
	}
}
