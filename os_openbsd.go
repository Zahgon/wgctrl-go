//go:build openbsd
// +build openbsd

package wgctrl

import (
	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
)

// newClients configures wginternal.Clients for OpenBSD systems.
func newClients() ([]wginternal.Client, error) { _ = "STUB: not implemented"; return nil, nil }

// OpenBSD has an in-kernel WireGuard implementation. Determine if it is
// available and make use of it if so.
