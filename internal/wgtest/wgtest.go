package wgtest

import (
	"net"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// MustCIDR converts CIDR string s into a net.IPNet or panics.
func MustCIDR(s string) net.IPNet { _ = "STUB: not implemented"; return *new(net.IPNet) }

// MustHexKey decodes a hex string s as a key or panics.
func MustHexKey(s string) wgtypes.Key { _ = "STUB: not implemented"; return *new(wgtypes.Key) }

// MustPresharedKey generates a preshared key or panics.
func MustPresharedKey() wgtypes.Key { _ = "STUB: not implemented"; return *new(wgtypes.Key) }

// MustPrivateKey generates a private key or panics.
func MustPrivateKey() wgtypes.Key { _ = "STUB: not implemented"; return *new(wgtypes.Key) }

// MustPublicKey generates a public key or panics.
func MustPublicKey() wgtypes.Key { _ = "STUB: not implemented"; return *new(wgtypes.Key) }

// MustUDPAddr parses s as a UDP address or panics.
func MustUDPAddr(s string) *net.UDPAddr { _ = "STUB: not implemented"; return nil }

func panicf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }
