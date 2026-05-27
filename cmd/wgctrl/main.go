// Command wgctrl is a testing utility for interacting with WireGuard via package
// wgctrl.
package main

import (
	"flag"
	"log"
	"net"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {
	flag.Parse()

	c, err := wgctrl.New()
	if err != nil {
		log.Fatalf("failed to open wgctrl: %v", err)
	}
	defer c.Close()

	var devices []*wgtypes.Device
	if device := flag.Arg(0); device != "" {
		d, err := c.Device(device)
		if err != nil {
			log.Fatalf("failed to get device %q: %v", device, err)
		}

		devices = append(devices, d)
	} else {
		devices, err = c.Devices()
		if err != nil {
			log.Fatalf("failed to get devices: %v", err)
		}
	}

	for _, d := range devices {
		printDevice(d)

		for _, p := range d.Peers {
			printPeer(p)
		}
	}
}

func printDevice(d *wgtypes.Device) { _ = "STUB: not implemented"; return }

func printPeer(p wgtypes.Peer) { _ = "STUB: not implemented"; return }

// TODO(mdlayher): get right endpoint with getnameinfo.

func ipsString(ipns []net.IPNet) string { _ = "STUB: not implemented"; return "" }
