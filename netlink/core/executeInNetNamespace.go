//go:build linux

package core

import (
	"fmt"
	"github.com/sam-caldwell/network/namespace"
	"runtime"
)

// executeInNetNamespace - set execution of the code following this call to the network namespace newNamespace, then moves the
// thread back to currentNamespace if open, otherwise to the current netns at the time the function was invoked.
//
// In case of success, the caller is expected to execute the returned function at the end of the code that needs to
// be executed in the network namespace.
//
// Example Implementation:
//
//	func myJobAt(...) error {
//	     d, err := executeInNetNamespace(...)
//	     if err != nil {
//	       return err
//	     }
//	     defer d()
//	     /* code which needs to be executed in specific netns */
//	 }
func executeInNetNamespace(newNamespace, currentNamespace namespace.Handle) (func(), error) {
	var (
		err                error
		moveBackFunc       func(namespace.Handle) error
		closeNamespaceFunc func() error
		unlockThreadFunc   func()
	)
	restoreFunc := func() {
		// ensure proper ordering
		if moveBackFunc != nil {
			if err = moveBackFunc(currentNamespace); err != nil {
				return
			}
		}
		if closeNamespaceFunc != nil {
			if err = closeNamespaceFunc(); err != nil {
				return
			}
		}
		if unlockThreadFunc != nil {
			unlockThreadFunc()
		}
	}
	if newNamespace.IsOpen() {
		runtime.LockOSThread()
		unlockThreadFunc = runtime.UnlockOSThread
		if !currentNamespace.IsOpen() {
			if currentNamespace, err = namespace.Get(); err != nil {
				restoreFunc()
				return nil, fmt.Errorf("error creating netlink socket "+
					"(could not get current namespace): %v", err)
			}
			closeNamespaceFunc = currentNamespace.Close
		}
		if err := namespace.Set(newNamespace); err != nil {
			restoreFunc()
			return nil, fmt.Errorf("error creating netlink socket "+
				"(failed to set into network namespace %d ): %v", newNamespace, err)
		}
		moveBackFunc = namespace.Set
	}
	return restoreFunc, nil
}
