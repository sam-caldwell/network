package core

import (
	"testing"
)

func TestDeserializeNetlinkMessageHeader(t *testing.T) {

	t.Run("valid input", func(t *testing.T) {
		//
		//	type NetlinkMessageHeader struct {
		//	   Len   uint32
		//	   Type  uint16
		//	   Flags uint16
		//	   Seq   uint32
		//	   Pid   uint32
		//	}
		testData := []byte{
			0x10, 0x00, 0x00, 0x00, // len
			0x18, 0x00, //   	  	   type
			0x01, 0x00, //     		   flags
			0x34, 0x12, 0x00, 0x00, // seq
			0x78, 0x56, 0x00, 0x00, // pid
		}

		// Call DeserializeNetlinkMessageHeader
		header, err := DeserializeNetlinkMessageHeader(testData)

		// Check for no errors
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check the deserialized header values
		if header.Len != 0x10 {
			t.Errorf("Expected Len to be %d, but got %d", 0x10, header.Len)
		}
		if header.Type != 0x18 {
			t.Errorf("Expected Type to be %d, but got %d", 0x18, header.Type)
		}
		if header.Flags != 0x01 {
			t.Errorf("Expected Flags to be 0x01, but got %d", header.Flags)
		}
		if header.Seq != 0x1234 {
			t.Errorf("Expected Seq to be 0x1234, but got %d", header.Seq)
		}
		if header.Pid != 0x5678 {
			t.Errorf("Expected PID to be 0x5678, but got %d", header.Pid)
		}
	})

	t.Run(ErrInputTooShort, func(t *testing.T) {
		// Prepare a byte slice with insufficient length (less than NLMSG_HDRLEN)
		buf := make([]byte, NetlinkMessageHeaderSize-1)

		// Call DeserializeNetlinkMessageHeader
		_, err := DeserializeNetlinkMessageHeader(buf)

		// Expect an error due to short byte slice
		if err == nil {
			t.Fatalf("Expected error due to short byte slice, but got none")
		}

		// Ensure the error message is correct
		expectedError := ErrInputTooShort
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
		}
	})
}
