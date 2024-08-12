package network

import (
	"testing"
)

func TestGetDefaultRouteInfo(t *testing.T) {
	t.Run("Test routes", func(t *testing.T) {
		// Get the expected RouteInfo for comparison.
		expected, err := getExpectedDefaultRoute()
		if err != nil {
			t.Error(err)
		}
		route, err := GetDefaultRouteInfo()
		if err != nil {
			t.Error(err)
		}
		if route != nil && !route.Equal(expected) {
			t.Errorf("expected: %v, actual: %v", expected, route)
		}
		if route != nil && !route.Equal(expected) {
			t.Errorf("expected: %v, actual: %v", expected, route)
		}
	})
}
