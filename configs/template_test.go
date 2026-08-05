package configs

import (
	"bytes"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	want, err := os.ReadFile("ephemeralscan.yaml.example")
	if err != nil {
		t.Fatalf("read example configuration: %v", err)
	}

	got := Template()
	if !bytes.Equal(got, want) {
		t.Error("Template() does not match ephemeralscan.yaml.example")
	}

	got[0] = 0
	if bytes.Equal(got, Template()) {
		t.Error("Template() returned mutable shared data")
	}
}
