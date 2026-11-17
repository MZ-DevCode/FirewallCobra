package internal

import "time"

type ConntracktKey struct {
	SrcIP    uint32
	DstIP    uint32
	SrcPort  uint16
	DstPort  uint16
	Protocol byte
}

type ConntracktInfo struct {
	State     string
	Timestamp time.Time
	Count     int
}
