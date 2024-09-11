package namespace

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"path/filepath"
	"runtime"
)

// Opsys - a direct interface for querying the operating system for a network namespace.
var Opsys LinuxNamespaceInterface

// LinuxNamespaceInterface - an empty struct representing the operating system interface for working with
// network namespaces.
type LinuxNamespaceInterface struct{}

// Get returns a handle to the current thread's network namespace.
// This function retrieves the network namespace associated with the calling thread by getting the process ID (PID)
// and thread ID (TID). It uses the `GetFromThread` function, which operates based on the Linux kernel's handling
// of network namespaces at the thread level.
//
// Linux network namespaces allow different processes or threads to have their own isolated network configurations
// (interfaces, routing, etc.). In Linux, a thread can be part of a different network namespace than its parent process.
//
// The `Get()` function utilizes the `os.Getpid()` and `unix.Gettid()` system calls to get the current process and
// thread ID, respectively, which are then passed to `GetFromThread` to retrieve the current thread's network namespace.
//
// Related Linux kernel sources:
//   - `get_task_ns()` and `task_active_pid_ns()` are used to retrieve a process/thread's network namespace.
//     See: https://github.com/torvalds/linux/blob/master/kernel/pid.c
//   - Network namespace management in the kernel can be found at:
//     https://github.com/torvalds/linux/blob/master/net/core/net_namespace.c
//
// Returns:
//   - Handle: A handle to the current thread's network namespace.
//   - error: An error if the handle cannot be retrieved.
func (ns *LinuxNamespaceInterface) Get() (Handle, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// os.Getpid() retrieves the current process ID
	// unix.Gettid() retrieves the current thread ID (since threads in Linux are treated as lightweight processes).
	return ns.GetFromThread(os.Getpid(), unix.Gettid())
}

// GetFromThread - return the network namespace handle for the given process id (pid) and thread id (tid)
func (ns *LinuxNamespaceInterface) GetFromThread(pid, tid int) (Handle, error) {
	return ns.GetFromPath(fmt.Sprintf("/proc/%d/task/%d/ns/net", pid, tid))
}

// GetFromPath - return a handle to the network namespace identified by the given path.
func (ns *LinuxNamespaceInterface) GetFromPath(path string) (Handle, error) {
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_CLOEXEC, 0)
	return Handle(fd), err
}

// GetFromPid - return the network namespace handle for the given pid.
func (ns *LinuxNamespaceInterface) GetFromPid(pid int) (Handle, error) {
	const path = "/proc/%d/ns/net"
	return ns.GetFromPath(fmt.Sprintf(path, pid))
}

// GetFromName - return a handle to the named network namespace.
func (ns *LinuxNamespaceInterface) GetFromName(name string) (Handle, error) {
	return ns.GetFromPath(filepath.Join(bindMountPath, name))
}
