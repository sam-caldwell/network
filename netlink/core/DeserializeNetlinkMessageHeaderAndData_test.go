package core

import (
	"errors"
	"golang.org/x/sys/unix"
	"testing"
)

func TestDeserializeNetlinkMessageHeaderAndData(t *testing.T) {
	t.Run("input too short", func(t *testing.T) {
		// Prepare a byte slice with insufficient length
		buf := make([]byte, unix.NLMSG_HDRLEN-1)
		_, _, _, err := DeserializeNetlinkMessageHeaderAndData(buf)
		// Expect an error due to short byte slice
		if err == nil {
			t.Fatalf("Expected error due to short byte slice, but got none")
		}
		// Ensure the error message is correct
		expectedError := "input too short to contain NlMsghdr"
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
		}
	})
	t.Run("invalid Netlink Message Header", func(t *testing.T) {
		var (
			header        *unix.NlMsghdr
			remainingData []byte
			msgLength     int
			err           error
		)

		t.Run("setup", func(t *testing.T) {
			/*
				 type NlMsghdr struct {
					Len   uint32  // 0x00 0x01 0x02 0x03
					Type  uint16  // 0x10 0x11
					Flags uint16  // 0x20 0x21
					Seq   uint32  // 0x30 0x31 0x32 0x33
					Pid   uint32  // 0x40 0x41 0x42 0x43
				}
			*/
			buf := []byte{
				0x00, 0x01, 0x02, 0x03,
				0x10, 0x11,
				0x20, 0x21,
				0x30, 0x031, 0x032, 0x33,
				0x40, 0x41, 0x42, 0x43,
				0x50,
			}
			header, remainingData, msgLength, err = DeserializeNetlinkMessageHeaderAndData(buf)
		})
		t.Run("verify no error state", func(t *testing.T) {
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
	})
}
