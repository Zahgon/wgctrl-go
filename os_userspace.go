//go:build !linux && !openbsd && !windows && !freebsd
// +build !linux,!openbsd,!windows,!freebsd

package wgctrl

import (
	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
)

// newClients configures wginternal.Clients for systems which only support
// userspace WireGuard implementations.
func newClients() ([]wginternal.Client, error) { _ = "STUB: not implemented"; return nil, nil }
