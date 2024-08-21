//go:build linux

package core

import (
	"fmt"
)

// String - Return string value of BridgeVlanInfo
// ToDo: Should we format this a bit better?
func (bridge *BridgeVlanInfo) String() string {
	return fmt.Sprintf("%+v", *bridge)
}
