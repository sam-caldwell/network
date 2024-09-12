package main

import (
	"github.com/sam-caldwell/network/namespace"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	const namespaceName = "testnamespace"

	log.Println("TestSetNamespace starting")

	log.Println("calling namespace.Get()")
	originalNamespace, err := namespace.Get()
	log.Println("return from namespace.Get()")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("originalNamespace: %v (from namespace.Get)", originalNamespace)

	log.Println("calling namespace.New()")
	newNamespace, err := namespace.New()
	log.Println("return from namespace.New()")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("    newNamespace: %v (from namespace.New)", newNamespace)

	if originalNamespace.Equal(newNamespace) {
		log.Fatalf("newNamespace create failed\n"+
			"originalNamespace: %v\n"+
			"     newNamespace: %v",
			originalNamespace, newNamespace)
	}

	// Attempt to set the namespace
	if err := namespace.SetNamespace(originalNamespace, unix.CLONE_NEWNET); err != nil {
		log.Fatalf("Failed to set namespace: %v", err)
	}

	log.Println("Successfully set the namespace.")
}
