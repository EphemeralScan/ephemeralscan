package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	path := writeConfig(t, `
version: 1
job:
  name: example-scan
target:
  url: https://example.com
provider:
  name: linode
  region: ca-central
  instance_type: g6-nanode-1
scanners:
  - name: zap
    profile: baseline
reports:
  directory: ./reports
  formats:
    - html
    - json
cleanup:
  always_destroy: true
`)

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Version != 1 {
		t.Errorf("Version = %d, want 1", cfg.Version)
	}
	if cfg.Job.Name != "example-scan" {
		t.Errorf("Job.Name = %q, want %q", cfg.Job.Name, "example-scan")
	}
	if cfg.Target.URL != "https://example.com" {
		t.Errorf("Target.URL = %q, want %q", cfg.Target.URL, "https://example.com")
	}
	if cfg.Provider.InstanceType != "g6-nanode-1" {
		t.Errorf("Provider.InstanceType = %q, want %q", cfg.Provider.InstanceType, "g6-nanode-1")
	}
	if len(cfg.Scanners) != 1 || cfg.Scanners[0].Name != "zap" {
		t.Errorf("Scanners = %#v, want one zap scanner", cfg.Scanners)
	}
	if len(cfg.Reports.Formats) != 2 {
		t.Errorf("Reports.Formats = %#v, want two formats", cfg.Reports.Formats)
	}
	if !cfg.Cleanup.AlwaysDestroy {
		t.Error("Cleanup.AlwaysDestroy = false, want true")
	}
}

func TestLoadValidation(t *testing.T) {
	tests := []struct {
		name    string
		content string
		wantErr string
	}{
		{
			name: "missing version",
			content: `
target:
  url: https://example.com
cleanup:
  always_destroy: true
`,
			wantErr: "version is required",
		},
		{
			name: "missing target URL",
			content: `
version: 1
cleanup:
  always_destroy: true
`,
			wantErr: "target.url is required",
		},
		{
			name: "blank target URL",
			content: `
version: 1
target:
  url: "  "
cleanup:
  always_destroy: true
`,
			wantErr: "target.url is required",
		},
		{
			name: "missing provider name",
			content: `
version: 1
target:
  url: https://example.com
cleanup:
  always_destroy: true
`,
			wantErr: "provider.name is required",
		},
		{
			name: "cleanup disabled",
			content: `
version: 1
target:
  url: https://example.com
provider:
  name: linode
cleanup:
  always_destroy: false
`,
			wantErr: "cleanup.always_destroy must be true",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Load(writeConfig(t, tt.content))
			if err == nil {
				t.Fatal("Load() error = nil, want error")
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("Load() error = %q, want it to contain %q", err, tt.wantErr)
			}
		})
	}
}

func TestLoadInvalidYAML(t *testing.T) {
	_, err := Load(writeConfig(t, "version: [\n"))
	if err == nil {
		t.Fatal("Load() error = nil, want error")
	}
	if !strings.Contains(err.Error(), "parse configuration") {
		t.Errorf("Load() error = %q, want parse context", err)
	}
}

func TestLoadMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.yaml")

	_, err := Load(path)
	if err == nil {
		t.Fatal("Load() error = nil, want error")
	}
	if !strings.Contains(err.Error(), "read configuration") {
		t.Errorf("Load() error = %q, want read context", err)
	}
}

func writeConfig(t *testing.T, content string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "ephemeralscan.yaml")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write configuration: %v", err)
	}

	return path
}
