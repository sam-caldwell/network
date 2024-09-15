package core

import (
	"testing"
)

func TestBpfNetlinkPolicySize(t *testing.T) {
	const expectedSizeInBytes = 1
	if BpfNetlinkPolicyEnumSize != expectedSizeInBytes {
		t.Fatal("size is wrong")
	}
}
