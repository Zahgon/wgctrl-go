//go:build linux
// +build linux

package wglinux

import (
	"net"

	"github.com/mdlayher/netlink"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// configAttrs creates the required encoded netlink attributes to configure
// the device specified by name using the non-nil fields in cfg.
func configAttrs(name string, cfg wgtypes.Config) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only apply peer attributes if necessary.

// Netlink arrays use type as an array index.

// ipBatchChunk is a tunable allowed IP batch limit per peer.
//
// Because we don't necessarily know how much space a given peer will occupy,
// we play it safe and use a reasonably small value.  Note that this constant
// is used both in this package and tests, so be aware when making changes.
const ipBatchChunk = 256

// peerBatchChunk specifies the number of peers that can appear in a
// configuration before we start splitting it into chunks.
const peerBatchChunk = 32

// shouldBatch determines if a configuration is sufficiently complex that it
// should be split into batches.
func shouldBatch(cfg wgtypes.Config) bool { _ = "STUB: not implemented"; return false }

// buildBatches produces a batch of configs from a single config, if needed.
func buildBatches(cfg wgtypes.Config) []wgtypes.Config {
	_ = "STUB: not implemented"
	// Is this a small configuration; no need to batch?
	return nil
}

// Use most fields of cfg for our "base" configuration, and only differ
// peers in each batch.

// Track the known peers so that peer IPs are not replaced if a single
// peer has its allowed IPs split into multiple batches.

// Iterate until no more allowed IPs.

// IPs all fit within a batch; we are done.

// IPs are larger than a single batch, copy a batch out and
// advance the cursor.

// IPs ended on a batch boundary; no more IPs left so end
// iteration after this loop.

// PublicKey denotes the peer and must be present.

// Apply the update only flag to every chunk to ensure
// consistency between batches when the kernel module processes
// them.

// It'd be a bit weird to have a remove peer message with many
// IPs, but just in case, add this to every peer's message.

// The IPs for this chunk.

// Only pass certain fields on the first occurrence of a peer, so
// that subsequent IPs won't be wiped out and space isn't wasted.

// Important: do not move or appending peers won't work.

// Add a peer configuration to this batch and keep going.

// Do not allow peer replacement beyond the first message in a batch,
// so we don't overwrite our previous batch work.

// encodePeer returns a function to encode PeerConfig nested attributes.
func encodePeer(p wgtypes.PeerConfig) func(ae *netlink.AttributeEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// Flags are stored in a single attribute.

// Only apply allowed IPs if necessary.

// encodeSockaddr returns a function which encodes a net.UDPAddr as raw
// sockaddr_in or sockaddr_in6 bytes.
func encodeSockaddr(endpoint net.UDPAddr) func() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil
}

// Is this an IPv6 address?

// IPv4 address handling.

// encodeAllowedIPs returns a function to encode allowed IP nested attributes.
func encodeAllowedIPs(ipns []net.IPNet) func(ae *netlink.AttributeEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure address is 4 bytes if IPv4.

// Netlink arrays use type as an array index.

// isValidIP determines if IP is a valid IPv4 or IPv6 address.
func isValidIP(ip net.IP) bool { _ = "STUB: not implemented"; return false }

// isIPv6 determines if IP is a valid IPv6 address.
func isIPv6(ip net.IP) bool { _ = "STUB: not implemented"; return false }

// sockaddrPort interprets port as a big endian uint16 for use passing sockaddr
// structures to the kernel.
func sockaddrPort(port int) uint16 { _ = "STUB: not implemented"; return 0 }
