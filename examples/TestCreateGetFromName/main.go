package main

import (
	"github.com/sam-caldwell/network/namespace"
	"log"
	"runtime"
)

func main() {

	const namespaceName = "namespaceNew"

	var err error
	var originalNamespace, newNamespace, ns namespace.Handle

	log.Println("starting TestCreateGetFromName")

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	log.Printf("Get originalNamespace")
	if originalNamespace, err = namespace.Get(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Create a new namespace with name (%s)", namespaceName)
	if newNamespace, err = namespace.NewWithName(namespaceName); err != nil {
		log.Fatal(err)
	}

	log.Printf("Verifying namespaces.  New %d Original %d", newNamespace, originalNamespace)
	if originalNamespace.Equal(newNamespace) {
		log.Fatalf("New ns failed.  original: %d, new: %d", originalNamespace, newNamespace)
	}

	log.Printf("GetFromName()")
	if ns, err = namespace.GetFromName(namespaceName); err != nil {
		log.Fatalf("GetFromName() failed with error: %v", err)
	} else {
		log.Printf("GetFromName() returned no error")
		if ns.Equal(newNamespace) {
			log.Fatalf("Get failed to get the newNamespace\n"+
				"     ns: %d\n"+
				"  newns: %d", ns, newNamespace)
		}
	}
}
