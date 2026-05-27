//go:build freebsd
// +build freebsd

package wgfreebsd

// #include <stdlib.h>
// #include <netinet/in.h>
import "C"

import (
	"net"
	"time"
	"unsafe"

	"golang.zx2c4.com/wireguard/wgctrl/internal/wgfreebsd/internal/nv"
	"golang.zx2c4.com/wireguard/wgctrl/internal/wgfreebsd/internal/wgh"
	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// ifGroupWG is the WireGuard interface group name passed to the kernel.
var ifGroupWG = [16]byte{0: 'w', 1: 'g'}

var _ wginternal.Client = &Client{}

// A Client provides access to FreeBSD WireGuard ioctl information.
type Client struct {
	// Hooks which use system calls by default, but can also be swapped out
	// during tests.
	close           func() error
	ioctlIfgroupreq func(*wgh.Ifgroupreq) error
	ioctlWGDataIO   func(uint, *wgh.WGDataIO) error
}

// New creates a new Client and returns whether or not the ioctl interface
// is available.
func New() (*Client, bool, error) {
	_ = "STUB: not implemented"
	// The FreeBSD ioctl interface operates on a generic AF_INET socket.
	return nil, false, nil
}

// TODO(mdlayher): find a call to invoke here to probe for availability.
// c.Devices won't work because it returns a "not found" error when the
// kernel WireGuard implementation is available but the interface group
// has no members.

// By default, use system call implementations for all hook functions.

// Close implements wginternal.Client.
func (c *Client) Close() error {
	_ = "STUB: not implemented"

	// Devices implements wginternal.Client.
	return nil
}

func (c *Client) Devices() ([]*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	return nil,

		// Query for devices in the "wg" group.
		nil
}

// Determine how many device names we must allocate memory for.

// ifg.Len is size in bytes; allocate enough memory for the correct number
// of wgh.Ifgreq and then store a pointer to the memory where the data
// should be written (ifgrs) in ifg.Groups.
//
// From a thread in golang-nuts, this pattern is valid:
// "It would be OK to pass a pointer to a struct to ioctl if the struct
// contains a pointer to other Go memory, but the struct field must have
// pointer type."
// See: https://groups.google.com/forum/#!topic/golang-nuts/FfasFTZvU_o.

// Now actually fetch the device names.

// Keep this alive until we're done doing the ioctl dance.

// Remove any trailing NULL bytes from the interface names.

// Device implements wginternal.Client.
func (c *Client) Device(name string) (*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// First, specify the name of the device and determine how much memory
// must be allocated.

// ioctl functions always return a wrapped unix.Errno value.
// Conform to the wgctrl contract by unwrapping some values:
//   ENXIO: "no such device": (no such WireGuard device)
//   EINVAL: "inappropriate ioctl for device" (device is not a
//	   WireGuard device)

// Allocated enough memory!

// Allocate the appropriate amount of memory and point the kernel at
// the first byte of our slice's backing array. When the loop continues,
// we will check if we've allocated enough memory.

// ConfigureDevice implements wginternal.Client.
func (c *Client) ConfigureDevice(name string, cfg wgtypes.Config) error {
	_ = "STUB: not implemented"
	// Check if there is a peer with the UpdateOnly flag set.
	// This is not supported on FreeBSD yet. So error out..
	// TODO(stv0g): remove this check once kernel support has landed.
	return nil
}

// Check that this device is really an existing kernel
// device

// ioctl functions always return a wrapped unix.Errno value.
// Conform to the wgctrl contract by unwrapping some values:
//   ENXIO: "no such device": (no such WireGuard device)
//   EINVAL: "inappropriate ioctl for device" (device is not a
//	   WireGuard device)

// deviceName converts an interface name string to the format required to pass
// with wgh.WGGetServ.
func deviceName(name string) ([16]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// ioctlIfgroupreq returns a function which performs the appropriate ioctl on
// fd to retrieve members of an interface group.
func ioctlIfgroupreq(fd int) func(*wgh.Ifgroupreq) error { _ = "STUB: not implemented"; return nil }

// ioctlWGDataIO returns a function which performs the appropriate ioctl on
// fd to issue a WireGuard data I/O.
func ioctlWGDataIO(fd int) func(uint, *wgh.WGDataIO) error { _ = "STUB: not implemented"; return nil }

// ioctl is a raw wrapper for the ioctl system call.
func ioctl(fd int, req uint, arg unsafe.Pointer) error { _ = "STUB: not implemented"; return nil }

func panicf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

func ntohs(i uint16) int { _ = "STUB: not implemented"; return 0 }

func htons(i int) uint16 { _ = "STUB: not implemented"; return 0 }

// parseEndpoint converts a struct sockaddr to a Go net.UDPAddr
func parseEndpoint(ep []byte) *net.UDPAddr { _ = "STUB: not implemented"; return nil }

// TODO(mdlayher): IPv6 zone?

// No endpoint configured.

func unparseEndpoint(ep net.UDPAddr) []byte { _ = "STUB: not implemented"; return nil }

// parseAllowedIP unpacks a net.IPNet from a WGAIP structure.
func parseAllowedIP(aip nv.List) net.IPNet { _ = "STUB: not implemented"; return *new(net.IPNet) }

func unparseAllowedIP(aip net.IPNet) nv.List { _ = "STUB: not implemented"; return *new(nv.List) }

// parseTimestamp parses a binary timestamp to a Go time.Time
func parseTimestamp(b []byte) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// TODO(stv0g): Handle non-little endian machines

// parsePeer unpacks a wgtypes.Peer from a name-value list (nvlist).
func parsePeer(v nv.List) wgtypes.Peer { _ = "STUB: not implemented"; return *new(wgtypes.Peer) }

// parseDevice decodes the device from a FreeBSD name-value list (nvlist)
func parseDevice(data []byte) (*wgtypes.Device, error) { _ = "STUB: not implemented"; return nil, nil }

// unparsePeerConfig encodes a PeerConfig to a name-value list (nvlist).
func unparsePeerConfig(cfg wgtypes.PeerConfig) nv.List {
	_ = "STUB: not implemented"
	return *new(nv.List)
}

// unparseDevice encodes the device configuration as a FreeBSD name-value list (nvlist).
func unparseConfig(cfg wgtypes.Config) nv.List { _ = "STUB: not implemented"; return *new(nv.List) }
