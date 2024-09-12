package namespace

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"path"
)

// NewWithName - create new named network namespace and set it as the current namespace, returning its handle.
func NewWithName(name string) (Handle, error) {

	if _, err := os.Stat(bindMountPath); os.IsNotExist(err) {
		err = os.MkdirAll(bindMountPath, 0o755)
		if err != nil {
			return None(), err
		}
	}

	newNamespace, err := New()
	if err != nil {
		return None(), err
	}

	namedPath := path.Join(bindMountPath, name)

	f, err := os.OpenFile(namedPath, os.O_CREATE|os.O_EXCL, 0o444)
	if err != nil {
		cerr := newNamespace.Close()
		if cerr != nil {
			return closedHandle, err
		}
		return None(), err
	}
	err = f.Close()
	if err != nil {
		return closedHandle, err
	}

	nsPath := fmt.Sprintf("/proc/%d/task/%d/ns/net", os.Getpid(), unix.Gettid())
	err = unix.Mount(nsPath, namedPath, "bind", unix.MS_BIND, "")
	if err != nil {
		cerr := newNamespace.Close()
		if cerr != nil {
			return closedHandle, err
		}
		return None(), err
	}

	return newNamespace, nil

}
