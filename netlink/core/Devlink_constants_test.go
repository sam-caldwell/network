//go:build linux

package core

import (
	"testing"
)

func TestDevlink_Constants(t *testing.T) {
	if DevlinkGenlName != "devlink" {
		t.Fatal("DevlinkGenlName wrong")
	}
	if DevlinkGenlVersion != 1 {
		t.Fatal("DevlinkGenlVersion wrong")
	}
	if DevlinkGenlMcgrpConfigName != "config" {
		t.Fatal("DevlinkGenlMcgrpConfigName wrong")
	}
}
