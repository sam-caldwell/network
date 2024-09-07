package core

import "net"

// SetSrcIP sets the source IP address for the pedit action.
// This function determines whether the IP is an IPv4 or IPv6 address and calls the appropriate
// function to set the source address (either `SetIPv4Src` or `SetIPv6Src`).
//
// Parameters:
//
//	ip - The source IP address to set (net.IP).
func (p *TcPedit) SetSrcIP(ip net.IP) {
	// Check if the IP address is an IPv4 address (it will return a non-nil value if true).
	if ip.To4() != nil {
		// If it's an IPv4 address, call SetIPv4Src to handle the IPv4 source.
		p.SetIPv4Src(ip)
	} else {
		// Otherwise, assume it's an IPv6 address and call SetIPv6Src to handle the IPv6 source.
		p.SetIPv6Src(ip)
	}
}
