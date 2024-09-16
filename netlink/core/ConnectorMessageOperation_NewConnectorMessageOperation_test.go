package core

import (
	"encoding/binary"
	"testing"
)

func TestNewConnectorMessageOperation(t *testing.T) {
	idx := uint32(123) // Example index
	val := uint32(456) // Example value
	op := uint32(789)  // Example operation

	// Create a new ConnectorMessageOperation
	cm := NewConnectorMessageOperation(idx, val, op)

	// Verify the ID fields (Idx and Val)
	if cm.ID.Idx != idx {
		t.Errorf("Expected ID Idx: %d, got: %d", idx, cm.ID.Idx)
	}
	if cm.ID.Val != val {
		t.Errorf("Expected ID Val: %d, got: %d", val, cm.ID.Val)
	}

	// Verify Ack is 0
	if cm.Ack != 0 {
		t.Errorf("Expected Ack: 0, got: %d", cm.Ack)
	}

	// Verify Seq is 1
	if cm.Seq != 1 {
		t.Errorf("Expected Seq: 1, got: %d", cm.Seq)
	}

	// Verify Length is correctly set (size of the op)
	expectedLength := uint16(binary.Size(op))
	if cm.Length != expectedLength {
		t.Errorf("Expected Length: %d, got: %d", expectedLength, cm.Length)
	}

	// Verify the operation value
	if cm.Operation != op {
		t.Errorf("Expected Operation: %d, got: %d", op, cm.Operation)
	}
}
