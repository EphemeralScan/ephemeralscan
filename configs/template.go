// Package configs provides configuration templates distributed with EphemeralScan.
package configs

import _ "embed"

//go:embed ephemeralscan.yaml.example
var template []byte

// Template returns a copy of the default YAML configuration template.
func Template() []byte {
	return append([]byte(nil), template...)
}
