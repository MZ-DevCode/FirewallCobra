package main

import (
	"firewallcobra/internal"
	"fmt"
	"log"
	"net"
)

func main() {
	ifce, err := internal.SetupDevice()
	if err != nil {
		log.Fatalf("Error setup TUN: %v", err)
	}

	log.Printf("TUN interface %s is running", ifce.Name())

	packet := make([]byte, 1500)

	for {
		n, err := ifce.Read(packet)
		if err != nil {
			log.Printf("Error reading from TUN: %v", err)
			continue
		}

		if n < 20 {
			continue
		}

		version := packet[0] >> 4
		if version != 4 {
			continue
		}

		ihl := packet[0] & 0x0F
		ttl := packet[8]
		protocolCode := packet[9]

		srcIP := net.IPv4(packet[12], packet[13], packet[14], packet[15])
		dstIP := net.IPv4(packet[16], packet[17], packet[18], packet[19])

		var protocol string = "OTHER"

		switch protocolCode {
		case 1:
			protocol = "ICMP"
		case 6:
			protocol = "TCP"
		case 17:
			protocol = "UDP"
		}

		fmt.Printf("[%s] src: %s -> dst: %s | ttl: %d | ihl: %d | len: %d\n",
			protocol, srcIP, dstIP, ttl, ihl, n)
	}
}
