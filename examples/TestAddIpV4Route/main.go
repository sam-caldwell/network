package main

import (
	"fmt"
	"github.com/sam-caldwell/network"
	"log"
	"net"
	"os/exec"
	"strings"
)

// checkRouteExistence - checks if the route exists in the system route table.
func checkRouteExistence(route *network.RouteInfo) bool {
	log.Println("Checking route existence")
	output, err := exec.Command("route", "-4").Output()
	log.Println(string(output))
	cmd := exec.Command("ip", "route", "show", route.Network.String())
	if output, err = cmd.CombinedOutput(); err != nil {
		log.Println("Checking route existence (fail)")
		return false
	}
	result := strings.Contains(string(output), route.Gateway.String())
	log.Printf("Checking route existence (result: %v)\n", result)
	return result
}

func createInterface(iface, ip string) error {
	var output []byte
	var err error
	log.Println("Creating interface")
	if output, err = exec.Command("ip", "link", "add", "name", iface, "type", "dummy").Output(); err != nil {
		return fmt.Errorf("error creating device: %v", err)
	}
	log.Printf("...dummy device created %s", output)
	if output, err = exec.Command("ip", "addr", "add", ip, "dev", iface).Output(); err != nil {
		return err
	}
	log.Printf("...ip address added to interface %s", output)
	if output, err = exec.Command("ip", "link", "set", iface, "up").Output(); err != nil {
		return err
	}
	if output, err = exec.Command("ip", "link").Output(); err != nil {
		return err
	}
	log.Printf("interface up %s", output)
	return nil
}

func deleteInterface(iface string) error {
	var output []byte
	var err error
	if output, err = exec.Command("ip", "link", "delete", "dummy").Output(); err != nil {
		return err
	}
	log.Println(string(output))
	return nil
}

func main() {
	log.Println("addRoute starting")

	func() {
		//Test the happy path
		happyRoute := network.RouteInfo{
			Interface: "dummy",
			Network:   net.ParseIP("127.1.0.0"),
			Gateway:   net.ParseIP("127.1.0.254"),
			Netmask:   network.MaskToIPMask(net.ParseIP("255.255.0.0")),
		}

		defer func() {
			log.Println("cleanup (happy path)...")
			defer func() {
				if err := deleteInterface(happyRoute.Interface); err != nil {
					log.Printf("error deleting interface: %v", err)
					return
				}
			}()
			cmd := exec.Command("ip", "route", "delete", happyRoute.Network.String())
			if _, err := cmd.CombinedOutput(); err != nil {
				log.Printf("error deleting route: %v", err)
			}
			log.Println("cleanup (happy path)...done")
		}()

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

		if err := createInterface(happyRoute.Interface, happyRoute.Network.String()); err != nil {
			log.Printf("error creating interface: %v", err)
			return
		}

		if err := happyRoute.Add(); err != nil {
			log.Printf("route.Add(): %v\n", err)
			return
		}

		if !checkRouteExistence(&happyRoute) {
			log.Println("Route was not added successfully")
			return
		}
		log.Println("happy path ok")
	}()
}
