package core

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"testing"
)

func (b *BridgeVlanInfo) write(b []byte) {
	native := NativeEndian()
	native.PutUint16(b[0:2], b.Flags)
	native.PutUint16(b[2:4], b.Vid)
}

func (b *BridgeVlanInfo) serializeSafe() []byte {
	length := SizeofBridgeVlanInfo
	b := make([]byte, length)
	b.write(b)
	return b
}

func deserializeBridgeVlanInfoSafe(b []byte) *BridgeVlanInfo {
	var msg = BridgeVlanInfo{}
	binary.Read(bytes.NewReader(b[0:SizeofBridgeVlanInfo]), NativeEndian(), &msg)
	return &msg
}

func TestBridgeVlanInfoDeserializeSerialize(t *testing.T) {
	var orig = make([]byte, SizeofBridgeVlanInfo)
	rand.Read(orig)
	safemsg := deserializeBridgeVlanInfoSafe(orig)
	msg := DeserializeBridgeVlanInfo(orig)
	testDeserializeSerialize(t, orig, safemsg, msg)
}
