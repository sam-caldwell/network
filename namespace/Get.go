//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
	"os"
	"runtime"
)

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
//
// Dependencies:
//
//	  +- Get() : returns file handle and error
//	     |
//	     +- os.Getpid()
//	     +- unix.Gettid()
//		    +- GetFromThread() : returns file handle and error
//	        |
//	        +- GetFromPath() : returns file handle and error from current/namespaces map
func Get() (Handle, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	return GetFromThread(
		// os.Getpid() retrieves the current process ID
		os.Getpid(),
		// unix.Gettid() retrieves the current thread ID (since threads in Linux are treated as lightweight processes).
		unix.Gettid(),
	)
}
