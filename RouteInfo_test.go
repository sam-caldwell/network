package network

import (
	"fmt"
	"testing"
)

func TestRouteInfo(t *testing.T) {
	t.Run("verify struct", func(t *testing.T) {
		network, _ := StringToIPNet("0.0.0.0/0")
		gateway := StringToIP("192.168.13.14")
		subnetMask, _ := StringToIPMask("255.255.255.0")
		_ = RouteInfo{
			Interface: "interface0",
			Network:   network,
			Gateway:   gateway,
			Netmask:   subnetMask,
		}
	})
	t.Run("verify ToString() method", func(t *testing.T) {
		interface0 := "interface0"
		network, _ := StringToIPNet("0.0.0.0/0")
		gateway := StringToIP("192.168.13.14")
		subnetMask, _ := StringToIPMask("255.255.255.0")
		route := RouteInfo{
			Interface: interface0,
			Network:   network,
			Gateway:   gateway,
			Netmask:   subnetMask,
		}
		expected := fmt.Sprintf("%s %s %s %s", interface0, network.String(), gateway.String(), subnetMask.String())
		if actual := route.ToString(); actual != expected {
			t.Fatalf("route mismatch\n"+
				"  actual: %s\n"+
				"expected: %s", actual, expected)
		}
	})

}
