package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestNewNetlinkRequest(t *testing.T) {
	// Test parameters
	proto := 16
	flags := unix.NLM_F_ACK

	// Reset the sequence number to a known value for testing
	nextSequenceNumber = 0

	// Create a new netlink request
	req := NewNetlinkRequest(proto, flags)

	// Verify the fields of the netlink request
	if req.Type != uint16(proto) {
		t.Errorf("Expected Type %d, got %d", proto, req.Type)
	}

	expectedFlags := unix.NLM_F_REQUEST | uint16(flags)
	if req.Flags != expectedFlags {
		t.Errorf("Expected Flags %d, got %d", expectedFlags, req.Flags)
	}

	expectedLen := uint32(NetlinkMessageHeaderSize)
	if req.Len != expectedLen {
		t.Errorf("Expected Len %d, got %d", expectedLen, req.Len)
	}

	expectedSeq := uint32(1) // Since we reset nextSequenceNumber to 0, the first call will be 1
	if req.Seq != expectedSeq {
		t.Errorf("Expected Seq %d, got %d", expectedSeq, req.Seq)
	}

	if req.Pid != 0 {
		t.Errorf("Expected Pid 0, got %d", req.Pid)
	}
}
