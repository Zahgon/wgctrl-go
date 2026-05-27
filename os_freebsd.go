//go:build freebsd
// +build freebsd

package wgctrl

import (
	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
)

// newClients configures wginternal.Clients for FreeBSD systems.
func newClients() ([]wginternal.Client, error) { _ = "STUB: not implemented"; return nil, nil }

// FreeBSD has an in-kernel WireGuard implementation. Determine if it is
// available and make use of it if so.
