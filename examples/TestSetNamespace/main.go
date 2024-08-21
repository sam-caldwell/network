package namespace

import (
	"github.com/sam-caldwell/network/namespace"
	"golang.org/x/sys/unix"
	"log"
	"os"
)

func main() {
	const namespaceName = "testnamespace"
	createTestNamespace := func() namespace.Handle {
		f, err := os.CreateTemp("", namespaceName)
		if err != nil {
			log.Fatalf("Failed to create temp namespace file: %v", err)
		}
		defer func() { _ = f.Close() }()
		// Create the namespace
		fd := f.Fd()
		return namespace.Handle(fd)
	}

	// Test SetNamespace function
	ns := createTestNamespace()

	defer func() {
		// Clean up temporary namespace file
		if err := os.Remove(namespaceName); err != nil {
			log.Fatalf("Failed to remove test namespace: %v", err)
		}
	}()

	// Attempt to set the namespace
	if err := namespace.SetNamespace(ns, unix.CLONE_NEWNET); err != nil {
		log.Fatalf("Failed to set namespace: %v", err)
	}
}
