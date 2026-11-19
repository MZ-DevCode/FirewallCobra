package main

import (
	"encoding/binary"
	"firewallcobra/internal"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ifce, err := internal.SetupDevice()
	if err != nil {
		log.Fatalf("Error setup TUN: %v", err)
	}

	var file *os.File

	file, err = os.OpenFile("firewall.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening firewall.log: %v\n", err)
	}

	defer file.Close()

	blacklist := internal.ReadBlacklist("blacklist.txt")
	log.Print("Loaded nlacklist succesfully")

	log.Printf("TUN interface %s is running", ifce.Name())

	packet := make([]byte, 1500)
	tracker := internal.NewConntrack()

	var flagsStr string

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

		ihl := int(packet[0]&0x0F) * 4
		ttl := packet[8]
		protocolCode := packet[9]

		srcIP := packet[12:16]
		dstIP := packet[16:20]

		transportLayer := packet[ihl:]

		var protocol string = "OTHER"
		var srcPort uint16 = 0
		var dstPort uint16 = 0

		switch protocolCode {
		case 1:
			protocol = "ICMP"
		case 6:
			protocol = "TCP"
			srcPort = binary.BigEndian.Uint16(transportLayer[0:2])
			dstPort = binary.BigEndian.Uint16(transportLayer[2:4])

			tcpFlags := transportLayer[13]

			isSYN := (tcpFlags & 0x02) != 0
			isACK := (tcpFlags & 0x10) != 0
			isFIN := (tcpFlags & 0x01) != 0
			isRST := (tcpFlags & 0x04) != 0

			if isSYN {
				flagsStr = "SYN "
			}
			if isACK {
				flagsStr = "ACK "
			}
			if isFIN {
				flagsStr = "FIN "
			}
			if isRST {
				flagsStr = "RST "
			}

		case 17:
			protocol = "UDP"
			srcPort = binary.BigEndian.Uint16(transportLayer[0:2])
			dstPort = binary.BigEndian.Uint16(transportLayer[2:4])

		}

		srcIPInt := binary.BigEndian.Uint32(srcIP)
		dstIPInt := binary.BigEndian.Uint32(dstIP)

		if blacklist[srcIPInt] {
			srcIPNet := net.IPv4(srcIP[0], srcIP[1], srcIP[2], srcIP[3])
			fmt.Printf("[IP] blocked ip: %s\n", srcIPNet)
			continue
		}

		info, exists := tracker.Observe(
			srcIPInt,
			dstIPInt,
			srcPort,
			dstPort,
			protocolCode,
		)

		line := fmt.Sprintf(
			"[%s] src: %s:%d -> dst: %s:%d | flags: %s | ttl: %d | ihl: %d | len: %d \npackets=%d | existed=%t\n",
			protocol,
			net.IPv4(srcIP[0], srcIP[1], srcIP[2], srcIP[3]),
			srcPort,
			net.IPv4(dstIP[0], dstIP[1], dstIP[2], dstIP[3]),
			dstPort,
			flagsStr,
			ttl,
			ihl,
			n,
			info.Count,
			exists,
		)

		fmt.Println(line)

		_, err = file.WriteString(line)
	}
}
