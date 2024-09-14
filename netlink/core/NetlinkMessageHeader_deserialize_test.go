package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestDeserializeNetlinkMessageHeader(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		// Prepare a byte slice with valid input
		buf := make([]byte, NetlinkMessageHeaderSize)

		// Populate the byte slice with values for NlMsghdr (using little-endian encoding)
		NativeEndian.PutUint32(buf[0:4], unix.NLMSG_HDRLEN+10) // Len = NLMSG_HDRLEN + 10
		NativeEndian.PutUint16(buf[4:6], unix.RTM_NEWROUTE)    // Type
		NativeEndian.PutUint16(buf[6:8], 0x01)                 // Flags
		NativeEndian.PutUint32(buf[8:12], 1234)                // Seq
		NativeEndian.PutUint32(buf[12:16], 5678)               // PID

		// Call DeserializeNetlinkMessageHeader
		hdr, err := DeserializeNetlinkMessageHeader(buf)

		// Check for no errors
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check the deserialized header values
		if hdr.Len != unix.NLMSG_HDRLEN+10 {
			t.Errorf("Expected Len to be %d, but got %d", unix.NLMSG_HDRLEN+10, hdr.Len)
		}
		if hdr.Type != unix.RTM_NEWROUTE {
			t.Errorf("Expected Type to be %d, but got %d", unix.RTM_NEWROUTE, hdr.Type)
		}
		if hdr.Flags != 0x01 {
			t.Errorf("Expected Flags to be 0x01, but got %d", hdr.Flags)
		}
		if hdr.Seq != 1234 {
			t.Errorf("Expected Seq to be 1234, but got %d", hdr.Seq)
		}
		if hdr.Pid != 5678 {
			t.Errorf("Expected PID to be 5678, but got %d", hdr.Pid)
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
		expectedError := "byte slice is too short"
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
		}
	})
}
