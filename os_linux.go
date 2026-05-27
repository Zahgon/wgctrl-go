//go:build linux
// +build linux

package wgctrl

import (
	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
)

// newClients configures wginternal.Clients for Linux systems.
func newClients() ([]wginternal.Client, error) { _ = "STUB: not implemented"; return nil, nil }

// Linux has an in-kernel WireGuard implementation. Determine if it is
// available and make use of it if so.

// Although it isn't recommended to use userspace implementations on Linux,
// it can be used. We make use of it in integration tests as well.

// Kernel devices seem to appear first in wg(8).
