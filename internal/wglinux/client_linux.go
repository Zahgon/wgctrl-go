//go:build linux
// +build linux

package wglinux

import (
	"syscall"

	"github.com/mdlayher/genetlink"
	"github.com/mdlayher/netlink"
	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var _ wginternal.Client = &Client{}

// A Client provides access to Linux WireGuard netlink information.
type Client struct {
	c      *genetlink.Conn
	family genetlink.Family

	interfaces func() ([]string, error)
}

// New creates a new Client and returns whether or not the generic netlink
// interface is available.
func New() (*Client, bool, error) { _ = "STUB: not implemented"; return nil, false, nil }

// Best effort version of netlink.Config.Strict due to CentOS 7.

// initClient is the internal Client constructor used in some tests.
func initClient(c *genetlink.Conn) (*Client, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// The generic netlink interface is not available.

// By default, gather only WireGuard interfaces using rtnetlink.

// Close implements wginternal.Client.
func (c *Client) Close() error {
	_ = "STUB: not implemented"

	// Devices implements wginternal.Client.
	return nil
}

func (c *Client) Devices() ([]*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	// By default, rtnetlink is used to fetch a list of all interfaces and then
	// filter that list to only find WireGuard interfaces.
	//
	// The remainder of this function assumes that any returned device from this
	// function is a valid WireGuard device.
	return nil, nil
}

// Device implements wginternal.Client.
func (c *Client) Device(name string) (*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	// Don't bother querying netlink with empty input.
	return nil, nil
}

// Fetching a device by interface index is possible as well, but we only
// support fetching by name as it seems to be more convenient in general.

// ConfigureDevice implements wginternal.Client.
func (c *Client) ConfigureDevice(name string, cfg wgtypes.Config) error {
	_ = "STUB: not implemented"
	// Large configurations are split into batches for use with netlink.
	return nil
}

// Request acknowledgement of our request from netlink, even though the
// output messages are unused.  The netlink package checks and trims the
// status code value.

// execute executes a single WireGuard netlink request with the specified command,
// header flags, and attribute arguments.
func (c *Client) execute(command uint8, flags netlink.HeaderFlags, attrb []byte) ([]genetlink.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We don't want to expose netlink errors directly to callers so unpack to
// something more generic.

// Expect all errors to conform to netlink.OpError.

// Convert "no such device" and "not a wireguard device" to an error
// compatible with os.ErrNotExist for easy checking.

// Expose the inner error directly (such as EPERM).

// rtnlInterfaces uses rtnetlink to fetch a list of WireGuard interfaces.
func rtnlInterfaces() ([]string, error) {
	_ = "STUB: not implemented"
	// Use the stdlib's rtnetlink helpers to get ahold of a table of all
	// interfaces, so we can begin filtering it down to just WireGuard devices.
	return nil, nil
}

// parseRTNLInterfaces unpacks rtnetlink messages and returns WireGuard
// interface names.
func parseRTNLInterfaces(msgs []syscall.NetlinkMessage) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only deal with link messages, and they must have an ifinfomsg
// structure appear before the attributes.

// Determine the interface's name and if it's a WireGuard device.

// Found one; append it to the list.

// wgKind is the IFLA_INFO_KIND value for WireGuard devices.
const wgKind = "wireguard"

// isWGKind parses netlink attributes to determine if a link is a WireGuard
// device, then populates ok with the result.
func isWGKind(ok *bool) func(b []byte) error { _ = "STUB: not implemented"; return nil }
