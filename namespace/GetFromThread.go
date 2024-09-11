//go:build linux

package namespace

import "fmt"

// GetFromThread - return the network namespace handle for the given process id (pid) and thread id (tid)
func GetFromThread(pid, tid int) (Handle, error) {
	return GetFromPath(fmt.Sprintf("/proc/%d/task/%d/ns/net", pid, tid))
}
