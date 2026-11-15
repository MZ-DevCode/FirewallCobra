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
}
