//go:build linux

package network

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/unix"
	"testing"
)

func TestCreateNlMsgHdrForNewRouteCreate(t *testing.T) {
	var expected bytes.Buffer

	// Setup expected output
	hdr := unix.NlMsghdr{
		Len:   uint32(unix.SizeofNlMsghdr + unix.SizeofRtMsg),
		Type:  unix.RTM_NEWROUTE,
		Flags: unix.NLM_F_CREATE | unix.NLM_F_REQUEST | unix.NLM_F_ACK,
		Seq:   1,
		Pid:   uint32(unix.Getpid()),
	}

	_ = binary.Write(&expected, binary.LittleEndian, hdr)

	// Generate actual output
	var actual bytes.Buffer
	if err := createNlMsgHdrForNewRouteCreate(&actual); err != nil {
		t.Fatalf("createNlMsgHdrForNewRouteCreate() error: %v", err)
	}

	// Compare expected and actual
	if !bytes.Equal(expected.Bytes(), actual.Bytes()) {
		t.Fatalf("Netlink message header mismatch.\n"+
			"Expected: %v\n"+
			"  Actual: %v", expected.Bytes(), actual.Bytes())
	}
}
