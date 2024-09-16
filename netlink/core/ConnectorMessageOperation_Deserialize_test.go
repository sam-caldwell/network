package core

import (
	"encoding/binary"
	"testing"
)

func TestConnectorMessageOperation_Deserialize(t *testing.T) {

	t.Run("test Deserialize() method", func(t *testing.T) {
		// Prepare test data with valid values
		testData := make([]byte, ConnectorMessageOperationSize)
		NativeEndian.PutUint32(testData[0:4], 0x01020304)   // ID.Idx
		NativeEndian.PutUint32(testData[4:8], 0x05060708)   // ID.Val
		NativeEndian.PutUint32(testData[8:12], 0x11121314)  // Seq
		NativeEndian.PutUint32(testData[12:16], 0x15161718) // Ack
		NativeEndian.PutUint16(testData[16:18], 0x1920)     // Length
		NativeEndian.PutUint16(testData[18:20], 0x2122)     // Flags
		NativeEndian.PutUint32(testData[20:24], 0x23242526) // Operation

		// Test valid deserialization
		t.Run("Valid Deserialization", func(t *testing.T) {
			var msg ConnectorMessageOperation
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
			if msg.Operation != 0x23242526 {
				t.Errorf("Expected Operation to be 0x23242526, got 0x%x", msg.Operation)
			}
		})

		// Test deserialization with insufficient data
		t.Run("Insufficient Data", func(t *testing.T) {
			shortData := testData[:10] // Shorter than ConnectorMessageOperationSize
			var msg ConnectorMessageOperation
			err := msg.Deserialize(shortData)
			if err == nil {
				t.Fatalf("Expected an error due to insufficient data, but got none")
			}
			if err.Error() != ErrInputTooShort {
				t.Errorf("Expected 'data too short to deserialize ConnectorMessageOperation', got %v", err)
			}
		})
	})
	t.Run("Test DeserializeCnMsgOp() function", func(t *testing.T) {
		// Prepare a byte slice that represents a serialized ConnectorMessageOperation
		// The byte slice contains a ConnectorMessage struct followed by the Operation field
		buf := make([]byte, 24)

		// Fill the byte slice with values corresponding to ConnectorMessageOperation
		// ConnectorMessage: CbID (Idx: 1, Val: 2), Seq: 100, Ack: 200, Length: 16, Flags: 0
		binary.LittleEndian.PutUint32(buf[0:4], 1)     // CbID.Idx
		binary.LittleEndian.PutUint32(buf[4:8], 2)     // CbID.Val
		binary.LittleEndian.PutUint32(buf[8:12], 100)  // Seq
		binary.LittleEndian.PutUint32(buf[12:16], 200) // Ack
		binary.LittleEndian.PutUint16(buf[16:18], 16)  // Length
		binary.LittleEndian.PutUint16(buf[18:20], 0)   // Flags

		// Operation field
		binary.LittleEndian.PutUint32(buf[20:24], 500) // Operation

		// Call DeserializeCnMsgOp with the byte slice
		cnMsgOp, err := DeserializeCnMsgOp(buf)
		if err != nil {
			t.Fatal(err)
		}

		// Check if the deserialized ConnectorMessageOperation struct matches expected values
		if cnMsgOp.ID.Idx != 1 {
			t.Errorf("Expected ID.Idx to be 1, but got %d", cnMsgOp.ID.Idx)
		}
		if cnMsgOp.ID.Val != 2 {
			t.Errorf("Expected ID.Val to be 2, but got %d", cnMsgOp.ID.Val)
		}
		if cnMsgOp.Seq != 100 {
			t.Errorf("Expected Seq to be 100, but got %d", cnMsgOp.Seq)
		}
		if cnMsgOp.Ack != 200 {
			t.Errorf("Expected Ack to be 200, but got %d", cnMsgOp.Ack)
		}
		if cnMsgOp.Length != 16 {
			t.Errorf("Expected Length to be 16, but got %d", cnMsgOp.Length)
		}
		if cnMsgOp.Flags != 0 {
			t.Errorf("Expected Flags to be 0, but got %d", cnMsgOp.Flags)
		}
		if cnMsgOp.Operation != 500 {
			t.Errorf("Expected Operation to be 500, but got %d", cnMsgOp.Operation)
		}
	})
}
