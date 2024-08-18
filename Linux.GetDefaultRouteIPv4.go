//go:build linux

package network

import (
	"fmt"
	"os"
	"strings"
)

// GetDefaultRouteIPv4 fetches the default route information for IPv4.
func GetDefaultRouteIPv4() (*RouteInfo, error) {
	/*
			 *   $ cat /proc/net/route
			 *   Iface   Destination     Gateway         Flags   RefCnt  Use     Metric  Mask            MTU     Window  IRTT
			 *   wlp4s0  00000000        0103A8C0        0003    0       0       600     00000000        0       0       0
			 *   virbr0  0000FEA9        00000000        0001    0       0       1000    0000FFFF        0       0       0
			 *   docker0 000011AC        00000000        0001    0       0       0       0000FFFF        0       0       0
			 *   wlp4s0  0003A8C0        00000000        0001    0       0       600     00FFFFFF        0       0       0
			 *   br-c7bddc1f0590 0031A8C0        00000000        0001    0       0       0       00FFFFFF        0       0       0
			 *   virbr0  007AA8C0        00000000        0001    0       0       0       00FFFFFF        0       0       0
			 *
			 *   $ route
			 *   Kernel IP routing table
			 *   Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
			 *   default         192.168.3.1     0.0.0.0         UG    600    0        0 wlp4s0
			 *   link-local      0.0.0.0         255.255.0.0     U     1000   0        0 virbr0
			 *   172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 docker0
			 *   192.168.3.0     0.0.0.0         255.255.255.0   U     600    0        0 wlp4s0
			 *   192.168.49.0    0.0.0.0         255.255.255.0   U     0      0        0 br-c7bddc1f0590
		 	 *   192.168.122.0   0.0.0.0         255.255.255.0   U     0      0        0 virbr0
	*/

	const (
		routeFile      = "/proc/net/route"
		defaultGateway = "0.0.0.0"
	)

	raw, err := os.ReadFile(routeFile)
	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(raw), "\n") {
		fields := strings.Fields(line)
		if fields[0] == "Iface" || len(fields) < 11 {
			continue
		}

		if destination := hexToIPv4(fields[1]); destination == defaultGateway {
			network := StringToIP(destination)
			netMask, err := StringToIPMask(hexToIPv4(fields[1]))
			if err != nil {
				return nil, err
			}
			gateway := StringToIP(hexToIPv4(fields[2]))

			return &RouteInfo{
				Network:   network,
				Interface: fields[0],
				Gateway:   gateway,
				Netmask:   netMask,
			}, nil
		}
	}

	return nil, fmt.Errorf("no default IPv4 route found")

}
