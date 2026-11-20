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

func NewConntrack( /*timeout time.Duration*/ ) *Conntrack {
	/*	if timeout <= 0 {
		timeout = time.Minute * 2
		}*/
	return &Conntrack{
		entries: make(map[ConntrackKey]ConntrackInfo),
		/* 	timeout: timeout, */
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

	val, ok := ct.entries[key]
	if !ok {
		val = ConntrackInfo{
			State: "SEEN",
		}
	}

	val.Count++
	val.Timestamp = time.Now()

	ct.entries[key] = val

	return val, !ok

}
