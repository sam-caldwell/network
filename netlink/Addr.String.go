package netlink

import (
	"fmt"
	"strings"
)

// String - returns CIDR IP Address and LABEL
func (addr Addr) String() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", addr.IPNet, addr.Label))
}
