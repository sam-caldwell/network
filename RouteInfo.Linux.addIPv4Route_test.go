//go:build linux

package network

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestAddIPv4Route(t *testing.T) {
	t.Run("Build test container", func(t *testing.T) {
		var output []byte
		var err error
		commandStr := []string{
			"build", "--tag", "network-test:latest", ".",
		}
		output, err = exec.Command("docker", commandStr...).Output()
		t.Logf("Container build output:\n===\n%s\n===\n", string(output))
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Run containerized test", func(t *testing.T) {
		GetCurrentDirectory := func() string {
			dir, _ := os.Getwd()
			absPath, _ := filepath.Abs(dir)
			return absPath
		}
		var output []byte
		var err error
		dir := GetCurrentDirectory()
		commandStr := []string{
			"run", "--cap-add=NET_ADMIN", "-v", dir + ":/opt", "network-test:latest",
		}
		output, err = exec.Command("docker", commandStr...).Output()
		t.Logf("Container run output:\n===\n%s\n===\n", string(output))
		if err != nil {
			t.Fatalf("containerized test failed:\n"+
				"error: %v\n"+
				"command: %v", err, strings.Join(commandStr, " "))
		}
	})
}
