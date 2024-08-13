//go:build !darwin && linux && !windows

package network

import (
	"testing"
)

func TestGetDefaultRouteInfo(t *testing.T) {
	t.Run("Test routes", func(t *testing.T) {
		// Get the expected RouteInfo for comparison.
		expectedV4, err := getExpectedDefaultRoute(false)
		if err != nil {
			t.Error(err)
		}
		expectedV6, err := getExpectedDefaultRoute(true)
		if err != nil {
			t.Error(err)
		}
		route4, route6, err := GetDefaultRouteInfo()
		if err != nil {
			t.Error(err)
		}
		if route4 != nil && !route4.Equal(expectedV4) {
			t.Errorf("expectedV4: %v, actual: %v", expectedV4, route4)
		}
		if route6 != nil && !route6.Equal(expectedV6) {
			t.Errorf("expected: %v, actual: %v", expectedV6, route6)
		}
	})
}
