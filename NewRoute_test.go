package network

import (
	"net"
	"testing"
)

func TestNewRoute(t *testing.T) {
	iface := "tap0"
	network := net.IPNet{
		IP:   net.ParseIP("10.0.0.0"),
		Mask: net.CIDRMask(24, 32),
	}
	gateway := net.ParseIP("10.0.0.1")

	route := NewRoute(iface, network, gateway)

	if !route.Gateway.Equal(gateway) {
		t.Fatalf("gateway mismatch")
	}
	if !route.Network.Equal(network.IP) {
		t.Fatalf("network mismatch")
	}
	if route.Netmask.String() != network.Mask.String() {
		t.Fatalf("netmask mismatch")
	}

}
