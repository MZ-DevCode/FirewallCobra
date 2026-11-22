package internal

import (
	"fmt"
	"os"

	"github.com/stretchr/testify/assert/yaml"
	"go.yaml.in/yaml/v3"
)

type RawRule struct {
	Name     string `yaml:"name"`
	Action   string `yaml:"action"`
	Protocol string `yaml:"protocol"`
	SrcIP    string `yaml:"src_ip"`
	DstIP    string `yaml:"dst_ip"`
	SrcPort  int    `yaml:"src_port"`
	DstPort  int    `yaml:"dst_port"`
}

type RawConfig struct {
	DefaultAction string    `yaml:"default_action"`
	Rules         []RawRule `yaml:"rules"`
}

func ParseYaml(path string) (RawConfig, error) {
	var config RawConfig

	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("read yaml file: %w", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("Error unmarshal yaml file: %w", err)
	}
}
