package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config describes an EphemeralScan configuration file.
type Config struct {
	Version int `yaml:"version"`

	Job Job `yaml:"job"`

	Target Target `yaml:"target"`

	Provider Provider `yaml:"provider"`

	Scanners []Scanner `yaml:"scanners"`

	Reports Reports `yaml:"reports"`

	Cleanup Cleanup `yaml:"cleanup"`
}

// Job contains execution metadata.
type Job struct {
	Name string `yaml:"name"`
}

// Target identifies the system to assess.
type Target struct {
	URL string `yaml:"url"`
}

// Provider describes the temporary infrastructure provider.
type Provider struct {
	Name         string `yaml:"name"`
	Region       string `yaml:"region"`
	InstanceType string `yaml:"instance_type"`
}

// Scanner describes a scanner and its execution profile.
type Scanner struct {
	Name    string `yaml:"name"`
	Profile string `yaml:"profile"`
}

// Reports configures report output.
type Reports struct {
	Directory string   `yaml:"directory"`
	Formats   []string `yaml:"formats"`
}

// Cleanup configures infrastructure cleanup.
type Cleanup struct {
	AlwaysDestroy bool `yaml:"always_destroy"`
}

// Load reads, parses, and validates a configuration file at path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read configuration %q: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse configuration %q: %w", path, err)
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("validate configuration %q: %w", path, err)
	}

	return &cfg, nil
}

func (c Config) validate() error {
	if c.Version == 0 {
		return fmt.Errorf("version is required")
	}

	if strings.TrimSpace(c.Target.URL) == "" {
		return fmt.Errorf("target.url is required")
	}

	if !c.Cleanup.AlwaysDestroy {
		return fmt.Errorf("cleanup.always_destroy must be true")
	}

	return nil
}
