package core

import (
	"testing"
)

func TestInfoMsg_Size(t *testing.T) {

	const expectedSizeInBytes = 16

	if IfInfoMsgSize != expectedSizeInBytes {
		t.Errorf("IfInfoMsgSize != expectedSizeInBytes")
	}

}
