package namespace

import (
	"fmt"
	"golang.org/x/sys/unix"
	"runtime"
	"sync"
)

// Tracker - Tracks network namespaces in memory and synchronizes with the OS when necessary
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
type Tracker struct {
	//
	sync.Mutex
	current    Handle            // The currently active namespace handle
	namespaces map[string]Handle // Map of namespaces and their file descriptors
}

// CreateNamespace - Creates a new network namespace and adds it to the tracker
func (nt *Tracker) CreateNamespace(name string) (Handle, error) {

	// Lock the OS thread to ensure thread consistency
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	nt.Lock()
	defer nt.Unlock()

	// Create a new namespace using the unshare system call
	if err := unix.Unshare(unix.CLONE_NEWNET); err != nil {
		return closedHandle, fmt.Errorf("failed to create namespace: %v", err)
	}

	// Get the handle for the new namespace
	handle, err := Get()
	if err != nil {
		return closedHandle, fmt.Errorf("failed to get new namespace handle: %v", err)
	}

	// Store the handle in memory
	nt.namespaces[name] = handle
	return handle, nil
}

// SwitchNamespace - Switches to the given namespace, tracking the current one for later restoration
func (nt *Tracker) SwitchNamespace(name string) (func(), error) {
	nt.Lock()
	defer nt.Unlock()

	// Get the requested namespace handle
	handle, exists := nt.namespaces[name]
	if !exists {
		return nil, fmt.Errorf("namespace '%s' not found", name)
	}

	// Store the current namespace to allow restoration
	if nt.current == closedHandle {
		currentNamespace, err := Get()
		if err != nil {
			return nil, fmt.Errorf("failed to get current namespace: %v", err)
		}
		nt.current = currentNamespace
	}

	// Switch to the requested namespace
	if err := Set(handle); err != nil {
		return nil, fmt.Errorf("failed to switch to namespace '%s': %v", name, err)
	}

	// Return a function to restore the previous namespace
	return func() {
		nt.Lock()
		defer nt.Unlock()

		if nt.current != closedHandle {
			_ = Set(nt.current) // Restore the previous namespace
			nt.current = closedHandle
		}
	}, nil
}

// DeleteNamespace - Removes a namespace from the tracker and closes the handle
func (nt *Tracker) DeleteNamespace(name string) error {
	nt.Lock()
	defer nt.Unlock()

	handle, exists := nt.namespaces[name]
	if !exists {
		return fmt.Errorf("namespace '%s' not found", name)
	}

	// Close the handle and remove it from the tracker
	if err := handle.Close(); err != nil {
		return fmt.Errorf("failed to close namespace '%s': %v", name, err)
	}
	delete(nt.namespaces, name)
	return nil
}
