package core

import (
	"testing"
)

func TestLwt_constants(t *testing.T) {
	if LwtBpfMaxHeadroom != 256 {
		t.Fatal("LwtBpfMaxHeadroom is wrong")
	}
}
