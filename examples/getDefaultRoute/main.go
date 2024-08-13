package main

import (
	"github.com/sam-caldwell/network"
	"log"
)

func main() {
	route4, route6, err := network.GetDefaultRouteInfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	if route4 == nil {
		log.Println("no ipv4 default route4 found")
	}
	log.Printf("            iface   network    gateway   mask")
	log.Printf("ipv4 route: %s", route4.ToString())
	if route6 == nil {
		log.Println("no ipv6 default route6 found")
	}
	log.Printf("ipv6 route: %s", route6.ToString())
}
