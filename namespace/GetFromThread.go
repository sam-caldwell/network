package namespace

import (
	"fmt"
)

// GetFromThread - return the network namespace handle for the given process id (pid) and thread id (tid)
func GetFromThread(pid, tid int) (Handle, error) {

	return GetFromPath(fmt.Sprintf(threadNamespacePath, pid, tid))

}
