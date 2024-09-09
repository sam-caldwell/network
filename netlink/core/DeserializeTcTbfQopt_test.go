package core

import (
	"encoding/binary"
	"testing"
)

// TestDeserializeTcTbfQopt tests the DeserializeTcTbfQopt function.
func TestDeserializeTcTbfQopt(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Prepare a byte slice for TcTbfQopt
		data := make([]byte, SizeOfTcTbfQopt)

		// Fill the byte slice with sample values for the first TcRateSpec (Rate)
		data[0] = 0x01                                      // CellLog
		data[1] = 0x02                                      // Linklayer
		binary.LittleEndian.PutUint16(data[2:], 0x0304)     // Overhead
		binary.LittleEndian.PutUint16(data[4:], 0x0506)     // CellAlign
		binary.LittleEndian.PutUint16(data[6:], 0x0708)     // Mpu
		binary.LittleEndian.PutUint32(data[8:], 0x090A0B0C) // Rate

		// Fill the byte slice with sample values for the second TcRateSpec (Peakrate)
		data[12] = 0x0D                                      // CellLog
		data[13] = 0x0E                                      // Linklayer
		binary.LittleEndian.PutUint16(data[14:], 0x0F10)     // Overhead
		binary.LittleEndian.PutUint16(data[16:], 0x1112)     // CellAlign
		binary.LittleEndian.PutUint16(data[18:], 0x1314)     // Mpu
		binary.LittleEndian.PutUint32(data[20:], 0x15161718) // Rate

		// Fill the byte slice with sample values for the rest of the fields (Limit, Buffer, Mtu)
		binary.LittleEndian.PutUint32(data[24:], 0x19202122) // Limit
		binary.LittleEndian.PutUint32(data[28:], 0x23242526) // Buffer
		binary.LittleEndian.PutUint32(data[32:], 0x2728292A) // Mtu

		// Call DeserializeTcTbfQopt
		result, err := DeserializeTcTbfQopt(data)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Verify the results for Rate (TcRateSpec)
		if result.Rate.CellLog != 0x01 {
			t.Errorf("Expected Rate.CellLog 0x01, got 0x%02x", result.Rate.CellLog)
		}
		if result.Rate.Linklayer != 0x02 {
			t.Errorf("Expected Rate.Linklayer 0x02, got 0x%02x", result.Rate.Linklayer)
		}
		if result.Rate.Overhead != 0x0304 {
			t.Errorf("Expected Rate.Overhead 0x0304, got 0x%04x", result.Rate.Overhead)
		}
		if result.Rate.CellAlign != 0x0506 {
			t.Errorf("Expected Rate.CellAlign 0x0506, got 0x%04x", result.Rate.CellAlign)
		}
		if result.Rate.Mpu != 0x0708 {
			t.Errorf("Expected Rate.Mpu 0x0708, got 0x%04x", result.Rate.Mpu)
		}
		if result.Rate.Rate != 0x090A0B0C {
			t.Errorf("Expected Rate.Rate 0x090A0B0C, got 0x%08x", result.Rate.Rate)
		}

		// Verify the results for Peakrate (TcRateSpec)
		if result.Peakrate.CellLog != 0x0D {
			t.Errorf("Expected Peakrate.CellLog 0x0D, got 0x%02x", result.Peakrate.CellLog)
		}
		if result.Peakrate.Linklayer != 0x0E {
			t.Errorf("Expected Peakrate.Linklayer 0x0E, got 0x%02x", result.Peakrate.Linklayer)
		}
		if result.Peakrate.Overhead != 0x0F10 {
			t.Errorf("Expected Peakrate.Overhead 0x0F10, got 0x%04x", result.Peakrate.Overhead)
		}
		if result.Peakrate.CellAlign != 0x1112 {
			t.Errorf("Expected Peakrate.CellAlign 0x1112, got 0x%04x", result.Peakrate.CellAlign)
		}
		if result.Peakrate.Mpu != 0x1314 {
			t.Errorf("Expected Peakrate.Mpu 0x1314, got 0x%04x", result.Peakrate.Mpu)
		}
		if result.Peakrate.Rate != 0x15161718 {
			t.Errorf("Expected Peakrate.Rate 0x15161718, got 0x%08x", result.Peakrate.Rate)
		}

		// Verify the remaining fields
		if result.Limit != 0x19202122 {
			t.Errorf("Expected Limit 0x19202122, got 0x%08x", result.Limit)
		}
		if result.Buffer != 0x23242526 {
			t.Errorf("Expected Buffer 0x23242526, got 0x%08x", result.Buffer)
		}
		if result.Mtu != 0x2728292A {
			t.Errorf("Expected Mtu 0x2728292A, got 0x%08x", result.Mtu)
		}
	})
	t.Run("tests DeserializeTcTbfQopt with short input.", func(t *testing.T) {
		// Prepare short data that is less than the expected size
		shortData := make([]byte, SizeOfTcTbfQopt-1)

		// Call DeserializeTcTbfQopt
		_, err := DeserializeTcTbfQopt(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		expectedErr := "DeserializeTcTbfQopt: insufficient byte slice size"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
}
