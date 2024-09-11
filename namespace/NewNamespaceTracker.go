package namespace

// NewNamespaceTracker - Initializes a new in-memory namespace tracker
//
// The tracker is needed to ensure that each call to the Get() method or GetFromThread(pid,tid) method will return
// the same file handle. Linux network namespaces allow different processes or threads to have their own isolated
// network configurations (interfaces, routing, etc.). In Linux, a thread can be part of a different network namespace
// than its parent process.
//
// Related Linux kernel sources:
//   - `get_task_ns()` and `task_active_pid_ns()` are used to retrieve a process/thread's network namespace.
//     See: https://github.com/torvalds/linux/blob/master/kernel/pid.c
//   - Network namespace management in the kernel can be found at:
//     https://github.com/torvalds/linux/blob/master/net/core/net_namespace.c
func NewNamespaceTracker() *Tracker {

	return &Tracker{

		current: closedHandle,

		namespaces: make(TrackerMap),
	}

}
