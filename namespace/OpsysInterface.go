package namespace

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// OpsysGet returns a handle to the current thread's network namespace.
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
func OpsysGet() (Handle, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// os.Getpid() retrieves the current process ID
	// unix.Gettid() retrieves the current thread ID (since threads in Linux are treated as lightweight processes).
	return OpsysGetFromThread(os.Getpid(), unix.Gettid())
}

// OpsysGetFromThread - return the network namespace handle for the given process id (pid) and thread id (tid)
func OpsysGetFromThread(pid, tid int) (Handle, error) {
	return OpsysGetFromPath(fmt.Sprintf("/proc/%d/task/%d/ns/net", pid, tid))
}

// OpsysGetFromPath - return a handle to the network namespace identified by the given path.
func OpsysGetFromPath(path string) (Handle, error) {
	fd, err := unix.Open(path, unix.O_RDONLY|unix.O_CLOEXEC, 0)
	return Handle(fd), err
}

// OpsysGetFromPid - return the network namespace handle for the given pid.
func OpsysGetFromPid(pid int) (Handle, error) {
	const nsPath = "/proc/%d/ns/net"
	return OpsysGetFromPath(fmt.Sprintf(nsPath, pid))
}

// OpsysGetFromName - return a handle to the named network namespace.
func OpsysGetFromName(name string) (Handle, error) {
	return OpsysGetFromPath(filepath.Join(bindMountPath, name))
}

// OpsysDeleteNamespace deletes a named network namespace
func OpsysDeleteNamespace(name string) error {
	namedPath := path.Join(bindMountPath, name)

	err := unix.Unmount(namedPath, unix.MNT_DETACH)
	if err != nil {
		return err
	}

	return os.Remove(namedPath)
}
