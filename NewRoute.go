package network

import "net"

// NewRoute - Create, configure and return a new RouteInfo object.
func NewRoute(iface string, network net.IPNet, gateway net.IP) *RouteInfo {
	return &RouteInfo{
		Interface: iface,
		Network:   network.IP,
		Gateway:   gateway,
		Netmask:   network.Mask,
	}
}
