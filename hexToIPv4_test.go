package network

import "testing"

func TestHexToIPv4(t *testing.T) {
	t.Run("test 192.168.1.0", func(t *testing.T) {
		hexIp := "0103A8C0"
		expected := "192.168.3.1"
		if actual := hexToIPv4(hexIp); actual != expected {
			t.Errorf("expected: %s, actual: %s", expected, actual)
		}
	})
	t.Run("test empty", func(t *testing.T) {
		hexIp := ""
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("expected panic")
			}
		}()
		_ = hexToIPv4(hexIp)
	})
}
