package main

import (
	"github.com/sam-caldwell/network/namespace"
	"log"
	"runtime"
)

func main() {
	log.Println("starting TestCreateGetSetNamespace")
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	const createdNamespaceName = "foo"

	originalNamespace, err := namespace.Get()
	if err != nil {
		log.Fatal(err)
	}

	newNamespace, err := namespace.NewWithName(createdNamespaceName)
	if err != nil {
		log.Fatal(err)
	}

	if originalNamespace.Equal(newNamespace) {
		log.Fatal("New ns failed")
	}

	createdNamespace, err := namespace.GetFromName(createdNamespaceName)
	if err != nil {
		log.Fatal(err)
	}

	if !createdNamespace.Equal(newNamespace) {
		log.Fatal("created namespace is equal to new namespace")
	}

	newNamespace, err = namespace.NewWithName("newer")
	if err != nil {
		log.Fatal(err)
	}

	if createdNamespace.Equal(newNamespace) {
		log.Fatal("created namespace is equal to new namespace")
	}

	if err := namespace.Set(originalNamespace); err != nil {
		log.Fatal("failed to set original namespace")
	}

	expectedNamespace, err := namespace.Get()
	if err != nil {
		log.Fatal(err)
	}

	if !originalNamespace.Equal(expectedNamespace) {
		log.Fatal("created namespace is equal to new namespace")
	}

}
