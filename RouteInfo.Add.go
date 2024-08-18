package network

import (
	"fmt"
)

// Add - Determine if RouteInfo is IPv4 or IPv6 and add the appropriate route to the host's route table
//
// This is a generic route info Add() method.  This will call the platform-specific addIPv4Route() or
// addIPv6Route() function as appropriate.
func (route *RouteInfo) Add() error {

	if route.Network.To4() != nil {
		return route.addIPv4Route()

	} else if route.Network.To16() != nil {
		//log.Println("create v6 route")
		return route.addIPv6Route()

	}

	return fmt.Errorf("unsupported IP version")

}
