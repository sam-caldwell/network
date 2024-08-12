package network

import (
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

	})

}
