package main

import (
	"github.com/sam-caldwell/network"
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {
	log.Println("addRoute")

	// checkRouteExistence checks if the route exists in the system route table.
	checkRouteExistence := func(route *network.RouteInfo) bool {
		cmd := exec.Command("ip", "route", "show", route.Network.String())
		output, err := cmd.CombinedOutput()
		if err != nil {
			return false
		}
		return strings.Contains(string(output), route.Gateway.String())
	}

	func() {
		//Test the happy path
		happyRoute := network.RouteInfo{}
		defer func() {
			//Cleanup function
			cmd := exec.Command("ip", "route", "delete", happyRoute.Network.String())
			if _, err := cmd.CombinedOutput(); err != nil {
				log.Fatalf("error deleting route: %v", err)
			}
		}()

		happyRoute = network.RouteInfo{
			Interface: "lo",
			Network:   net.ParseIP("127.1.0.0"),
			Gateway:   net.ParseIP("127.0.0.254"),
			Netmask:   network.MaskToIPMask(net.ParseIP("255.255.0.0")),
		}
		if strings.TrimSpace(happyRoute.Interface) == "" {
			log.Fatalf("interface is empty")
		}
		if happyRoute.Network == nil {
			log.Fatalf("network is empty")
		}
		if happyRoute.Gateway == nil {
			log.Fatalf("gateway is empty")
		}
		if happyRoute.Netmask == nil {
			log.Fatalf("netmask is empty")
		}
		if err := happyRoute.Add(); err != nil {
			log.Printf("route.Add(): %v\n", err)
			return
		}
		if !checkRouteExistence(&happyRoute) {
			log.Println("Route was not added successfully")
			return
		}
	}()
	func() {
		//sadRoute:=network.RouteInfo{}
		defer func() {
			//cmd := exec.Command("ip", "route", "delete", sadRoute.Network.String())
			//if _, err := cmd.CombinedOutput(); err != nil {
			//	t.Fatalf("error deleting route: %v", err)
			//}
		}()
	}()
}
