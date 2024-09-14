package core

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/unix"
	"testing"
)

func TestDeserializeNetlinkMessage(t *testing.T) {
	t.Run(ErrInputTooShort, func(t *testing.T) {
		buf := make([]byte, unix.NLMSG_HDRLEN-1) // Insufficient length
		_, _, _, err := DeserializeNetlinkMessage(buf)

		if err == nil {
			t.Fatalf("Expected error due to short byte slice, but got none")
		}

		expectedError := ErrInputTooShort
		if err.Error() != expectedError {
			t.Fatalf("Expected error message '%s', but got '%s'", expectedError, err.Error())
		}
	})

	t.Run("invalid Netlink Message Header", func(t *testing.T) {
		buf := []byte{
			0x01, 0x02, 0x03, 0x04, // Len
			0x11, 0x12, // Type
			0x21, 0x22, // Flags
			0x31, 0x32, 0x33, 0x34, // Seq
			0x41, 0x42, 0x43, 0x44, // Pid
			0x51, // Extra byte (invalid)
		}

		header, remainingData, msgLength, err := DeserializeNetlinkMessage(buf)

		if err == nil {
			t.Fatalf("expected error.  got none.")
		}
		if err.Error() != ErrInputTooLarge {
			t.Fatalf("Expected error message '%s', but got '%s'", ErrInputTooLarge, err.Error())
		}
		if header != nil {
			t.Fatalf("Expected header to be nil, got %v", header)
		}
		if remainingData != nil {
			t.Fatalf("Expected remaining data to be nil, got %v", remainingData)
		}
		if msgLength != 0 {
			t.Fatal("Expected msgLength to be 0, got ", msgLength)
		}
	})

	t.Run("test with a valid netlink header", func(t *testing.T) {
		t.Skip("disabled")
		netlinkMsg := []byte{
			// Netlink Header (NlMsghdr)
			0x1c, 0x1d, 0x1e, 0x1f, // Len: 28 (16-byte header + 12-byte payload)
			0x01, 0x00, // Type: 1 (NLMSG_NOOP, for example)
			0x01, 0x00, // Flags: 1 (NLM_F_REQUEST)
			0x78, 0x56, 0x34, 0x12, // Seq: 0x12345678
			0xef, 0xcd, 0xab, 0x89, // Pid: 0x89abcdef

			// Payload (data)
			0xaa, 0xbb, 0xcc, 0xdd, // Some example payload data
			0x11, 0x22, 0x33, 0x44,
			0x55, 0x66, 0x77, 0x88,
		}

		header, remainingData, msgLength, err := DeserializeNetlinkMessage(netlinkMsg)
		if err != nil {
			t.Fatalf("error should be nil, but got %v", err)
		}

		expectedLen := binary.LittleEndian.Uint32([]byte{0x1c, 0x1d, 0x1e, 0x1f})
		if header.Len != expectedLen {
			t.Fatalf("Len should be %d, but got %d", expectedLen, header.Len)
		}

		expectedRemainingData := []byte{
			0xaa, 0xbb, 0xcc, 0xdd, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		}
		if !bytes.Equal(expectedRemainingData, remainingData) {
			t.Fatalf("remainingData should be %v, but got %v", expectedRemainingData, remainingData)
		}

		expectedMsgLength := nlmAlignOf(int(header.Len))
		if msgLength != expectedMsgLength {
			t.Fatalf("msgLength should be %d, but got %d", expectedMsgLength, msgLength)
		}
	})

}
