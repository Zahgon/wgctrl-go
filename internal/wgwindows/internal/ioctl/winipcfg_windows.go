/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2021 WireGuard LLC. All Rights Reserved.
 */

package ioctl

import (
	"net"
)

// AddressFamily enumeration specifies protocol family and is one of the windows.AF_* constants.
type AddressFamily uint16

// RawSockaddrInet union contains an IPv4, an IPv6 address, or an address family.
// https://docs.microsoft.com/en-us/windows/desktop/api/ws2ipdef/ns-ws2ipdef-_sockaddr_inet
type RawSockaddrInet struct {
	Family AddressFamily
	data   [26]byte
}

func ntohs(i uint16) uint16 { _ = "STUB: not implemented"; return 0 }

func htons(i uint16) uint16 { _ = "STUB: not implemented"; return 0 }

// SetIP method sets family, address, and port to the given IPv4 or IPv6 address and port.
// All other members of the structure are set to zero.
func (addr *RawSockaddrInet) SetIP(ip net.IP, port uint16) error {
	_ = "STUB: not implemented"
	return nil
}

// IP returns IPv4 or IPv6 address, or nil if the address is neither.
func (addr *RawSockaddrInet) IP() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

// Port returns the port if the address if IPv4 or IPv6, or 0 if neither.
func (addr *RawSockaddrInet) Port() uint16 { _ = "STUB: not implemented"; return 0 }
