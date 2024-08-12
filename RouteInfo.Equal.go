package network

// Equal - Compare two RouteInfo objects and return boolean true if equal
func (route *RouteInfo) Equal(rhs *RouteInfo) bool {
	if route == nil || rhs == nil {
		return route == nil && rhs == nil
	}

	return route.Interface == rhs.Interface &&
		route.Network.String() == rhs.Network.String() &&
		route.Netmask.String() == rhs.Netmask.String() &&
		route.Gateway.String() == rhs.Gateway.String() &&
		CompareIPMask(route.Netmask, rhs.Netmask)
}
