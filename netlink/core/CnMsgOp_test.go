package core

import (
	"testing"
	"unsafe"
)

func TestCnMsgOp(t *testing.T) {
	t.Run("CnMsgOp struct", func(t *testing.T) {
		t.Run("test CnMsgOp fields", func(t *testing.T) {
			_ = CnMsgOp{
				CnMsg: CnMsg{
					ID: CbID{
						Idx: uint32(0),
						Val: uint32(0),
					},
					Seq:    uint32(0),
					Ack:    uint32(0),
					Length: uint16(0),
					Flags:  uint16(0),
				},
				Op: uint32(0),
			}
		})
		t.Run("size check", func(t *testing.T) {
			var o CnMsgOp
			t.Run("Size of CnMsgOp", func(t *testing.T) {
				if unsafe.Sizeof(o.CnMsg) != unsafe.Sizeof(CnMsg{}) {
					t.Fatalf("CnMsg size verification")
				}
				if unsafe.Sizeof(o.Op) != unsafe.Sizeof(uint32(0)) {
					t.Fatalf("Op size verification")
				}
			})
		})
	})
	t.Run("test .Len() method", func(t *testing.T) {
		var o CnMsgOp
		if o.Len() != int(unsafe.Sizeof(CnMsgOp{})) {
			t.Fatalf("CnMsgOp.Len() method size mismatch")
		}
	})
	t.Run("test .Serialize() method", func(t *testing.T) {
		// Test data
		testMsg := CnMsgOp{
			CnMsg: CnMsg{
				ID: CbID{
					Idx: 1,
					Val: 2,
				},
				Seq:    3,
				Ack:    4,
				Length: 5,
				Flags:  6,
			},
			Op: 7,
		}

		// Expected byte slice (based on the values above)
		// We will generate it manually using the same approach as the Serialize function
		expected := make([]byte, SizeOfCnMsgOp)

		NativeEndian.PutUint32(expected[0:], testMsg.ID.Idx)
		NativeEndian.PutUint32(expected[4:], testMsg.ID.Val)
		NativeEndian.PutUint32(expected[8:], testMsg.Seq)
		NativeEndian.PutUint32(expected[12:], testMsg.Ack)
		NativeEndian.PutUint16(expected[16:], testMsg.Length)
		NativeEndian.PutUint16(expected[18:], testMsg.Flags)
		NativeEndian.PutUint32(expected[20:], testMsg.Op)

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
			shortData := testData[:10] // Shorter than SizeOfCnMsgOp
			var msg CnMsgOp
			err := msg.Deserialize(shortData)
			if err == nil {
				t.Fatalf("Expected an error due to insufficient data, but got none")
			}
			if err.Error() != "input too short" {
				t.Errorf("Expected 'data too short to deserialize CnMsgOp', got %v", err)
			}
		})
	})
}
