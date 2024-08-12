package network

import (
	"fmt"
	"net"
)

// RouteInfo - Describe a single route
type RouteInfo struct {
	Interface string
	Network   net.IP
	Gateway   net.IP
	Netmask   net.IPMask
}

// ToString - Return the route to string (interface network gateway mask)
func (route *RouteInfo) ToString() string {
	return fmt.Sprintf("%s %s %s %s",
		route.Interface, route.Network.String(), route.Gateway.String(), route.Netmask.String())
}
