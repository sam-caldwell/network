//go:build linux

package core

import (
	"fmt"
)

func (b *BridgeVlanInfo) String() string {
	return fmt.Sprintf("%+v", *b)
}
