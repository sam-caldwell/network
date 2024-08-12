package network

import "fmt"

// RouteInfo - Describe a single route
type RouteInfo struct {
	Interface string
	Network   string
	Gateway   string
	Netmask   string
}

// ToString - Return the route to string (interface network gateway mask)
func (route *RouteInfo) ToString() string {
	return fmt.Sprintf("%s %s %s %s", route.Interface, route.Network, route.Gateway, route.Netmask)
}
