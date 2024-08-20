//go:build linux

package namespace

import "path/filepath"

// GetFromName gets a handle to a named network namespace such as one
// created by `ip netns add`.
func GetFromName(name string) (Handle, error) {
	return GetFromPath(filepath.Join(bindMountPath, name))
}
