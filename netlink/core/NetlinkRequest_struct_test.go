package core

import (
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestNetlinkRequest(t *testing.T) {
	t.Run("Test the NetlinkRequest structure", func(t *testing.T) {
		// Expected size of the NetlinkRequest struct
		const expectedSizeInBytes = int(unsafe.Sizeof(unix.NlMsghdr{}) +
			unsafe.Sizeof([]NetlinkRequestData{}) +
			unsafe.Sizeof([]byte{}) +
			unsafe.Sizeof(map[IpProtocol]*SocketHandle{}))

		t.Run("size check", func(t *testing.T) {
			t.Run("verify NetlinkRequestSize", func(t *testing.T) {
				if NetlinkRequestSize != expectedSizeInBytes {
					t.Fatalf("NetlinkRequestSize mismatch. Expected: %d, Actual: %d",
						expectedSizeInBytes, NetlinkRequestSize)
				}
			})
			if actualSize := int(unsafe.Sizeof(NetlinkRequest{})); actualSize != expectedSizeInBytes {
				t.Errorf("Expected size: %d bytes, but got: %d bytes", expectedSizeInBytes, actualSize)
			}
		})

		t.Run("field check", func(t *testing.T) {
			_ = NetlinkRequest{
				NlMsghdr: unix.NlMsghdr{
					Len:   uint32(0),
					Type:  uint16(0),
					Flags: uint16(0),
					Seq:   uint32(0),
					Pid:   uint32(0),
				},
				Data:    []NetlinkRequestData{},
				RawData: []byte{},
				Sockets: make(map[IpProtocol]*SocketHandle),
			}
		})
	})

}
