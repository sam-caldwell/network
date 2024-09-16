package core

import (
	"bytes"
	"testing"
)

// Unit test for DeserializeToList
func TestDeserializeToList(t *testing.T) {
	t.Run("test with empty input", func(t *testing.T) {
		messages, err := DeserializeToList([]byte{})
		if err != nil {
			if err.Error() != ErrInputTooShort {
				t.Fatalf("DeserializeToList returned error: %v", err)
			}
		}
		if len(messages) != 0 {
			t.Fatalf("Expected 0 messages, got %d", len(messages))
		}
	})

	t.Run("Malformed message: header length less than NetlinkMessageHeaderSize", func(t *testing.T) {
		testData := []byte{0x00, 0x00, 0x00, 0x00}
		if _, err := DeserializeToList(testData); err == nil {
			t.Fatalf("Expected error for malformed input, but got none")
		}
	})

	t.Run("test with a set of happy-path messages", func(t *testing.T) {
		actualData := []byte{
			// Message 0:
			// Netlink Header (NetlinkMessageHeader)
			0x1C, 0x00, 0x00, 0x00, // Len: 28
			0x01, 0x00, // Type: 1
			0x01, 0x00, // Flags: 1
			0x78, 0x56, 0x34, 0x12, // Seq: 0x12345678
			0xEF, 0xCD, 0xAB, 0x89, // Pid: 0x89ABCDEF
			// Payload (data)
			0xAA, 0xBB, 0xCC, 0xDD,
			0x11, 0x22, 0x33, 0x44,
			0x55, 0x66, 0x77, 0x88,

			// Message 1:
			// Netlink Header (NetlinkMessageHeader)
			0x1C, 0x00, 0x00, 0x00, // Len: 28
			0x01, 0x00, // Type: 1
			0x01, 0x00, // Flags: 1
			0x79, 0x56, 0x34, 0x12, // Seq: 0x12345679
			0xEF, 0xCD, 0xAB, 0x89, // Pid
			// Payload (data)
			0xAA, 0xBB, 0xCC, 0xDD,
			0x11, 0x22, 0x33, 0x44,
			0x55, 0x66, 0x77, 0x88,

			// Message 2:
			// Netlink Header (NetlinkMessageHeader)
			0x1C, 0x00, 0x00, 0x00, // Len: 28
			0x01, 0x00, // Type: 1
			0x01, 0x00, // Flags: 1
			0x80, 0x56, 0x34, 0x12, // Seq: 0x12345680
			0xEF, 0xCD, 0xAB, 0x89, // Pid
			// Payload (data)
			0xAA, 0xBB, 0xCC, 0xDD,
			0x11, 0x22, 0x33, 0x44,
			0x55, 0x66, 0x77, 0x88,
		}

		expected := []NetlinkMessage{
			{
				Header: NetlinkMessageHeader{
					Len:   28,
					Type:  0x01,
					Flags: 0x01,
					Seq:   0x12345678,
					Pid:   0x89ABCDEF,
				},
				Data: []byte{
					0xAA, 0xBB, 0xCC, 0xDD,
					0x11, 0x22, 0x33, 0x44,
					0x55, 0x66, 0x77, 0x88,
				},
			},
			{
				Header: NetlinkMessageHeader{
					Len:   28,
					Type:  0x01,
					Flags: 0x01,
					Seq:   0x12345679,
					Pid:   0x89ABCDEF,
				},
				Data: []byte{
					0xAA, 0xBB, 0xCC, 0xDD,
					0x11, 0x22, 0x33, 0x44,
					0x55, 0x66, 0x77, 0x88,
				},
			},
			{
				Header: NetlinkMessageHeader{
					Len:   28,
					Type:  0x01,
					Flags: 0x01,
					Seq:   0x12345680,
					Pid:   0x89ABCDEF,
				},
				Data: []byte{
					0xAA, 0xBB, 0xCC, 0xDD,
					0x11, 0x22, 0x33, 0x44,
					0x55, 0x66, 0x77, 0x88,
				},
			},
		}

		actualDataList, err := DeserializeToList(actualData)
		if err != nil {
			t.Fatalf("DeserializeToList returned error: %v", err)
		}

		if len(actualDataList) != len(expected) {
			t.Fatalf("Expected %d messages, got %d", len(expected), len(actualDataList))
		}

		for i, actual := range actualDataList {
			expectedMsg := expected[i]
			if actual.Header != expectedMsg.Header {
				t.Errorf("Message %d header mismatch.\nExpected: %+v\nGot:      %+v", i, expectedMsg.Header, actual.Header)
			}
			if !bytes.Equal(actual.Data, expectedMsg.Data) {
				t.Errorf("Message %d data mismatch.\nExpected: %v\nGot:      %v", i, expectedMsg.Data, actual.Data)
			}
		}
	})
}
