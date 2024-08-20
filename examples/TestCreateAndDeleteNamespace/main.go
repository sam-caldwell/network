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

	newNamespace, err := namespace.New()
	if err != nil {
		log.Fatal(err)
	}

	if originalNamespace.Equal(newNamespace) {
		log.Fatal("New ns failed")
	}

}
