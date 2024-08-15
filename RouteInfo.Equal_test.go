package network

import (
	"testing"
)

func TestRouteInfo_Equal(t *testing.T) {

	t.Run("test Route==nil expect inequality", func(t *testing.T) {
		var lhs RouteInfo
		var rhs *RouteInfo
		if lhs.Equal(rhs) {
			t.Errorf("lhs.Equal(rhs)==true, want false")
		}
	})
	t.Run("test nil==nil expect equality", func(t *testing.T) {
		var lhs *RouteInfo
		var rhs *RouteInfo
		if !lhs.Equal(rhs) {
			t.Errorf("lhs.Equal(rhs)==false, want true")
		}
	})
	t.Run("test empty lhs==rhs, expect equality", func(t *testing.T) {
		var lhs RouteInfo
		var rhs RouteInfo
		if !lhs.Equal(&rhs) {
			t.Errorf("lhs.Equal(&rhs)==false, want true")
		}
	})
	t.Run("test lhs == rhs with data, expect equality", func(t *testing.T) {
		networkInterface := "myInterface"
		network := StringToIP("0.0.0.0")
		gateway := StringToIP("10.10.10.10")
		netMask, _ := StringToIPMask("255.255.255.255")

		var lhs = RouteInfo{
			Interface: networkInterface,
			Network:   network,
			Gateway:   gateway,
			Netmask:   netMask,
		}
		var rhs = RouteInfo{
			Interface: networkInterface,
			Network:   network,
			Gateway:   gateway,
			Netmask:   netMask,
		}
		if !lhs.Equal(&rhs) {
			t.Errorf("lhs.Equal(&rhs)==false, want true")
		}
	})
	t.Run("test lhs != rhs with data, expect inequality (interface)", func(t *testing.T) {
		networkInterfaceL := "lhs"
		networkInterfaceR := "rhs"
		network := StringToIP("0.0.0.0")
		gateway := StringToIP("10.10.10.10")
		netMask, _ := StringToIPMask("255.255.255.255")

		var lhs = RouteInfo{
			Interface: networkInterfaceL,
			Network:   network,
			Gateway:   gateway,
			Netmask:   netMask,
		}
		var rhs = RouteInfo{
			Interface: networkInterfaceR,
			Network:   network,
			Gateway:   gateway,
			Netmask:   netMask,
		}
		if lhs.Equal(&rhs) {
			t.Errorf("lhs.Equal(&rhs)==true, want false")
		}
	})
	t.Run("test lhs != rhs with data, expect inequality (network)", func(t *testing.T) {
		networkInterfaceL := "myInterface"
		networkInterfaceR := "myInterface"
		networkL := StringToIP("1.1.1.1")
		networkR := StringToIP("0.0.0.0")
		gatewayL := StringToIP("10.10.10.10")
		gatewayR := StringToIP("10.10.10.10")
		netMaskL, _ := StringToIPMask("255.255.255.255")
		netMaskR, _ := StringToIPMask("255.255.255.255")

		var lhs = RouteInfo{
			Interface: networkInterfaceL,
			Network:   networkL,
			Gateway:   gatewayL,
			Netmask:   netMaskL,
		}
		var rhs = RouteInfo{
			Interface: networkInterfaceR,
			Network:   networkR,
			Gateway:   gatewayR,
			Netmask:   netMaskR,
		}
		if lhs.Equal(&rhs) {
			t.Errorf("lhs.Equal(&rhs)==true, want false")
		}
	})
	t.Run("test lhs != rhs with data, expect inequality (gateway)", func(t *testing.T) {
		networkInterfaceL := "myInterface"
		networkInterfaceR := "myInterface"
		networkL := StringToIP("0.0.0.0")
		networkR := StringToIP("0.0.0.0")
		gatewayL := StringToIP("10.10.10.10")
		gatewayR := StringToIP("11.11.11.11")
		netMaskL, _ := StringToIPMask("255.255.255.255")
		netMaskR, _ := StringToIPMask("255.255.255.255")

		var lhs = RouteInfo{
			Interface: networkInterfaceL,
			Network:   networkL,
			Gateway:   gatewayL,
			Netmask:   netMaskL,
		}
		var rhs = RouteInfo{
			Interface: networkInterfaceR,
			Network:   networkR,
			Gateway:   gatewayR,
			Netmask:   netMaskR,
		}
		if lhs.Equal(&rhs) {
			t.Errorf("lhs.Equal(&rhs)==true, want false")
		}
	})
	t.Run("test lhs != rhs with data, expect inequality (netmask)", func(t *testing.T) {
		networkInterfaceL := "myInterface"
		networkInterfaceR := "myInterface"
		networkL := StringToIP("0.0.0.0")
		networkR := StringToIP("0.0.0.0")
		gatewayL := StringToIP("10.10.10.10")
		gatewayR := StringToIP("10.10.10.10")
		netMaskL, _ := StringToIPMask("255.255.255.255")
		netMaskR, _ := StringToIPMask("0.0.0.0")

		var lhs = RouteInfo{
			Interface: networkInterfaceL,
			Network:   networkL,
			Gateway:   gatewayL,
			Netmask:   netMaskL,
		}
		var rhs = RouteInfo{
			Interface: networkInterfaceR,
			Network:   networkR,
			Gateway:   gatewayR,
			Netmask:   netMaskR,
		}
		if lhs.Equal(&rhs) {
			t.Errorf("lhs.Equal(&rhs)==true, want false")
		}
	})

}
