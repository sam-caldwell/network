//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
)

// SetNamespace - set the namespace (using golang.org/x/sys/unix.Setns)
//
// We keep this wrapper around unix.Setns() because SetNamespace() is more readable
// and less likely to be confused with a few other things we are building (e.g. name server).
// ...and there's the golang love of changing things under us that make wrappers appealing.
//
// See https://man7.org/linux/man-pages/man2/setns.2.html
func SetNamespace(namespace Handle, namespaceType NsTypes) error {
	return unix.Setns(int(namespace), int(namespaceType))
}
