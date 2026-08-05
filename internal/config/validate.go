package config

import (
	"fmt"
	"strings"
)

func (c Config) validate() error {
	if c.Version == 0 {
		return fmt.Errorf("version is required")
	}

	if strings.TrimSpace(c.Target.URL) == "" {
		return fmt.Errorf("target.url is required")
	}

	if strings.TrimSpace(c.Provider.Name) == "" {
		return fmt.Errorf("provider.name is required")
	}

	if !c.Cleanup.AlwaysDestroy {
		return fmt.Errorf("cleanup.always_destroy must be true")
	}

	return nil
}
