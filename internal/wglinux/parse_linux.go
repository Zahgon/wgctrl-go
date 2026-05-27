//go:build linux
// +build linux

package wglinux

import (
	"net"
	"time"
	"unsafe"

	"github.com/mdlayher/genetlink"
	"github.com/mdlayher/netlink"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// parseDevice parses a Device from a slice of generic netlink messages,
// automatically merging peer lists from subsequent messages into the Device
// from the first message.
func parseDevice(msgs []genetlink.Message) (*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First message contains our target device.

// Gather the known peers so that we can merge
// them later if needed

// Any subsequent messages have their peer contents merged into the
// first "target" message.

// parseDeviceLoop parses a Device from a single generic netlink message.
func parseDeviceLoop(m genetlink.Message) (*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ignored; interface index isn't exposed at all in the userspace
// configuration protocol, and name is more friendly anyway.

// Netlink array of peers.
//
// Errors while parsing are propagated up to top-level ad.Err check.

// Initialize to the number of peers in this decoder and begin
// handling nested Peer attributes.

// parseAllowedIPs parses a wgtypes.Peer from a netlink attribute payload.
func parsePeer(ad *netlink.AttributeDecoder) wgtypes.Peer {
	_ = "STUB: not implemented"
	return *new(wgtypes.Peer)
}

// parseAllowedIPs parses a slice of net.IPNet from a netlink attribute payload.
func parseAllowedIPs(ipns *[]net.IPNet) func(ad *netlink.AttributeDecoder) error {
	_ = "STUB: not implemented"
	return nil
}

// Initialize to the number of allowed IPs and begin iterating through
// the netlink array to decode each one.

// Allowed IP nested attributes.

// The address family determines the correct number of bits in
// the mask.

// parseKey parses a wgtypes.Key from a byte slice.
func parseKey(key *wgtypes.Key) func(b []byte) error { _ = "STUB: not implemented"; return nil }

// parseAddr parses a net.IP from raw in_addr or in6_addr struct bytes.
func parseAddr(ip *net.IP) func(b []byte) error { _ = "STUB: not implemented"; return nil }

// Okay to convert directly to net.IP; memory layout is identical.

// parseSockaddr parses a *net.UDPAddr from raw sockaddr_in or sockaddr_in6 bytes.
func parseSockaddr(endpoint *net.UDPAddr) func(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// IPv4 address parsing.

// IPv6 address parsing.

// timespec32 is a unix.Timespec with 32-bit integers.
type timespec32 struct {
	Sec  int32
	Nsec int32
}

// timespec64 is a unix.Timespec with 64-bit integers.
type timespec64 struct {
	Sec  int64
	Nsec int64
}

const (
	sizeofTimespec32 = int(unsafe.Sizeof(timespec32{}))
	sizeofTimespec64 = int(unsafe.Sizeof(timespec64{}))
)

// parseTimespec parses a time.Time from raw timespec bytes.
func parseTimespec(t *time.Time) func(b []byte) error { _ = "STUB: not implemented"; return nil }

// It would appear that WireGuard can return a __kernel_timespec which
// uses 64-bit integers, even on 32-bit platforms. Clarification of this
// behavior is being sought in:
// https://lists.zx2c4.com/pipermail/wireguard/2019-April/004088.html.
//
// In the mean time, be liberal and accept 32-bit and 64-bit variants.

// Only set fields if UNIX timestamp value is greater than 0, so the
// caller will see a zero-value time.Time otherwise.

// mergeDevices merges Peer information from d into target.  mergeDevices is
// used to deal with multiple incoming netlink messages for the same device.
func mergeDevices(target, d *wgtypes.Device, knownPeers map[wgtypes.Key]int) {
	_ = "STUB: not implemented"
	return

	// Peer is already known, append to it's allowed IP networks
}

// New peer, add it to the target peers.
