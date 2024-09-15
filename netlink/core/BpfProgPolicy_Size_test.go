package core

import (
	"testing"
)

func TestBpfProgPolicySize(t *testing.T) {
	const expectedSizeInBytes = 1
	if BpfProgPolicySize != expectedSizeInBytes {
		t.Fatal("size is wrong")
	}
}
