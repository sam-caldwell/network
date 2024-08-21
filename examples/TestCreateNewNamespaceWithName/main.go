package main

import (
	"github.com/sam-caldwell/network/namespace"
	"log"
	"runtime"
)

func main() {
	const expectedNamespaceName = "foo"
	log.Println("starting TestCreateNewNamespaceWithName")
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	log.Println("getting originalNamespace")
	originalNamespace, err := namespace.Get()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("creating namespace (with name)")
	newNamespace, err := namespace.NewWithName(expectedNamespaceName)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Test equality to see if we failed to create a namespace")
	if originalNamespace.Equal(newNamespace) {
		log.Fatal("New ns failed")
	}

	log.Println("Calling GetFromName()")
	if actualNamespace, err := namespace.GetFromName(expectedNamespaceName); err != nil {
		log.Fatalf("GetFromName failed: %v", err)
	} else {
		if actualNamespace == -1 {
			log.Fatal("GetFromName returned -1 namespace")
		}
		log.Println("Calling GetFromName() to test outcome")
		if !actualNamespace.Equal(newNamespace) {
			log.Fatal("either GetFromName() or NewWithName() failed")
		}
	}
}
