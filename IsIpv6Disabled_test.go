package network

import (
	"os"
	"strings"
	"testing"
)

func TestIsIPv6Disabled(t *testing.T) {
	const sourceFile = "/proc/sys/net/ipv6/conf/all/disable_ipv6"
	var data []byte
	var err error

	if data, err = os.ReadFile(sourceFile); err != nil {
		t.Fatalf("error reading /proc/sys/net/ipv6/conf/all/disable_ipv6 (%v)", err)
	}

	expected := strings.TrimSpace(string(data)) == "0" //expect disabled/not disabled based on system setting

	actual, err := IsIpv6Disabled()
	if err != nil {
		t.Fatalf("error checking IsIpv6Disabled (%v)", err)
	}

	if actual != expected {
		t.Fatalf("expected: %v, actual: %v (data:%v)", expected, actual, string(data))
	}
}
