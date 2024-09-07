package core

import (
	"testing"
)

func TestConstant_NfnetlinkV0(t *testing.T) {
	if NfnetlinkV0 != 0 {
		t.Fatal("NfnetlinkV0 should be 0")
	}
}
