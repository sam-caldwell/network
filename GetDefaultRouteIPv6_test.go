//go:build !darwin && linux && !windows

package network

import (
	"testing"
)

func TestGetDefaultRouteIPv6(t *testing.T) {
	//Fetch route info from the command line
	expected, err := getExpectedDefaultRoute(true)
	if err != nil {
		t.Fatalf("getExpectedDefaultRoute(): failed to fetch route info from command\n"+
			"error: %v", err)
	}
	if expected == nil {
		t.Fatalf("getExpectedDefaultRoute(): failed to get default route info from command\n"+
			"expected: %v\n", expected)
	}

	// Call the function under test
	result, err := GetDefaultRouteIPv6()
	if err != nil {
		t.Fatalf("GetDefaultRouteIPv6() returned error: '%v'", err)
	}
	if result == nil {
		t.Fatalf("GetDefaultRouteIPv6() failed to get default route info.\n"+
			"result: %v\n", result)
	}

	// Compare the results
	if !expected.Equal(result) {
		t.Fatalf("GetDefaultRouteIPv4() mismatch\n"+
			"  actual = %v\n"+
			"expected = %v", result, expected)
	}
}
