Network/Namespace
=================

## Description
Provide an interface for handling network namespaces in pure golang.

## Usage
## Local Build and Test ##
## Usage
### Build/Test (local dev)
#### Install package
* `go get github.com/sam-caldwell/network/namespace`

#### Run tests
```shell
cd <network package root>
go test -v ./...
```

### Usage Examples
```go
package main

import (
	"github.com/sam-caldwell/network/namespace"
	"log"
	"net"
	"runtime"
)

func main() {
	// Lock our OS Thread to prevent accidental namespace switch...it would be bad.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	// Save the current network namespace
	origins, _ := namespace.Get()
	defer func() { _ = origins.Close() }()

	// Create a new network namespace
	newNamespace, _ := namespace.New()
	defer func() { _ = newNamespace.Close() }()

	// Do something with the network namespace
	netInterfaces, _ := net.Interfaces()
	log.Printf("Interfaces: %v\n", netInterfaces)

	// Switch back to the original namespace
	namespace.Set(origins)
}

```