//go:build windows
// +build windows

package wguser

import (
	"net"
)

// Expected prefixes when dealing with named pipes.
const (
	pipePrefix = `\\.\pipe\`
	wgPrefix   = `ProtectedPrefix\Administrators\WireGuard\`
)

// dial is the default implementation of Client.dial.
func dial(device string) (net.Conn, error) { _ = "STUB: not implemented"; return *new(net.Conn), nil }

// find is the default implementation of Client.find.
func find() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// findNamedPipes looks for Windows named pipes that match the specified
// search string prefix.
func findNamedPipes(search string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Thanks @zx2c4 for the tips on the appropriate Windows APIs here:
// https://א.cc/dHGpnhxX/c.

// Append * to find all named pipes.

// FindClose is used to close file search handles instead of the typical
// CloseHandle used elsewhere, see:
// https://docs.microsoft.com/en-us/windows/desktop/api/fileapi/nf-fileapi-findclose.

// Check the first file's name for a match, but also keep searching for
// WireGuard named pipes until no more files can be iterated.

// Concatenate strings directly as filepath.Join appears to break the
// named pipe prefix convention.
