package wguser

import (
	"io"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// configureDevice configures a device specified by its path.
func (c *Client) configureDevice(device string, cfg wgtypes.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// Start with set command.

// Add any necessary configuration from cfg, then finish with an empty line.

// Apply configuration for the device and then check the error number.

// errno=0 indicates success, anything else returns an error number that
// matches definitions from errno.h.

// TODO(mdlayher): return actual errno on Linux?

// writeConfig writes textual configuration to w as specified by cfg.
func writeConfig(w io.Writer, cfg wgtypes.Config) { _ = "STUB: not implemented"; return }

// hexKey encodes a wgtypes.Key into a hexadecimal string.
func hexKey(k wgtypes.Key) string { _ = "STUB: not implemented"; return "" }
