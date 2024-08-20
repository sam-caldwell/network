//go:build linux

package namespace

import "fmt"

// GetFromThread gets a handle to the network namespace of a given pid and tid.
func GetFromThread(pid, tid int) (Handle, error) {
	return GetFromPath(fmt.Sprintf("/proc/%d/task/%d/ns/net", pid, tid))
}
