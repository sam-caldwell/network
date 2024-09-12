package core

import (
	"errors"
	"golang.org/x/sys/unix"
	"reflect"
	"testing"
	"unsafe"
)

type NrDummyStruct struct {
	lenData  int
	byteData []byte
	errData  error
}

func (nr NrDummyStruct) Len() int {
	return 0
}
func (nr NrDummyStruct) Serialize() ([]byte, error) {
	return []byte{}, nil
}

func TestNetlinkRequest(t *testing.T) {
	t.Run("Test the NetlinkRequest structure", func(t *testing.T) {
		t.Run("size check", func(t *testing.T) {
			// Expected size of the NetlinkRequest struct
			const expectedSizeInBytes = unsafe.Sizeof(unix.NlMsghdr{}) +
				unsafe.Sizeof([]NetlinkRequestData{}) +
				unsafe.Sizeof([]byte{}) +
				unsafe.Sizeof(map[IpProtocol]*SocketHandle{})

			if actualSize := unsafe.Sizeof(NetlinkRequest{}); actualSize != expectedSizeInBytes {
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
	t.Run("Test the AddData() method", func(t *testing.T) {
		// Dummy data for testing
		dummyData1 := NrDummyStruct{
			lenData:  int(1),
			byteData: []byte{0x01},
			errData:  nil,
		}
		dummyData2 := NrDummyStruct{
			lenData:  int(2),
			byteData: []byte{0x02, 0x03},
			errData:  errors.New("fake error"),
		}

		// Create a new NetlinkRequest
		req := NetlinkRequest{}

		// Add first data item. Verify that Data slice now contains one element
		if req.AddData(dummyData1); len(req.Data) != 1 {
			t.Errorf("Expected Data to have 1 element, got %d", len(req.Data))
		}
		// Add second data item, Verify that Data slice now contains two elements
		if req.AddData(dummyData2); len(req.Data) != 2 {
			t.Errorf("Expected Data to have 2 elements, got %d", len(req.Data))
		}

		// Check that both elements were correctly added
		if !reflect.DeepEqual(req.Data[0], dummyData1) {
			t.Errorf("Expected first element to be %v, got %v", dummyData1, req.Data[0])
		}
		if !reflect.DeepEqual(req.Data[1], dummyData2) {
			t.Errorf("Expected first element to be %v, got %v", dummyData2, req.Data[1])
		}

	})

}
