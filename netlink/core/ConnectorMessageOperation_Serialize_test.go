package core

import (
	"testing"
)

func TestConnectorMessageOperation_Serialize(t *testing.T) {
	// Test data
	testMsg := ConnectorMessageOperation{
		ConnectorMessage: ConnectorMessage{
			ID: CbID{
				Idx: 1,
				Val: 2,
			},
			Seq:    3,
			Ack:    4,
			Length: 5,
			Flags:  6,
		},
		Operation: 7,
	}

	// Expected byte slice (based on the values above)
	// We will generate it manually using the same approach as the Serialize function
	expected := make([]byte, ConnectorMessageOperationSize)

	NativeEndian.PutUint32(expected[0:], testMsg.ID.Idx)
	NativeEndian.PutUint32(expected[4:], testMsg.ID.Val)
	NativeEndian.PutUint32(expected[8:], testMsg.Seq)
	NativeEndian.PutUint32(expected[12:], testMsg.Ack)
	NativeEndian.PutUint16(expected[16:], testMsg.Length)
	NativeEndian.PutUint16(expected[18:], testMsg.Flags)
	NativeEndian.PutUint32(expected[20:], testMsg.Operation)

	// Call the Serialize function
	actual, err := testMsg.Serialize()

	// Verify there was no error during serialization
	if err != nil {
		t.Fatalf("unexpected error during serialization: %v", err)
	}

	// Check that the actual result matches the expected result
	if len(actual) != len(expected) {
		t.Fatalf("serialized data length mismatch: got %v, want %v", len(actual), len(expected))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("byte mismatch at index %d: got %v, want %v", i, actual[i], expected[i])
		}
	}
}
