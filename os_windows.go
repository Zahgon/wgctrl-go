//go:build windows
// +build windows

package wgctrl

import (
	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
)

// newClients configures wginternal.Clients for Windows systems.
func newClients() ([]wginternal.Client, error) { _ = "STUB: not implemented"; return nil, nil }

// Windows has an in-kernel WireGuard implementation.
