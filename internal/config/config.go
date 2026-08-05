package config

type Config struct {
	Version int `yaml:"version"`

	Job Job `yaml:"job"`

	Target Target `yaml:"target"`

	Provider Provider `yaml:"provider"`

	Scanners []Scanner `yaml:"scanners"`

	Reports Reports `yaml:"reports"`

	Cleanup Cleanup `yaml:"cleanup"`
}

type Job struct {
	Name string `yaml:"name"`
}

type Target struct {
	URL string `yaml:"url"`
}

type Provider struct {
	Name         string `yaml:"name"`
	Region       string `yaml:"region"`
	InstanceType string `yaml:"instance_type"`
}

type Scanner struct {
	Name    string `yaml:"name"`
	Profile string `yaml:"profile"`
}

type Reports struct {
	Directory string   `yaml:"directory"`
	Formats   []string `yaml:"formats"`
}

type Cleanup struct {
	AlwaysDestroy bool `yaml:"always_destroy"`
}
