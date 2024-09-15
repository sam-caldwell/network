package core

import (
	"golang.org/x/sys/unix"
	"sync/atomic"
	"testing"
)

func TestNewNetlinkRequest(t *testing.T) {

	t.Run("NewNetlinkRequest", func(t *testing.T) {
		// Save the original value of the nextSequenceNumber and restore it later to avoid test interference.
		originalSequenceNumber := nextSequenceNumber
		defer atomic.StoreUint32(&nextSequenceNumber, originalSequenceNumber)

		t.Run("Test NewNetlinkRequest basic fields", func(t *testing.T) {
			// Set a known value for the next sequence number for testing.
			atomic.StoreUint32(&nextSequenceNumber, 42)

			// Create a new NetlinkRequest.
			proto := 100            // Example protocol
			flags := unix.NLM_F_ACK // Example flags
			req := NewNetlinkRequest(proto, flags)

			// Verify that the message header fields are set correctly.
			if req.Len != uint32(NetlinkMessageHeaderSize) {
				t.Errorf("Expected Len: %d, got: %d", NetlinkMessageHeaderSize, req.Len)
			}
			if req.Type != uint16(proto) {
				t.Errorf("Expected Type: %d, got: %d", proto, req.Type)
			}
			expectedFlags := NlmFRequest | uint16(flags)
			if req.Flags != expectedFlags {
				t.Errorf("Expected Flags: %d, got: %d", expectedFlags, req.Flags)
			}
			expectedSeq := uint32(43) // nextSequenceNumber was set to 42, and it should be incremented.
			if req.Seq != expectedSeq {
				t.Errorf("Expected Seq: %d, got: %d", expectedSeq, req.Seq)
			}
		})

		t.Run("Test NewNetlinkRequest sequence number increment", func(t *testing.T) {
			// Reset the sequence number to a known value.
			atomic.StoreUint32(&nextSequenceNumber, 100)

			// Create two new NetlinkRequest instances to check sequence number increment.
			req1 := NewNetlinkRequest(1, 0)
			req2 := NewNetlinkRequest(1, 0)

			// Verify that the sequence numbers increment as expected.
			if req1.Seq != 101 {
				t.Errorf("Expected Seq for req1: %d, got: %d", 101, req1.Seq)
			}
			if req2.Seq != 102 {
				t.Errorf("Expected Seq for req2: %d, got: %d", 102, req2.Seq)
			}
		})
	})
}
