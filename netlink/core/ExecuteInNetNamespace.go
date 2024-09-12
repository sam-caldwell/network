package core

import (
	"fmt"
	"github.com/sam-caldwell/network/namespace"
	"runtime"
)

// ExecuteInNetNamespace - set execution of the code following this call to the newNamespace.
// Once the code is executed, it moves the thread back to the originalNamespace.
//
// In case of success, the caller is expected to execute the returned function at the end of the code that needs to
// be executed in the new network namespace.
//
// Example Usage:
//
//	func myJobAt(...) error {
//	     d, err := ExecuteInNetNamespace(...)
//	     if err != nil {
//	       return err
//	     }
//	     defer d()
//	     /* code which needs to be executed in specific netns */
//	 }
func ExecuteInNetNamespace(newNamespace, currentNamespace namespace.Handle) (func(), error) {
	var (
		err                error
		moveBackFunc       func(namespace.Handle) error
		closeNamespaceFunc func() error
		unlockThreadFunc   func()
	)

	// Function to restore the original namespace
	restoreFunc := func() {
		// Ensure proper ordering
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

	// Check if the newNamespace is open
	if newNamespace.IsOpen() {
		// Lock the OS thread to ensure the same thread is used throughout
		runtime.LockOSThread()
		unlockThreadFunc = runtime.UnlockOSThread

		// If currentNamespace is not open, get it
		if !currentNamespace.IsOpen() {
			if currentNamespace, err = namespace.Get(); err != nil {
				restoreFunc()
				return nil, fmt.Errorf("error getting current namespace: %v", err)
			}
			closeNamespaceFunc = currentNamespace.Close
		}

		// Switch to the new namespace
		if err := namespace.Set(newNamespace); err != nil {
			restoreFunc()
			return nil, fmt.Errorf("failed to set into network namespace %d: %v", newNamespace, err)
		}
		// Prepare move back function to return to the original namespace
		moveBackFunc = namespace.Set
	}

	// Return the function to restore the original namespace after execution
	return restoreFunc, nil
}
