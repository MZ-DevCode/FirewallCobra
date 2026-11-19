package internal

import "time"

type ConntrackKey struct {
	SrcIP    uint32
	DstIP    uint32
	SrcPort  uint16
	DstPort  uint16
	Protocol byte
}

type ConntrackInfo struct {
	State     string
	Timestamp time.Time
	Count     int
}

type Conntrack struct {
	entries map[ConntrackKey]ConntrackInfo
	timeout time.Duration
}
