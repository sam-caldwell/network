package main

import (
	"github.com/sam-caldwell/network/namespace"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	const namespaceName = "testnamespace"

	originalNamespace, err := namespace.Get()
	if err != nil {
		log.Fatal(err)
	}
	newNamespace, err := namespace.New()
	if err != nil {
		log.Fatal(err)
	}
	if originalNamespace.Equal(newNamespace) {
		log.Fatal("newNamespace create failed")
	}
	// Attempt to set the namespace
	if err := namespace.SetNamespace(originalNamespace, unix.CLONE_NEWNET); err != nil {
		log.Fatalf("Failed to set namespace: %v", err)
	}

	log.Println("Successfully set the namespace.")
}
