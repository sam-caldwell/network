package core

import (
	"encoding/binary"
	"testing"
)

// Assuming you have a Deserialize method for CnMsgOp that handles deserialization correctly.
func TestDeserializeCnMsgOp(t *testing.T) {
	// Prepare a byte slice that represents a serialized CnMsgOp
	// The byte slice contains a CnMsg struct followed by the Op field
	buf := make([]byte, 24)

	// Fill the byte slice with values corresponding to CnMsgOp
	// CnMsg: CbID (Idx: 1, Val: 2), Seq: 100, Ack: 200, Length: 16, Flags: 0
	binary.LittleEndian.PutUint32(buf[0:4], 1)     // CbID.Idx
	binary.LittleEndian.PutUint32(buf[4:8], 2)     // CbID.Val
	binary.LittleEndian.PutUint32(buf[8:12], 100)  // Seq
	binary.LittleEndian.PutUint32(buf[12:16], 200) // Ack
	binary.LittleEndian.PutUint16(buf[16:18], 16)  // Length
	binary.LittleEndian.PutUint16(buf[18:20], 0)   // Flags

	// Op field
	binary.LittleEndian.PutUint32(buf[20:24], 500) // Op

	// Call DeserializeCnMsgOp with the byte slice
	cnMsgOp, err := DeserializeCnMsgOp(buf)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the deserialized CnMsgOp struct matches expected values
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
	if cnMsgOp.Op != 500 {
		t.Errorf("Expected Op to be 500, but got %d", cnMsgOp.Op)
	}
}
