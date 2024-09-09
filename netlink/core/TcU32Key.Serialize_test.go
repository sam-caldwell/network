package core

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func TestTcU32KeySerialize(t *testing.T) {
	// Create a sample TcU32Key object
	key := TcU32Key{
		Mask:    0x01020304,
		Val:     0x05060708,
		Off:     0x090A0B0C,
		OffMask: 0x0D0E0F10,
	}

	// Call the Serialize method
	serialized, err := key.Serialize()
	if err != nil {
		t.Fatalf("Unexpected error during serialization: %v", err)
	}

	// Expected serialized output
	expected := make([]byte, 16) // 4 fields, each 4 bytes long

	// Serialize expected values manually to ensure correctness
	binary.BigEndian.PutUint32(expected[0:4], key.Mask)
	binary.BigEndian.PutUint32(expected[4:8], key.Val)
	binary.LittleEndian.PutUint32(expected[8:12], uint32(key.Off))
	binary.LittleEndian.PutUint32(expected[12:16], uint32(key.OffMask))

	// Check if the serialized output matches the expected output
	if !reflect.DeepEqual(serialized, expected) {
		t.Fatalf("Serialized output doesn't match expected.\nExpected: %v\nGot: %v", expected, serialized)
	}
}
