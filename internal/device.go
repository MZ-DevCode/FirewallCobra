package internal

import (
	"log"
	"os/exec"

	"github.com/songgao/water"
)

func SetupDevice(*water.Interface, error) {
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}

	name := ifce.Name()

	cmd := exec.Command("sudo", "ip", "addr", "10.0.0.1/24", "dev", name)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("sudo", "ip", "link", "set", "dev", name, "up")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	packet := make([]byte, 2000)
	for {
		n, err := ifce.Read(packet)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Packet Received: % x\n", packet[:n])
	}
}
