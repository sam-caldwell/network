//go:build linux

package namespace

import (
	"fmt"
)

// GetFromPid - return the network namespace handle for the given pid.
func GetFromPid(pid int) (Handle, error) {
	const nsPath = "/proc/%d/ns/net"
	return GetFromPath(fmt.Sprintf(nsPath, pid))
}
