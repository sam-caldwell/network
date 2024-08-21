package core

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"golang.org/x/sys/unix"
	"testing"
)

func TestIfaCacheInfoDeserializeSerialize(t *testing.T) {

	deserializeIfaCacheInfoSafe := func(b []byte) *IfaCacheInfo {
		var msg = IfaCacheInfo{}
		if err := binary.Read(bytes.NewReader(b[0:unix.SizeofIfaCacheinfo]), NativeEndian(), &msg); err != nil {
			return nil
		}
		return &msg
	}

	var orig = make([]byte, unix.SizeofIfaCacheinfo)
	if _, err := rand.Read(orig); err != nil {
		t.Fatal(err)
	}
	safemsg := deserializeIfaCacheInfoSafe(orig)
	msg := DeserializeIfaCacheInfo(orig)
	testDeserializeSerialize(t, orig, safemsg, msg)
}
