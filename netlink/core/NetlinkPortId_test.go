package core

import (
	"testing"
)

func TestNetlinkPortId(t *testing.T) {
	if NetlinkPortId != 0 {
		t.Fatal("NetlinkPortId is wrong")
	}
}
