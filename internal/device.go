package internal

import (
	"log"
	"os/exec"

	"github.com/songgao/water"
)

func SetupDevice() (*water.Interface, error) {
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}

	name := ifce.Name()

	cmd := exec.Command("sudo", "ip", "addr", "add", "10.0.0.1/24", "dev", name)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("sudo", "ip", "link", "set", "dev", name, "up")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("sudo", "ip", "route", "add", "8.8.8.8", "dev", name)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	return ifce, nil
}
