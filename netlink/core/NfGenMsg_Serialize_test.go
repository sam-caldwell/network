package core

import (
	"encoding/binary"
	"testing"
)

func TestNfGenMsg_Serialize(t *testing.T) {
	t.Run("Serialize NfGenMsg", func(t *testing.T) {
		// Define a sample NfGenMsg with known values
		nfGenMsg := NfGenMsg{
			NfgenFamily: 0x01,
			Version:     0x02,
			ResId:       0x1234,
		}

		// Expected serialized byte slice
		expected := make([]byte, NfGenMsgSize)
		expected[0] = nfGenMsg.NfgenFamily
		expected[1] = nfGenMsg.Version
		binary.BigEndian.PutUint16(expected[2:4], nfGenMsg.ResId)

		// Call Serialize method
		serialized, err := nfGenMsg.Serialize()
		if err != nil {
			t.Fatalf("Unexpected error during serialization: %v", err)
		}

		// Check if the serialized output matches the expected bytes
		if len(serialized) != len(expected) {
			t.Fatalf("Serialized output length mismatch. Expected: %d, Got: %d", len(expected), len(serialized))
		}
		for i := range expected {
			if serialized[i] != expected[i] {
				t.Errorf("Mismatch at byte %d. Expected: %x, Got: %x", i, expected[i], serialized[i])
			}
		}
	})
}
