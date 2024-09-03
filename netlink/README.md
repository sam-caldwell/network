Network/Netlink
===============

## Description
Provide a netlink implementation in golang to allow programs in user-space to communicate with
a Linux Kernel.  The ultimate goal of this package is Linux specific with an eye on being used
by the larger `network` package to create a uniform interface for the same functionality.

```text
     +---------------------+      +---------------------+
     | (3) application "A" |      | (3) application "B" |
     +------+--------------+      +--------------+------+
            |                                    |
            \                                    /
             \                                  /
              |                                |
      +-------+--------------------------------+-------+
      |        :                               :       |   user-space
 =====+        :   (5)  kernel socket API      :       +================
      |        :                               :       |   kernel-space
      +--------+-------------------------------+-------+
               |                               |
         +-----+-------------------------------+----+
         |        (1)  Netlink subsystem            |
         +---------------------+--------------------+
                               |
         +---------------------+--------------------+
         |       (2) Generic Netlink bus            |
         +--+--------------------------+-------+----+
            |                          |       |
    +-------+---------+                |       |
    |  (4) controller |               /         \
    +-----------------+              /           \
                                     |           |
                  +------------------+--+     +--+------------------+
                  | (3) kernel user "X" |     | (3) kernel user "Y" |
                  +---------------------+     +---------------------+
```

## Usage
### Build/Test (local dev)
#### Install package
* `go get github.com/sam-caldwell/network/netlink`
#### Run tests
```shell
cd <network package root>
go test -v ./...
```
#### Testing (requires root):
* `go test -v .` (this will run the tests as root in a container)
### Usage Examples
#### To add a new IP address to an interface:
```go
package main

import (
    "github.com/sam-caldwell/network/netlink"
)

func main() {
    lo, _ := netlink.LinkByName("eth0")
    addr, _ := netlink.ParseAddr("10.11.12.13/32")
    netlink.AddressAdd(lo, addr)
}

```
#### To add a bridge (bridging eth0):

> Use the `NewLinkAttrs()` constructor function rather than create a DIY object.
> This `NewLinkAttrs()` function will set all the default values for you, and they
> can be changed/updated over time.

```go
package main

import (
	"fmt"
	"github.com/sam-caldwell/network/netlink"
)

func main() {
	link := netlink.NewLinkAttrs()
	link.Name = "bridge0"
	bridge0 := &netlink.Bridge{LinkAttrs: link}
	if err := netlink.LinkAdd(bridge0); err != nil {
		fmt.Printf("could not add %s: %v\n", link.Name, err)
	}
	eth0, _ := netlink.LinkByName("eth0")
	netlink.LinkSetMaster(eth0, bridge0)
}
```
