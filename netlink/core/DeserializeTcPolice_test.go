package core

import (
	"testing"
)

// TestDeserializeTcPolice tests the DeserializeTcPolice function.
func TestDeserializeTcPolice(t *testing.T) {
	// Prepare test data for TcPolice
	data := []byte{
		// TcPolice
		0x01, 0x00, 0x00, 0x00, // TcPolice: Index uint32
		0x02, 0x00, 0x00, 0x00, // TcPolice: Action int32
		0x64, 0x00, 0x00, 0x00, // TcPolice: Limit uint32
		0xC8, 0x00, 0x00, 0x00, // TcPolice: Burst uint32
		0xDC, 0x05, 0x00, 0x00, // TcPolice: Mtu uint32
		//
		// Rate:=TcRateSpec struct {
		//
		0x16,       //                TcRateSpec (Rate): CellLog   uint8
		0x17,       //                TcRateSpec (Rate): Linklayer uint8
		0x03, 0x00, //                TcRateSpec (Rate): Overhead  uint16
		0x04, 0x00, //                TcRateSpec (Rate): CellAlign int16
		0x05, 0x00, //                TcRateSpec (Rate): Mpu       uint16
		0x06, 0x00, 0x00, 0x00, //    TcRateSpec (Rate): Rate      uint32
		//
		// PeakRate:=TcRateSpec struct {
		//
		0x20,       //                TcRateSpec (Rate): CellLog   uint8
		0x21,       //                TcRateSpec (Rate): Linklayer uint8
		0x22, 0x00, //                TcRateSpec (Rate): Overhead  uint16
		0x23, 0x00, //                TcRateSpec (Rate): CellAlign int16
		0x24, 0x00, //                TcRateSpec (Rate): Mpu       uint16
		0x25, 0x00, 0x00, 0x00, //    TcRateSpec (Rate): Rate      uint32
		// TcPolice
		0x2c, 0x01, 0x00, 0x00, // TcPolice: Refcnt int32
		0xFF, 0x00, 0x00, 0x00, // TcPolice: Bindcnt int32
		0x00, 0xFF, 0x00, 0x00, // TcPolice: Capab uint32
	}

	// Call DeserializeTcPolice
	police, err := DeserializeTcPolice(data)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify the output
	if police.Index != 1 {
		t.Errorf("Expected Index 1, got 0x%02x", police.Index)
	}
	if police.Action != 2 {
		t.Errorf("Expected Action 2, got 0x%02x", police.Action)
	}
	if police.Limit != 100 {
		t.Errorf("Expected Limit 100, got 0x%02x", police.Limit)
	}
	if police.Burst != 200 {
		t.Errorf("Expected Burst 200, got 0x%02x", police.Burst)
	}
	if police.Mtu != 0x5DC {
		t.Errorf("Expected Mtu 0x5DC, got 0x%02x", police.Mtu)
	}
	if police.Rate.Overhead != 0x03 {
		t.Errorf("Expected Rate.Overhead 0x03, got 0x%02x", police.Rate.Overhead)
	}
	if police.Rate.CellAlign != 0x04 {
		t.Errorf("Expected Rate.CellAlign 0x04, got 0x%02x", police.Rate.CellAlign)
	}
	if police.PeakRate.Overhead != 0x22 {
		t.Errorf("Expected PeakRate.Overhead 0x22, got 0x%02x", police.PeakRate.Overhead)
	}
	if police.PeakRate.CellAlign != 0x23 {
		t.Errorf("Expected PeakRate.CellAlign 0x23, got 0x%02x", police.PeakRate.CellAlign)
	}
	if police.Refcnt != 300 {
		t.Errorf("Expected Refcnt 300, got %d", police.Refcnt)
	}
	if police.Bindcnt != 255 {
		t.Errorf("Expected Bindcnt 65535, got %d", police.Bindcnt)
	}
	if police.Capab != 65280 {
		t.Errorf("Expected Capab 65535, got %d", police.Capab)
	}
}
