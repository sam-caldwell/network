package network

import "testing"

func TestGetExpectedRouteInfo(t *testing.T) {
	expected, err := getExpectedDefaultRoute()
	if err != nil {
		t.Error(err)
	}
	if expected.Network.String() != "0.0.0.0" {
		t.Fatalf("expected default gateway")
	}
	if expected.Gateway.String() == "" {
		t.Fatal("expected default gateway")
	}
	if expected.Interface == "" {
		t.Fatal("expected default interface")
	}
}
