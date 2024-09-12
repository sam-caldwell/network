package main

import (
	"fmt"
	"github.com/sam-caldwell/network/namespace"
	"github.com/sam-caldwell/network/netlink/core"
	"log"
	"os"
)

// Main entry point for the test runner inside the Docker container
func main() {
	const closedFile = -1
	type TestData struct {
		new      namespace.Handle
		current  namespace.Handle
		protocol core.IpProtocol
	}
	var (
		err              error
		currentNamespace namespace.Handle
		newNamespace     namespace.Handle
	)
	if currentNamespace, err = namespace.New(); err != nil {
		log.Fatal(err)
	}
	if newNamespace, err = namespace.New(); err != nil {
		log.Fatal(err)
	}
	testData := []TestData{
		{new: closedFile, current: closedFile, protocol: core.IpProtoTCP},
		{new: newNamespace, current: currentNamespace, protocol: 1},
	}

	for _, v := range testData {
		socket, err := core.GetNetlinkSocketAt(v.new, v.current, v.protocol)
		if err != nil {
			if socket != nil {
				log.Println("Error: socket should be nil (err != nil)")
			}
			log.Fatal(err)
		}
		//ToDo: test the socket if no error occurs.
	}

	fmt.Println("All tests passed successfully!")
	os.Exit(0)
}
