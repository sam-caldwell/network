package network

// GetDefault - Get the default network address, gateway address, subnet mask and any errors.
//
// This is a generic route info GetDefault() method.  This will call the platform-specific GetDefaultRouteIPv6() or
// GetDefault() function as appropriate.
func (route *RouteInfo) GetDefault(useV6 bool) (err error) {

	var defaultRoute *RouteInfo

	if useV6 {
		if defaultRoute, err = GetDefaultRouteIPv6(); err != nil {
			return err
		}
	} else {
		if defaultRoute, err = GetDefaultRouteIPv4(); err != nil {
			return err
		}
	}

	*route = *defaultRoute

	return nil

}
