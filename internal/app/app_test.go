package app

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/EphemeralScan/ephemeralscan/configs"
)

func TestConfigInit(t *testing.T) {
	t.Chdir(t.TempDir())

	var stdout bytes.Buffer
	if err := configCommand([]string{"init"}, &stdout); err != nil {
		t.Fatalf("configCommand() error = %v", err)
	}

	got, err := os.ReadFile(configFileName)
	if err != nil {
		t.Fatalf("read created configuration: %v", err)
	}
	if !bytes.Equal(got, configs.Template()) {
		t.Error("created configuration does not match template")
	}
	if !strings.Contains(stdout.String(), "Created ephemeralscan.yaml") {
		t.Errorf("output = %q, want creation message", stdout.String())
	}
}

func TestConfigInitDoesNotOverwriteExistingFile(t *testing.T) {
	t.Chdir(t.TempDir())

	want := []byte("existing configuration\n")
	if err := os.WriteFile(configFileName, want, 0o600); err != nil {
		t.Fatalf("write existing configuration: %v", err)
	}

	var stdout bytes.Buffer
	err := configCommand([]string{"init"}, &stdout)
	if err == nil {
		t.Fatal("configCommand() error = nil, want error")
	}
	if !strings.Contains(err.Error(), "already exists") {
		t.Errorf("configCommand() error = %q, want already exists message", err)
	}

	got, err := os.ReadFile(configFileName)
	if err != nil {
		t.Fatalf("read existing configuration: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("existing configuration = %q, want %q", got, want)
	}
	if stdout.Len() != 0 {
		t.Errorf("output = %q, want no success message", stdout.String())
	}
}

func TestConfigCommandUsage(t *testing.T) {
	err := configCommand(nil, &bytes.Buffer{})
	if err == nil {
		t.Fatal("configCommand() error = nil, want error")
	}
	if !strings.Contains(err.Error(), "ephemeralscan config init") {
		t.Errorf("configCommand() error = %q, want usage", err)
	}
}
