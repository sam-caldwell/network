package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
	"testing"
)

func TestDeserializeNetlinkMessageHeaderAndData(t *testing.T) {
	t.Run("input too short", func(t *testing.T) {
		buf := make([]byte, unix.NLMSG_HDRLEN-1) // Insufficient length
		_, _, _, err := DeserializeNetlinkMessageHeaderAndData(buf)

		if err == nil {
			t.Fatalf("Expected error due to short byte slice, but got none")
		}

		expectedError := "input too short to contain NlMsghdr"
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
		}
	})

	t.Run("invalid Netlink Message Header", func(t *testing.T) {
		buf := []byte{
			0x00, 0x01, 0x02, 0x03, // Len
			0x10, 0x11, // Type
			0x20, 0x21, // Flags
			0x30, 0x31, 0x32, 0x33, // Seq
			0x40, 0x41, 0x42, 0x43, // Pid
			0x50, // Extra byte (invalid)
		}

		header, remainingData, msgLength, err := DeserializeNetlinkMessageHeaderAndData(buf)

		if header != nil {
			t.Errorf("header should be nil, but got %v", header)
		}
		if remainingData != nil {
			t.Errorf("remainingData should be nil, but got %v", remainingData)
		}
		if msgLength != 0 {
			t.Errorf("msgLength should be 0, but got %d", msgLength)
		}
		if err == nil {
			t.Fatalf("expected error, but got none")
		}
		if !errors.Is(unix.EINVAL, err) {
			t.Errorf("Expected error '%v', but got '%v'", unix.EINVAL, err)
		}
	})

	t.Run("test with a valid netlink header", func(t *testing.T) {
		netlinkMsg := []byte{
			// Netlink Header (NlMsghdr)
			0x1c, 0x00, 0x00, 0x00, // Len: 28 (16-byte header + 12-byte payload)
			0x01, 0x00, // Type: 1 (NLMSG_NOOP, for example)
			0x01, 0x00, // Flags: 1 (NLM_F_REQUEST)
			0x78, 0x56, 0x34, 0x12, // Seq: 0x12345678
			0xef, 0xcd, 0xab, 0x89, // Pid: 0x89abcdef

			// Payload (data)
			0xaa, 0xbb, 0xcc, 0xdd, // Some example payload data
			0x11, 0x22, 0x33, 0x44,
			0x55, 0x66, 0x77, 0x88,
		}

		header, remainingData, msgLength, err := DeserializeNetlinkMessageHeaderAndData(netlinkMsg)
		if err != nil {
			t.Errorf("error should be nil, but got %v", err)
		}

		expectedLen := binary.LittleEndian.Uint32([]byte{0x1c, 0x00, 0x00, 0x00})
		if header.Len != expectedLen {
			t.Errorf("Len should be %d, but got %d", expectedLen, header.Len)
		}

		expectedRemainingData := []byte{
			0xaa, 0xbb, 0xcc, 0xdd, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		}
		if !bytes.Equal(expectedRemainingData, remainingData) {
			t.Errorf("remainingData should be %v, but got %v", expectedRemainingData, remainingData)
		}

		expectedMsgLength := nlmAlignOf(int(header.Len))
		if msgLength != expectedMsgLength {
			t.Fatalf("msgLength should be %d, but got %d", expectedMsgLength, msgLength)
		}
	})
}
