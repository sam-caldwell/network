package core

import "net"

// SetDstIP sets the destination IP address for the pedit action.
// This function determines whether the IP is an IPv4 or IPv6 address and calls the appropriate
// function to set the destination address (either `SetIPv4Dst` or `SetIPv6Dst`).
//
// Parameters:
//
//	ip - The destination IP address to set (net.IP).
func (p *TcPedit) SetDstIP(ip net.IP) {
	// Check if the IP address is an IPv4 address (it will return a non-nil value if true).
	if ip.To4() != nil {
		// If it's an IPv4 address, call SetIPv4Dst to handle the IPv4 destination.
		p.SetIPv4Dst(ip)
	} else {
		// Otherwise, assume it's an IPv6 address and call SetIPv6Dst to handle the IPv6 destination.
		p.SetIPv6Dst(ip)
	}
}
