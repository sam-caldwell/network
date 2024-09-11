package namespace

import (
	"errors"
	"fmt"
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
	sync.Mutex
	// current - The currently active namespace handle
	current Id
	// namespaces - Map of namespaces and their file descriptors
	namespaces TrackerMap
}

// Id - Namespace Identifier.  This is an abstract identifier used to represent a network namespace in
// memory, which is mapped to the underlying linux file descriptor.
type Id int64

// TrackerMap is a map of namespace Id to the Handle (file descriptor).
type TrackerMap map[Id]Handle

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

// CreateNamespace - Creates a new network namespace and adds it to the tracker
func (nt *Tracker) CreateNamespace() (Handle, error) {

	// Lock the OS thread to ensure thread consistency
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	nt.Lock()
	defer nt.Unlock()

	newHandle, err := NewHandle()
	return newHandle, err
}

// Get - Return the current namespace or error if closed or not found.
func (nt *Tracker) Get() (h Handle, err error) {
	nt.Lock()
	defer nt.Unlock()

	if nt.current == closedHandle {
		return closedHandle, errors.New("current namespace is closed")
	}
	return nt.getByNamespaceIdUnsafe(nt.current)
}

// GetByNamespaceId - Return the namespace Handle matching the given Id value.
func (nt *Tracker) GetByNamespaceId(id Id) (h Handle, err error) {
	nt.Lock()
	defer nt.Unlock()
	return nt.getByNamespaceIdUnsafe(id)
}

// getByNamespaceIdUnsafe - Return the namespace Handle matching the given Id value with no locking. This method is
// not safe, as it has no locking.  Use the exported GetByNamespaceId() instead unless locking is otherwise handled.
func (nt *Tracker) getByNamespaceIdUnsafe(id Id) (h Handle, err error) {
	if ns, found := nt.namespaces[nt.current]; found {
		return ns, nil
	}
	return closedHandle, errors.New("current namespace is closed")
}

// SwitchNamespace - Switches to the given namespace, tracking the current one for later restoration.
func (nt *Tracker) SwitchNamespace(ns Id) (func() error, error) {

	previousNsId := nt.current

	rollbackFunc := func() (err error) {
		nt.Lock()
		defer nt.Unlock()

		// Get the requested namespace handle
		if _, exists := nt.namespaces[previousNsId]; !exists {
			return fmt.Errorf("namespace '%d' not found", ns)
		}

		// Store the current namespace to allow restoration
		if previousNsId == closedHandle {
			return fmt.Errorf("current namespace is closed")
		}
		nt.current = previousNsId
		return nil
	}

	// Get the requested namespace handle
	if _, exists := nt.namespaces[ns]; !exists {
		return nil, fmt.Errorf("namespace '%d' not found", ns)
	}

	// Store the current namespace to allow restoration
	if nt.current == closedHandle {
		return nil, fmt.Errorf("current namespace is closed")
	}
	nt.current = ns

	return rollbackFunc, nil
}

// DeleteNamespace - Removes a namespace from the tracker and closes the handle
func (nt *Tracker) DeleteNamespace(ns Id) error {
	// Lock the OS thread to ensure thread consistency
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	nt.Lock()
	defer nt.Unlock()

	if _, exists := nt.namespaces[ns]; !exists {
		return fmt.Errorf("namespace '%d' not found", ns)
	}
	handle, exists := nt.namespaces[ns]
	if !exists {
		return fmt.Errorf("namespace '%d' not found", ns)
	}
	// Close the handle and remove it from the tracker
	if err := handle.Close(); err != nil {
		return fmt.Errorf("failed to close namespace '%d': %v", ns, err)
	}
	delete(nt.namespaces, ns)
	return nil
}
