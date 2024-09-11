Network/Namespace
=================

## Description
Provide an interface for handling network namespaces in pure golang. Linux network namespaces allow different 
processes or threads to have their own isolated network configurations (interfaces, routing, etc.). In Linux, a 
thread can be part of a different network namespace than its parent process.

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
