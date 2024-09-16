package core

import (
	"testing"
	"unsafe"
)

func TestNetlinkMessageStructure(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		// Calculate the expected size of the NetlinkMessage
		expectedSize := unsafe.Sizeof(NetlinkMessageHeader{}) + unsafe.Sizeof([]byte{})
		actualSize := unsafe.Sizeof(NetlinkMessage{})

		if actualSize != expectedSize {
			t.Errorf("Expected size %d but got %d", expectedSize, actualSize)
		}
	})

	t.Run("TestStructure", func(t *testing.T) {
		// Create a NetlinkMessage with sample data
		msg := NetlinkMessage{
			Header: NetlinkMessageHeader{
				Len:   16,
				Type:  RtmNewLink,
				Flags: NlmFRequest,
				Seq:   12345,
				Pid:   54321,
			},
			Data: []byte{0x01, 0x02, 0x03, 0x04},
		}

		// Verify that the fields are set correctly
		if msg.Header.Len != 16 {
			t.Errorf("Expected Header.Len to be 16 but got %d", msg.Header.Len)
		}
		if msg.Header.Type != RtmNewLink {
			t.Errorf("Expected Header.Type to be RTM_NEWLINK but got %d", msg.Header.Type)
		}
		if msg.Header.Flags != NlmFRequest {
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
}
