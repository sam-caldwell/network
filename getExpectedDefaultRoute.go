package network

import (
	"fmt"
	"os/exec"
	"strings"
)

// getExpectedDefaultRoute - run "route -v4" and parses the output to obtain route info.
func getExpectedDefaultRoute() (*RouteInfo, error) {
	/*
	 *  $ route -4
	 *  Kernel IP routing table
	 *	Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
	 *	default         192.168.3.1     0.0.0.0         UG    600    0        0 wlp4s0
	 *	link-local      0.0.0.0         255.255.0.0     U     1000   0        0 virbr0
	 *	172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 docker0
	 *	192.168.3.0     0.0.0.0         255.255.255.0   U     600    0        0 wlp4s0
	 *	192.168.49.0    0.0.0.0         255.255.255.0   U     0      0        0 br-c7bddc1f0590
	 *	192.168.122.0   0.0.0.0         255.255.255.0   U     0      0        0 virbr0
	 */
	output, err := exec.Command("route", "-v4").Output()
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(string(output), "\n") {
		if strings.HasPrefix(line, "default") {
			fields := strings.Fields(line)
			if len(fields) < 8 {
				continue
			}
			if fields[0] == "default" {
				fields[0] = "0.0.0.0"
			}
			network := StringToIP(fields[0])
			gateway := StringToIP(fields[1])
			mask, err := StringToIPMask(fields[2])
			if err != nil {
				return nil, err
			}
			return &RouteInfo{
				Network:   network,
				Interface: fields[len(fields)-1], //It's the last field
				Gateway:   gateway,
				Netmask:   mask,
			}, nil
		}
	}

	return nil, fmt.Errorf("no default route found")
}
