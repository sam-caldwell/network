package network

import "fmt"

// ToString - Return the route to string (interface network gateway mask)
func (route *RouteInfo) ToString() string {
	return fmt.Sprintf("%s %s %s %s",
		route.Interface, route.Network.String(), route.Gateway.String(), route.Netmask.String())
}
