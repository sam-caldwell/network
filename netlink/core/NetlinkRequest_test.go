package core

import (
	"bytes"
	"errors"
	"github.com/sam-caldwell/network/test"
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
		// Expected size of the NetlinkRequest struct
		const expectedSizeInBytes = int(unsafe.Sizeof(unix.NlMsghdr{}) +
			unsafe.Sizeof([]NetlinkRequestData{}) +
			unsafe.Sizeof([]byte{}) +
			unsafe.Sizeof(map[IpProtocol]*SocketHandle{}))

		t.Run("size check", func(t *testing.T) {
			t.Run("verify SizeOfNetlinkRequest", func(t *testing.T) {
				if SizeOfNetlinkRequest != expectedSizeInBytes {
					t.Fatalf("SizeOfNetlinkRequest mismatch. Expected: %d, Actual: %d",
						expectedSizeInBytes, SizeOfNetlinkRequest)
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

	t.Run("Test AddRawData() method", func(t *testing.T) {
		// Create a new NetlinkRequest
		req := NetlinkRequest{}

		// First set of raw data to add
		rawData1 := []byte{0x01, 0x02, 0x03, 0x04}

		// Add first set of raw data
		req.AddRawData(rawData1)

		// Check that RawData contains the correct bytes after the first addition
		if !bytes.Equal(req.RawData, rawData1) {
			t.Errorf("Expected RawData to be %v, got %v", rawData1, req.RawData)
		}

		// Second set of raw data to add
		rawData2 := []byte{0x05, 0x06, 0x07, 0x08}

		// Add second set of raw data
		req.AddRawData(rawData2)

		// Expected concatenated result
		expectedRawData := append(rawData1, rawData2...)

		// Check that RawData contains the correct bytes after the second addition
		if !bytes.Equal(req.RawData, expectedRawData) {
			t.Errorf("Expected RawData to be %v, got %v", expectedRawData, req.RawData)
		}
	})

	t.Run("Test NetlinkRequest.ExecuteIter() method happy path", func(t *testing.T) {
		t.Skip("disabled")
		const testDockerImage = "network-test:latest"
		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatalf("container build error: %v", err)
			}
		})
		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestNetlinkRequestExecuteFuncHappy"); err != nil {
				t.Fatalf("test container failed: %v", err)
			}
		})
	})
	t.Run("Test NetlinkRequest.ExecuteIter() method error scenario", func(t *testing.T) {
		t.Skip("disabled")
		const testDockerImage = "network-test:latest"
		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatalf("container build error: %v", err)
			}
		})
		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestNetlinkRequestExecuteFuncErrorScenario"); err != nil {
				t.Fatalf("test container failed: %v", err)
			}
		})
	})
}
