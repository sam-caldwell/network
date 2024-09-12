package main

import (
	"github.com/sam-caldwell/network/namespace"
	"log"
	"runtime"
)

func main() {
	log.Println("starting TestCreateNewNamespace")
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	originalNamespace, err := namespace.Get()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("originalNamespace: %v", originalNamespace)

	newNamespace, err := namespace.New()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("newNamespace: %v", newNamespace)

	if originalNamespace.Equal(newNamespace) {
		log.Fatal("New ns failed")
	}
	log.Printf("namespace equality: true\n"+
		"originalNamespace: %v\n"+
		"     newNamespace: %v",
		originalNamespace, newNamespace)
}
