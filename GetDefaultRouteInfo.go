package network

// GetDefaultRouteInfo - Get the default network address, gateway address, subnet mask and any errors.
func GetDefaultRouteInfo() (route *RouteInfo, err error) {
	// First, check IPv4 routes
	if route, err = GetDefaultRouteIPv4(); err != nil {
		return nil, err
	}
	if route == nil {
		//If no IPv4 default route, check IPv6 routes
		if route, err = GetDefaultRouteIPv6(); err != nil {
			return nil, err
		}
	}
	return route, nil
}
