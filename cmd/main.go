package main

import (
	"firewallcobra/internal"
	"log"
)

func main() {
	ifce, err := internal.SetupDevice()
	if err != nil {
		log.Fatalf("Error setup TUN: %v", err)
	}

	log.Printf("TUN interface %s is running", ifce.Name())

	mtu := make([]byte, 1500)

	for {
		n, err := ifce.Read(mtu)
		if err != nil {
			log.Printf("Error reading from TUN")
		}

		if n < 20 {
			continue
		}
	}
}
