package config

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
