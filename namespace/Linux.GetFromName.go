//go:build linux

package namespace

import "path/filepath"

// GetFromName - return a handle to the named network namespace.
//
// Testing Note: `ip netns add` will create a namespace which we can fetch
// for testing.
func GetFromName(name string) (Handle, error) {
	return GetFromPath(filepath.Join(bindMountPath, name))
}
