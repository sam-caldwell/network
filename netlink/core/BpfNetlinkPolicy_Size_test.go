package core

import (
	"testing"
)

func TestBpfNetlinkPolicySize(t *testing.T) {
	const expectedSizeInBytes = 1
	if BpfNetlinkPolicySize != expectedSizeInBytes {
		t.Fatal("size is wrong")
	}
}
