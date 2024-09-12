package main

import (
	"github.com/sam-caldwell/network/namespace"
	"github.com/sam-caldwell/network/netlink/core"
	"log"
	"runtime"
)

func main() {
	log.Println("Test: Happy Path")
	var (
		err              error
		currentNamespace namespace.Handle
		newNamespace     namespace.Handle
		activeNs         namespace.Handle
		restoreFunc      func()
	)

	// Get the current namespace
	if currentNamespace, err = namespace.Get(); err != nil {
		log.Fatalf("Failed to get current namespace: %v", err)
	}
	defer func() { _ = currentNamespace.Close() }()
	log.Printf("currentNamespace: %v", currentNamespace)

	if newNamespace, err = namespace.New(); err != nil {
		log.Fatalf("Failed to create new namespace: %v", err)
	}
	defer func() { _ = newNamespace.Close() }()
	log.Printf("newNamespace: %v", newNamespace)

	log.Println("Lock the OS thread to ensure namespace changes affect the current thread")
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	log.Printf("Calling ExecuteInNetNamespace()\n"+
		"currentNamespace: %v\n"+
		"    newNamespace: %v",
		currentNamespace, newNamespace)

	// Execute the code in the new namespace and switch back
	if restoreFunc, err = core.ExecuteInNetNamespace(newNamespace, currentNamespace); err != nil {
		log.Fatalf("Expected no error, but got: %v", err)
	}

	if restoreFunc == nil {
		log.Fatalf("Expected a restore function but got nil")
	}

	// Reset activeNs from earlier to be sure of state
	activeNs = 0

	// Check if we're in the new namespace
	if activeNs, err = namespace.Get(); err != nil {
		log.Fatalf("Failed to get current active namespace: %v", err)
	}
	if !activeNs.Equal(newNamespace) {
		log.Fatalf("Expected to be in new namespace.\n"+
			"    activeNs: %v\n"+
			"newNamespace: %v", activeNs, newNamespace)
	}
	log.Printf("namespace mismatch.\n"+
		"    activeNs: %v\n"+
		"newNamespace: %v", activeNs, newNamespace)

	// Now move back to the original namespace
	restoreFunc()

	// Check if we've returned to the original namespace
	if activeNs, err = namespace.Get(); err != nil {
		log.Fatalf("Failed to get current active namespace: %v", err)
	}
	if activeNs != currentNamespace {
		log.Fatalf("Expected to return to original namespace, but we're in: %v", activeNs)
	}
}
