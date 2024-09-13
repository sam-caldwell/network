package core

import (
	"testing"
)

func TestGenericNetlinkMessageSize(t *testing.T) {
	const expectedSizeInBytes = 2
	if GenericNetlinkMessageSize != expectedSizeInBytes {
		t.Fatalf("GenericNetlinkMessageSize is wrong")
	}
}
