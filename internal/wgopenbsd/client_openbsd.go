//go:build openbsd
// +build openbsd

package wgopenbsd

import (
	"net"
	"unsafe"

	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
	"golang.zx2c4.com/wireguard/wgctrl/internal/wgopenbsd/internal/wgh"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// ifGroupWG is the WireGuard interface group name passed to the kernel.
var ifGroupWG = [16]byte{0: 'w', 1: 'g'}

var _ wginternal.Client = &Client{}

// A Client provides access to OpenBSD WireGuard ioctl information.
type Client struct {
	// Hooks which use system calls by default, but can also be swapped out
	// during tests.
	close           func() error
	ioctlIfgroupreq func(ifg *wgh.Ifgroupreq) error
	ioctlWGDataIO   func(data *wgh.WGDataIO) error
}

// New creates a new Client and returns whether or not the ioctl interface
// is available.
func New() (*Client, bool, error) {
	_ = "STUB: not implemented"
	// The OpenBSD ioctl interface operates on a generic AF_INET socket.
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
// must be allocated in order to store the WGInterfaceIO structure and
// any trailing WGPeerIO/WGAIPIOs.

// TODO: consider preallocating some memory to avoid a second system call
// if it proves to be a concern.

// ioctl functions always return a wrapped unix.Errno value.
// Conform to the wgctrl contract by unwrapping some values:
//   ENXIO: "no such device": (no such WireGuard device)
//   ENOTTY: "inappropriate ioctl for device" (device is not a
//	   WireGuard device)

// Allocated enough memory!

// Ensure we don't unsafe cast into uninitialized memory. We need at very
// least a single WGInterfaceIO with no peers.

// Allocate the appropriate amount of memory and point the kernel at
// the first byte of our slice's backing array. When the loop continues,
// we will check if we've allocated enough memory.

// parseDevice unpacks a Device from ifio, along with its associated peers
// and their allowed IPs.
func parseDevice(name string, ifio *wgh.WGInterfaceIO) (*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The kernel populates ifio.Flags to indicate which fields are present.

// If there were no peers, exit early so we do not advance the pointer
// beyond the end of the WGInterfaceIO structure.

// Set our pointer to the beginning of the first peer's location in memory.

// Same idea, we know how many allowed IPs we need to account for, so
// reserve the space and advance the pointer through each WGAIP structure.

// Prepare for the next iteration.

// ConfigureDevice implements wginternal.Client.
func (c *Client) ConfigureDevice(name string, cfg wgtypes.Config) error {
	_ = "STUB: not implemented"
	// Currently read-only: we must determine if a device belongs to this driver,
	// and if it does, return a sentinel so integration tests that configure a
	// device can be skipped.
	return nil
}

// deviceName converts an interface name string to the format required to pass
// with wgh.WGGetServ.
func deviceName(name string) ([16]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// parsePeer unpacks a wgtypes.Peer from a WGPeerIO structure.
func parsePeer(pio *wgh.WGPeerIO) wgtypes.Peer {
	_ = "STUB: not implemented"
	return *new(wgtypes.Peer)
}

// Only set last handshake if a non-zero timespec was provided, matching
// the time.Time.IsZero() behavior of internal/wglinux.

// Conversion required for GOARCH=386.

// parseAllowedIP unpacks a net.IPNet from a WGAIP structure.
func parseAllowedIP(aip *wgh.WGAIPIO) net.IPNet { _ = "STUB: not implemented"; return *new(net.IPNet) }

// parseEndpoint parses a peer endpoint from a wgh.WGIP structure.
func parseEndpoint(ep [28]byte) *net.UDPAddr {
	_ = "STUB: not implemented"
	// sockaddr* structures have family at index 1.
	return nil
}

// TODO(mdlayher): IPv6 zone?

// No endpoint configured.

// bePort interprets a port integer stored in native endianness as a big
// endian value. This is necessary for proper endpoint port handling on
// little endian machines.
func bePort(port uint16) int { _ = "STUB: not implemented"; return 0 }

// ioctlIfgroupreq returns a function which performs the appropriate ioctl on
// fd to retrieve members of an interface group.
func ioctlIfgroupreq(fd int) func(*wgh.Ifgroupreq) error { _ = "STUB: not implemented"; return nil }

// ioctlWGDataIO returns a function which performs the appropriate ioctl on
// fd to issue a WireGuard data I/O.
func ioctlWGDataIO(fd int) func(*wgh.WGDataIO) error { _ = "STUB: not implemented"; return nil }

// ioctl is a raw wrapper for the ioctl system call.
func ioctl(fd int, req uint, arg unsafe.Pointer) error {
	_ = "STUB: not implemented"
	//lint:ignore SA1019 temporarily permitted until we switch to a libc wrapper
	return nil
}

func panicf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }
