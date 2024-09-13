package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestInfoMsg_NewIfInfomsgChild(t *testing.T) {
	// Create a parent RtAttr
	parent := &RtAttr{
		RtAttr: unix.RtAttr{
			Len:  0,
			Type: 0,
		},
		Data:     nil,
		children: []NetlinkRequestData{},
	}

	// Define InterfaceFamily
	family := InterfaceFamily(unix.AF_INET)

	// Call NewIfInfomsgChild
	child := NewIfInfomsgChild(parent, family)

	// Verify that child is not nil
	if child == nil {
		t.Fatal("Expected child not to be nil")
	}

	// Verify that child's Family field is set correctly
	if child.Family != uint8(family) {
		t.Errorf("Expected child Family to be %d, got %d", family, child.Family)
	}

	// Verify that child is appended to parent.children
	if len(parent.children) != 1 {
		t.Fatalf("Expected parent to have 1 child, got %d", len(parent.children))
	}

	// Verify that the child in parent.children is the same as the one returned
	if parent.children[0] != child {
		t.Error("Expected parent.children[0] to be the same as child")
	}
}
