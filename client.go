package wgctrl

import (
	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// Expose an identical interface to the underlying packages.
var _ wginternal.Client = &Client{}

// A Client provides access to WireGuard device information.
type Client struct {
	// Seamlessly use different wginternal.Client implementations to provide an
	// interface similar to wg(8).
	cs []wginternal.Client
}

// New creates a new Client.
func New() (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// Close releases resources used by a Client.
func (c *Client) Close() error { _ = "STUB: not implemented"; return nil }

// Devices retrieves all WireGuard devices on this system.
func (c *Client) Devices() ([]*wgtypes.Device, error) { _ = "STUB: not implemented"; return nil, nil }

// Device retrieves a WireGuard device by its interface name.
//
// If the device specified by name does not exist or is not a WireGuard device,
// an error is returned which can be checked using `errors.Is(err, os.ErrNotExist)`.
func (c *Client) Device(name string) (*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConfigureDevice configures a WireGuard device by its interface name.
//
// Because the zero value of some Go types may be significant to WireGuard for
// Config fields, only fields which are not nil will be applied when
// configuring a device.
//
// If the device specified by name does not exist or is not a WireGuard device,
// an error is returned which can be checked using `errors.Is(err, os.ErrNotExist)`.
func (c *Client) ConfigureDevice(name string, cfg wgtypes.Config) error {
	_ = "STUB: not implemented"
	return nil
}
