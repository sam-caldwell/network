package network

// GetDefaultRouteInfo - Get the default network address, gateway address, subnet mask and any errors.
func GetDefaultRouteInfo() (v4, v6 *RouteInfo, err error) {
	// First, check IPv4 routes
	if v4, err = GetDefaultRouteIPv4(); err != nil {
		return nil, nil, err
	}

	// If no IPv4 default route, check IPv6 routes
	if v6, err = GetDefaultRouteIPv6(); err != nil {
		return nil, nil, err
	}

	return v4, v6, nil
}
