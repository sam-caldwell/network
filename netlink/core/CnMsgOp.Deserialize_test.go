package core

import (
	"testing"
)

func TestCnMsgOp_Deserialize(t *testing.T) {
	// Prepare test data with valid values
	testData := make([]byte, SizeOfCnMsgOp)
	NativeEndian.PutUint32(testData[0:4], 0x01020304)   // ID.Idx
	NativeEndian.PutUint32(testData[4:8], 0x05060708)   // ID.Val
	NativeEndian.PutUint32(testData[8:12], 0x11121314)  // Seq
	NativeEndian.PutUint32(testData[12:16], 0x15161718) // Ack
	NativeEndian.PutUint16(testData[16:18], 0x1920)     // Length
	NativeEndian.PutUint16(testData[18:20], 0x2122)     // Flags
	NativeEndian.PutUint32(testData[20:24], 0x23242526) // Op

	// Test valid deserialization
	t.Run("Valid Deserialization", func(t *testing.T) {
		var msg CnMsgOp
		err := msg.Deserialize(testData)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Check the deserialized values
		if msg.ID.Idx != 0x01020304 {
			t.Errorf("Expected ID.Idx to be 0x01020304, got 0x%x", msg.ID.Idx)
		}
		if msg.ID.Val != 0x05060708 {
			t.Errorf("Expected ID.Val to be 0x05060708, got 0x%x", msg.ID.Val)
		}
		if msg.Seq != 0x11121314 {
			t.Errorf("Expected Seq to be 0x11121314, got 0x%x", msg.Seq)
		}
		if msg.Ack != 0x15161718 {
			t.Errorf("Expected Ack to be 0x15161718, got 0x%x", msg.Ack)
		}
		if msg.Length != 0x1920 {
			t.Errorf("Expected Length to be 0x1920, got 0x%x", msg.Length)
		}
		if msg.Flags != 0x2122 {
			t.Errorf("Expected Flags to be 0x2122, got 0x%x", msg.Flags)
		}
		if msg.Op != 0x23242526 {
			t.Errorf("Expected Op to be 0x23242526, got 0x%x", msg.Op)
		}
	})

	// Test deserialization with insufficient data
	t.Run("Insufficient Data", func(t *testing.T) {
		shortData := testData[:10] // Shorter than SizeofCnMsgOp
		var msg CnMsgOp
		err := msg.Deserialize(shortData)
		if err == nil {
			t.Fatalf("Expected an error due to insufficient data, but got none")
		}
		if err.Error() != "data too short to deserialize CnMsgOp" {
			t.Errorf("Expected 'data too short to deserialize CnMsgOp', got %v", err)
		}
	})
}
