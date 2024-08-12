package network

import (
	"net"
)

// RouteInfo - Describe a single route
type RouteInfo struct {
	Interface string
	Network   net.IP
	Gateway   net.IP
	Netmask   net.IPMask
}
