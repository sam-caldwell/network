package main

import (
	"github.com/sam-caldwell/network/namespace"
	"github.com/sam-caldwell/network/netlink/core"
	"log"
	"runtime"
)

func main() {
	log.Println("Test: SAD PATH: error on setting invalid namespace")

	// Get the current namespace
	currentNamespace, err := namespace.Get()
	if err != nil {
		log.Fatalf("Failed to get current namespace: %v", err)
	}
	defer func() { _ = currentNamespace.Close() }()

	// Create a new network namespace
	newNamespace, err := namespace.New()
	if err != nil {
		log.Fatalf("Failed to create new namespace: %v", err)
	}
	defer func() { _ = newNamespace.Close() }()

	// Use a closed namespace to induce an error
	_ = newNamespace.Close()

	// Lock the OS thread to ensure namespace changes affect the current thread
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	log.Printf("Calling ExecuteInNetNamespace()\n"+
		"      newNamespace: %v\n"+
		"  currentNamespace: %v\n",
		newNamespace, currentNamespace)

	// Attempt to switch to a closed namespace, which should fail
	_, err = core.ExecuteInNetNamespace(newNamespace, currentNamespace)
	if err == nil {
		log.Fatal("Expected an error, but got nil")
	}

	// Check if we've returned to the original namespace
	activeNs, err := namespace.Get()
	if err != nil {
		log.Fatalf("Failed to get current active namespace: %v", err)
	}
	if activeNs != currentNamespace {
		log.Fatalf("Expected to stay in original namespace, but we're in: %v", activeNs)
	}
}
