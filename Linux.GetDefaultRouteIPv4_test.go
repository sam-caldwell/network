//go:build !darwin && linux && !windows

package network

import (
	"testing"
)

func TestGetDefaultRouteIPv4(t *testing.T) {

	//Fetch route info from the command line
	expected, err := getExpectedDefaultRoute(false)
	if err != nil {
		t.Fatalf("failed to fetch route info from command\n"+
			"error: %v", err)
	}
	if expected == nil {
		t.Fatalf("failed to get default route info from command\n"+
			"expected: %v\n", expected)
	}

	// Call the function under test
	result, err := GetDefaultRouteIPv4()
	if err != nil {
		t.Fatalf("GetDefaultRouteIPv4() returned error: %v", err)
	}
	if result == nil {
		t.Fatalf("failed to get default route info.\n"+
			"result: %v\n", result)
	}

	// Compare the results
	if !expected.Equal(result) {
		t.Fatalf("GetDefaultRouteIPv4() mismatch\n"+
			"  actual = %v\n"+
			"expected = %v", result, expected)
	}
}
