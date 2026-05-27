package wguser

import (
	"io"
	"net"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// The WireGuard userspace configuration protocol is described here:
// https://www.wireguard.com/xplatform/#cross-platform-userspace-implementation.

// getDevice gathers device information from a device specified by its path
// and returns a Device.
func (c *Client) getDevice(device string) (*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get information about this device.

// Parse the device from the incoming data stream.

// TODO(mdlayher): populate interface index too?

// parseDevice parses a Device and its Peers from an io.Reader.
func parseDevice(r io.Reader) (*wgtypes.Device, error) { _ = "STUB: not implemented"; return nil, nil }

// Empty line, done parsing.

// All data is in key=value format.

// A deviceParser accumulates information about a Device and its Peers.
type deviceParser struct {
	d   wgtypes.Device
	err error

	parsePeers    bool
	peers         int
	hsSec, hsNano int
}

// Device returns a Device or any errors that were encountered while parsing
// a Device.
func (dp *deviceParser) Device() (*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Compute remaining fields of the Device now that all parsing is done.

// Parse parses a single key/value pair into fields of a Device.
func (dp *deviceParser) Parse(key, value string) { _ = "STUB: not implemented"; return }

// 0 indicates success, anything else returns an error number that matches
// definitions from errno.h.

// TODO(mdlayher): return actual errno on Linux?

// We've either found the first peer or the next peer.  Stop parsing
// Device fields and start parsing Peer fields, including the public
// key indicated here.

// Are we parsing peer fields?

// Device field parsing.

// curPeer returns the current Peer being parsed so its fields can be populated.
func (dp *deviceParser) curPeer() *wgtypes.Peer { _ = "STUB: not implemented"; return nil }

// peerParse parses a key/value field into the current Peer.
func (dp *deviceParser) peerParse(key, value string) { _ = "STUB: not implemented"; return }

// Assume that we've seen both seconds and nanoseconds and populate this
// field now. However, if both fields were set to 0, assume we have never
// had a successful handshake with this peer, and return a zero-value
// time.Time to our callers.

// parseKey parses a Key from a hex string.
func (dp *deviceParser) parseKey(s string) wgtypes.Key {
	_ = "STUB: not implemented"
	return *new(wgtypes.Key)
}

// parseInt parses an integer from a string.
func (dp *deviceParser) parseInt(s string) int { _ = "STUB: not implemented"; return 0 }

// parseInt64 parses an int64 from a string.
func (dp *deviceParser) parseInt64(s string) int64 { _ = "STUB: not implemented"; return 0 }

// parseAddr parses a UDP address from a string.
func (dp *deviceParser) parseAddr(s string) *net.UDPAddr { _ = "STUB: not implemented"; return nil }

// parseInt parses an address CIDR from a string.
func (dp *deviceParser) parseCIDR(s string) *net.IPNet { _ = "STUB: not implemented"; return nil }
