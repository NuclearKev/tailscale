//go:build netbsd

package routetable

import "golang.org/x/sys/unix"

// NetBSD route flag values (match NetBSD kernel headers)
const (
	RTF_BROADCAST = 0x80000
	RTF_LOCAL     = 0x40000
	RTF_MULTICAST = 0x200
)

const (
	ribType        = unix.NET_RT_DUMP
	parseType      = unix.NET_RT_IFLIST
	rmExpectedType = unix.RTM_GET

	// Nothing to skip
	skipFlags = 0
)

var flags = map[int]string{
	unix.RTF_BLACKHOLE: "blackhole",
	RTF_BROADCAST:      "broadcast",
	unix.RTF_GATEWAY:   "gateway",
	unix.RTF_HOST:      "host",
	RTF_LOCAL:          "local",
	RTF_MULTICAST:      "multicast",
	unix.RTF_REJECT:    "reject",
	unix.RTF_STATIC:    "static",
	unix.RTF_UP:        "up",
}
