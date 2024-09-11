package namespace

import "path/filepath"

// GetFromName - return a handle to the named network namespace.
func GetFromName(name string) (Handle, error) {

	return GetFromPath(filepath.Join(bindMountPath, name))

}
