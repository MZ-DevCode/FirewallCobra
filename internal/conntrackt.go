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

func NewConntrack() *Conntrack {
	return &Conntrack{
		entries: make(map[ConntrackKey]ConntrackInfo),
	}
}

func (ct *Conntrack) Observe(SrcIP, DstIP uint32, SrcPort, DstPort uint16, protocol byte) (ConntrackInfo, bool) {
	key := ConntrackKey{
		SrcIP:    SrcIP,
		DstIP:    DstIP,
		SrcPort:  SrcPort,
		DstPort:  DstPort,
		Protocol: protocol,
	}
}
