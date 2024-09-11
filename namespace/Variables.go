package namespace

import (
	"sync"
)

// Track network namespaces in memory
//
// The tracker is needed to ensure that each call to its Get() method will return the same underlying file descriptor.
// Simply calling OpsysGet() will query the system and return a unique (new) file descriptor for the same file each
// time.  But this Tracker will only return the same file descriptor because the map tracks the underlying socket/file.
//

// Related Linux kernel sources:
//   - `get_task_ns()` and `task_active_pid_ns()` are used to retrieve a process/thread's network namespace.
//     See: https://github.com/torvalds/linux/blob/master/kernel/pid.c
//   - Network namespace management in the kernel can be found at:
//     https://github.com/torvalds/linux/blob/master/net/core/net_namespace.c
var (
	// lock - for variable for safety.
	lock sync.Mutex

	// current - the current (active) namespace
	current Handle

	// namespace - a map of all known namespaces.  This maps the path/filename to the file descriptor.
	namespaces map[string]Handle = make(map[string]Handle)
)
