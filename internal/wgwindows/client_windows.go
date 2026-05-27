package wgwindows

import (
	"golang.org/x/sys/windows"

	"golang.zx2c4.com/wireguard/wgctrl/internal/wginternal"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var _ wginternal.Client = &Client{}

// A Client provides access to WireGuardNT ioctl information.
type Client struct {
	cachedInterfaces map[string]*uint16
	lastLenGuess     uint32
}

var (
	deviceClassNetGUID     = windows.GUID{0x4d36e972, 0xe325, 0x11ce, [8]byte{0xbf, 0xc1, 0x08, 0x00, 0x2b, 0xe1, 0x03, 0x18}}
	deviceInterfaceNetGUID = windows.GUID{0xcac88484, 0x7515, 0x4c03, [8]byte{0x82, 0xe6, 0x71, 0xa8, 0x7a, 0xba, 0xc3, 0x61}}
	devpkeyWgName          = windows.DEVPROPKEY{
		FmtID: windows.DEVPROPGUID{0x65726957, 0x7547, 0x7261, [8]byte{0x64, 0x4e, 0x61, 0x6d, 0x65, 0x4b, 0x65, 0x79}},
		PID:   windows.DEVPROPID_FIRST_USABLE + 1,
	}
)

var enumerator = `SWD\WireGuard`

func init() {
	if maj, min, _ := windows.RtlGetNtVersionNumbers(); (maj == 6 && min <= 1) || maj < 6 {
		enumerator = `ROOT\WIREGUARD`
	}
}

func (c *Client) refreshInterfaceCache() error { _ = "STUB: not implemented"; return nil }

func (c *Client) interfaceHandle(name string) (handle windows.Handle, err error) {
	_ = "STUB: not implemented"
	return *new(windows.Handle), nil
}

// Devices implements wginternal.Client.
func (c *Client) Devices() ([]*wgtypes.Device, error) { _ = "STUB: not implemented"; return nil, nil }

// New creates a new Client
func New() *Client {
	_ = "STUB: not implemented"

	// Close implements wginternal.Client.
	return nil
}

func (c *Client) Close() error {
	_ = "STUB: not implemented"

	// Device implements wginternal.Client.
	return nil
}

func (c *Client) Device(name string) (*wgtypes.Device, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConfigureDevice implements wginternal.Client.
func (c *Client) ConfigureDevice(name string, cfg wgtypes.Config) error {
	_ = "STUB: not implemented"
	return nil
}
