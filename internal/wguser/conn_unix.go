//go:build !windows
// +build !windows

package wguser

import (
	"net"
)

// dial is the default implementation of Client.dial.
func dial(device string) (net.Conn, error) { _ = "STUB: not implemented"; return *new(net.Conn), nil }

// find is the default implementation of Client.find.
func find() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// It seems that /var/run is a common location between Linux and the
// BSDs, even though it's a symlink on Linux.

// findUNIXSockets looks for UNIX socket files in the specified directories.
func findUNIXSockets(dirs []string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
