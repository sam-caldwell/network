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

	originalNamespace, err := namespace.Get()
	if err != nil {
		log.Fatal(err)
	}
	newNamespace, err := namespace.NewWithName(expectedNamespaceName)
	if err != nil {
		log.Fatal(err)
	}
	if originalNamespace.Equal(newNamespace) {
		log.Fatal("New ns failed")
	}
	if actualNamespace, err := namespace.GetFromName(expectedNamespaceName); err != nil {
		log.Fatal(err)
	} else {
		if !actualNamespace.Equal(newNamespace) {
			log.Fatal("either GetFromName() or NewWithName() failed")
		}
	}
}
