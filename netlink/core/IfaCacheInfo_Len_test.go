package core

import (
	"testing"
)

func TestIfaCacheInfo_Len(t *testing.T) {
	const expectedSizeInBytes = IfaCacheInfoSize
	var o IfaCacheInfo
	if sz := o.Len(); sz != expectedSizeInBytes {
		t.Fatalf("Len() method failed")
	}
}
