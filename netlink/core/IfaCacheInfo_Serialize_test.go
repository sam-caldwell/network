package core

import (
	"bytes"
	"golang.org/x/sys/unix"
	"testing"
)

func TestIfaCacheInfo_Serialize(t *testing.T) {
	testInput := IfaCacheInfo{
		IfaCacheinfo: unix.IfaCacheinfo{
			Prefered: uint32(01),
			Valid:    uint32(02),
			Cstamp:   uint32(03),
			Tstamp:   uint32(04),
		},
	}
	expected := []byte{
		0x01, 0x00, 0x00, 0x00,
		0x02, 0x00, 0x00, 0x00,
		0x03, 0x00, 0x00, 0x00,
		0x04, 0x00, 0x00, 0x00,
	}
	data, err := testInput.Serialize()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, expected) {
		t.Fatalf("Serialize() method failed")
	}
}
