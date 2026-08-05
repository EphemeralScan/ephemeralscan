package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

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
