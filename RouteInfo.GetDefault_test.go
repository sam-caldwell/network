//go:build !darwin && linux && !windows

package network

import (
	"testing"
)

func TestRouteInfo_GetDefault(t *testing.T) {
	t.Run("Test routes", func(t *testing.T) {
		var expectedV4 *RouteInfo
		var expectedV6 *RouteInfo
		var route4 *RouteInfo
		var route6 *RouteInfo
		var err error
		// Get the expected RouteInfo for comparison.
		if expectedV4, err = getExpectedDefaultRoute(false); err != nil {
			t.Fatalf("getExpectedDefaultRoute():\n"+
				"expectedV4: %v\n"+
				"       err: %v", expectedV4, err)
		}
		if expectedV6, err = getExpectedDefaultRoute(true); err != nil {
			t.Fatalf("getExpectedDefaultRoute():\n"+
				"expectedV6: %v\n"+
				"       err: %v", expectedV6, err)
		}
		if route4, route6, err = GetDefaultRouteInfo(); err != nil {
			t.Fatalf("GetDefaultRouteInfo() failed\n"+
				"route4: %v\n"+
				"route6: %v\n"+
				" error: %v", route4, route6, err)
		}
		if route4 != nil && !route4.Equal(expectedV4) {
			t.Fatalf("GetDefaultRouteInfo()\n"+
				"expectedV4: %v,\n"+
				"    actual: %v", expectedV4, route4)
		}
		if route6 != nil && !route6.Equal(expectedV6) {
			t.Fatalf("GetDefaultRouteInfo()\n"+
				"expectedV6: %v,\n"+
				"    actual: %v", expectedV6, route6)
		}
	})
}
