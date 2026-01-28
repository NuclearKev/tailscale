// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux

//--------------------------------This issues have to solve-----------------
// 1️⃣ Temporary boot test
// At the bootloader:
// boot -d
// boot netbsd -o ip.maxsockbuf=16777216
// If that fixes everything → kernel limit confirmed.
// 2️⃣ Permanent fix (advanced)
// Rebuild kernel with:
// options MAXSOCKBUF=16777216
// Then reinstall kernel and reboot.
// This is outside Tailscale’s control.

package prober

import (
	"fmt"
	"net/netip"

	"github.com/tailscale/netlink"
	"go4.org/netipx"
)

const tunName = "derpprobe"

func configureTUN(addr netip.Prefix, tunname string) error {
	link, err := netlink.LinkByName(tunname)
	if err != nil {
		return fmt.Errorf("failed to look up link %q: %w", tunname, err)
	}

	// We need to bring the TUN device up before assigning an address. This
	// allows the OS to automatically create a route for it. Otherwise, we'd
	// have to manually create the route.
	if err := netlink.LinkSetUp(link); err != nil {
		return fmt.Errorf("failed to bring tun %q up: %w", tunname, err)
	}

	if err := netlink.AddrReplace(link, &netlink.Addr{IPNet: netipx.PrefixIPNet(addr)}); err != nil {
		return fmt.Errorf("failed to add address: %w", err)
	}

	return nil
}
