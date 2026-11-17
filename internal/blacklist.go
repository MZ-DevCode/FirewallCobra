package internal

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func ReadBlacklist(filename string) map[uint32]bool {
	blacklist := make(map[uint32]bool)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return blacklist
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		ip := net.ParseIP(line)
		if ip == nil {
			fmt.Printf("Error: invalid IP format")
			continue
		}
		ipv4 := ip.To4()
		if ipv4 != nil {
			ipInt := binary.BigEndian.Uint32(ipv4)
			blacklist[ipInt] = true
		}

	}
	return blacklist
}
