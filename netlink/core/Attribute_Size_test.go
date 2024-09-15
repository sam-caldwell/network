package core

import (
	"testing"
	"unsafe"
)

func TestAttributeSize(t *testing.T) {
	if sz := int(unsafe.Sizeof(Attribute{})); sz != AttributeSize {
		t.Fatalf("Attribute size is wrong.")
	}
}
