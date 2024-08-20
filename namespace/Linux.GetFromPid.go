//go:build linux

package namespace

import "fmt"

// GetFromPid gets a handle to the network namespace of a given pid.
func GetFromPid(pid int) (Handle, error) {
	return GetFromPath(fmt.Sprintf("/proc/%d/ns/net", pid))
}
