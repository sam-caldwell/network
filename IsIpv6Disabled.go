//go:build !darwin && linux && !windows

package network

import (
	"fmt"
	"os"
	"strings"
)

// IsIpv6Disabled - Query the Operating system to see if IPv6 is disabled.
func IsIpv6Disabled() (bool, error) {
	const sourceFile = "/proc/sys/net/ipv6/conf/all/disable_ipv6"
	var data []byte
	var err error
	if data, err = os.ReadFile(sourceFile); err != nil {
		return false, fmt.Errorf("error reading /proc/sys/net/ipv6/conf/all/disable_ipv6 (%v)", err)
	}
	return strings.TrimSpace(string(data)) == "0", nil
}
