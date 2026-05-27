package wguser

import (
	"net"

	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var _ wginternal.Client = &Client{}

// A Client provides access to userspace WireGuard device information.
type Client struct {
	dial func(device string) (net.Conn, error)
	find func() ([]string, error)
}

// New creates a new Client.
func New() (*Client, error) {
	_ = "STUB: not implemented"

	// Operating system-specific functions which can identify and connect
	// to userspace WireGuard devices. These functions can also be
	// overridden for tests.
	return nil, nil
}

// Close implements wginternal.Client.
func (c *Client) Close() error {
	_ = "STUB: not implemented"

	// Devices implements wginternal.Client.
	return nil
}

func (c *Client) Devices() ([]*wgtypes.Device, error) { _ = "STUB: not implemented"; return nil, nil }

// Device implements wginternal.Client.
func (c *Client) Device(name string) (*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConfigureDevice implements wginternal.Client.
func (c *Client) ConfigureDevice(name string, cfg wgtypes.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// deviceName infers a device name from an absolute file path with extension.
func deviceName(sock string) string { _ = "STUB: not implemented"; return "" }

func panicf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }
