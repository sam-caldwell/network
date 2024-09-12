package core

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestNetlinkMessageStructure(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		// Calculate the expected size of the NetlinkMessage
		expectedSize := unsafe.Sizeof(unix.NlMsghdr{}) + unsafe.Sizeof([]byte{})
		actualSize := unsafe.Sizeof(NetlinkMessage{})

		if actualSize != expectedSize {
			t.Errorf("Expected size %d but got %d", expectedSize, actualSize)
		}
	})

	t.Run("TestStructure", func(t *testing.T) {
		// Create a NetlinkMessage with sample data
		msg := NetlinkMessage{
			Header: unix.NlMsghdr{
				Len:   16,
				Type:  unix.RTM_NEWLINK,
				Flags: unix.NLM_F_REQUEST,
				Seq:   12345,
				Pid:   54321,
			},
			Data: []byte{0x01, 0x02, 0x03, 0x04},
		}

		// Verify that the fields are set correctly
		if msg.Header.Len != 16 {
			t.Errorf("Expected Header.Len to be 16 but got %d", msg.Header.Len)
		}
		if msg.Header.Type != unix.RTM_NEWLINK {
			t.Errorf("Expected Header.Type to be RTM_NEWLINK but got %d", msg.Header.Type)
		}
		if msg.Header.Flags != unix.NLM_F_REQUEST {
			t.Errorf("Expected Header.Flags to be NLM_F_REQUEST but got %d", msg.Header.Flags)
		}
		if msg.Header.Seq != 12345 {
			t.Errorf("Expected Header.Seq to be 12345 but got %d", msg.Header.Seq)
		}
		if msg.Header.Pid != 54321 {
			t.Errorf("Expected Header.Pid to be 54321 but got %d", msg.Header.Pid)
		}

		// Verify the data slice
		expectedData := []byte{0x01, 0x02, 0x03, 0x04}
		if len(msg.Data) != len(expectedData) {
			t.Errorf("Expected Data length to be %d but got %d", len(expectedData), len(msg.Data))
		}
		for i := range msg.Data {
			if msg.Data[i] != expectedData[i] {
				t.Errorf("Expected Data[%d] to be %d but got %d", i, expectedData[i], msg.Data[i])
			}
		}
	})
	t.Run("Deserialize NetlinkMessage Header and data", func(t *testing.T) {
		t.Run(ErrInputTooShort, func(t *testing.T) {
			buf := make([]byte, unix.NLMSG_HDRLEN-1) // Insufficient length
			_, _, _, err := DeserializeNetlinkMessage(buf)

			if err == nil {
				t.Fatalf("Expected error due to short byte slice, but got none")
			}

			expectedError := "input too short"
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

			header, remainingData, msgLength, err := DeserializeNetlinkMessage(buf)

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

			header, remainingData, msgLength, err := DeserializeNetlinkMessage(netlinkMsg)
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

	})
}
