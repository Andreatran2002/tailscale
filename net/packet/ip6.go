// Copyright (c) 2020 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packet

import (
	"fmt"

	"inet.af/netaddr"
)

type IP6 [16]byte

func IP6FromNetaddr(ip netaddr.IP) IP6 {
	if !ip.Is6() {
		panic(fmt.Sprintf("IP6FromNetaddr called with non-v6 addr %q", ip))
	}
	return IP6(ip.As16())
}

func (ip IP6) Netaddr() netaddr.IP {
	return netaddr.IPFrom16(ip)
}

func (ip IP6) String() string {
	return ip.Netaddr().String()
}

const ip6HeaderLength = 40
