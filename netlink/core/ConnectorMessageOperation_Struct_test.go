package core

import (
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestCnMsgOp(t *testing.T) {
	t.Run("ConnectorMessageOperation struct", func(t *testing.T) {
		t.Run("test ConnectorMessageOperation fields", func(t *testing.T) {
			_ = ConnectorMessageOperation{
				ConnectorMessage: ConnectorMessage{
					ID: CbID{
						Idx: uint32(0),
						Val: uint32(0),
					},
					Seq:    uint32(0),
					Ack:    uint32(0),
					Length: uint16(0),
					Flags:  uint16(0),
				},
				Operation: uint32(0),
			}
		})
		t.Run("size check", func(t *testing.T) {
			var o ConnectorMessageOperation
			t.Run("Size of ConnectorMessageOperation", func(t *testing.T) {
				if unsafe.Sizeof(o.ConnectorMessage) != unsafe.Sizeof(ConnectorMessage{}) {
					t.Fatalf("ConnectorMessage size verification")
				}
				if unsafe.Sizeof(o.Operation) != unsafe.Sizeof(uint32(0)) {
					t.Fatalf("Operation size verification")
				}
			})
		})
	})
	t.Run("Basic ConnectorMessageOperation creation", func(t *testing.T) {
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
	})
	t.Run("test .Len() method", func(t *testing.T) {
		var o ConnectorMessageOperation
		if o.Len() != int(unsafe.Sizeof(ConnectorMessageOperation{})) {
			t.Fatalf("ConnectorMessageOperation.Len() method size mismatch")
		}
	})
	t.Run("test .Serialize() method", func(t *testing.T) {
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
	})

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
