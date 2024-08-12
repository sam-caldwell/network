package network

import (
	"fmt"
	"os"
	"strings"
)

// GetDefaultRouteIPv6 fetches the default route information for IPv6.
func GetDefaultRouteIPv6() (*RouteInfo, error) {
	/*
	 *   $ route -6
	 *   Kernel IPv6 routing table
	 *   Destination                    Next Hop                   Flag Met Ref Use If
	 *   ip6-localhost/128              [::]                       U    256 1     0 lo
	 *   fe80::/64                      [::]                       U    1024 1     0 wlp4s0
	 *   [::]/0                         [::]                       !n   -1  1     0 lo
	 *   ip6-localhost/128              [::]                       Un   0   10     0 lo
	 *   fe80::1fa8:4cd:8210:5e87/128   [::]                       Un   0   2     0 wlp4s0
	 *   ip6-mcastprefix/8              [::]                       U    256 9     0 wlp4s0
	 *   [::]/0                         [::]                       !n   -1  1     0 lo
	 *
	 *   $ cat /proc/net/ipv6_route
	 *   00000000000000000000000000000001 80 00000000000000000000000000000000 00 00000000000000000000000000000000 00000100 00000001 00000000 00000001       lo
	 *   fe800000000000000000000000000000 40 00000000000000000000000000000000 00 00000000000000000000000000000000 00000400 00000001 00000000 00000001   wlp4s0
	 *   00000000000000000000000000000000 00 00000000000000000000000000000000 00 00000000000000000000000000000000 ffffffff 00000001 00000000 00200200       lo
	 *   00000000000000000000000000000001 80 00000000000000000000000000000000 00 00000000000000000000000000000000 00000000 0000000a 00000000 80200001       lo
	 *   fe800000000000001fa804cd82105e87 80 00000000000000000000000000000000 00 00000000000000000000000000000000 00000000 00000002 00000000 80200001   wlp4s0
	 *   ff000000000000000000000000000000 08 00000000000000000000000000000000 00 00000000000000000000000000000000 00000100 00000009 00000000 00000001   wlp4s0
	 *   00000000000000000000000000000000 00 00000000000000000000000000000000 00 00000000000000000000000000000000 ffffffff 00000001 00000000 00200200       lo
	 */
	const (
		routeFile      = "/proc/net/ipv6_route"
		defaultGateway = "::"
	)
	raw, err := os.ReadFile(routeFile)
	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(raw), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 10 {
			continue // Ignore invalid lines
		}
		// Check for the default route in IPv6 (all zeros for the destination)
		if destination := hexToIPv6(fields[0]); destination == defaultGateway {
			network := StringToIP(destination)
			netMask, err := StringToIPMask(hexToIPv6(fields[1]))
			if err != nil {
				return nil, err
			}
			gateway := StringToIP(hexToIPv6(fields[4]))

			return &RouteInfo{
				Network:   network,
				Interface: fields[9], // Interface name
				Gateway:   gateway,   // Gateway address
				Netmask:   netMask,   // Prefix length for IPv6, not the netmask
			}, nil
		}
	}
	return nil, fmt.Errorf("no default IPv6 route found")
}
