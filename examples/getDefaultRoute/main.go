package main

import (
	"github.com/sam-caldwell/network"
	"log"
)

func main() {
	v4, v6, err := network.GetDefaultRouteInfo()
	if err != nil {
		log.Fatal(err)
		return
	}
	if v4 == nil {
		log.Println("no ipv4 default route found")
	}
	log.Printf("ipv4 route: %s", v4.ToString())
	if v6 == nil {
		log.Println("no ipv6 default route found")
	}
	log.Printf("ipv6 route: %s", v6.ToString())
}
