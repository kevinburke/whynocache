package main

import (
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
	e := os.Getenv("UUID_ENV_VAR")
	if e == "" {
		t.Fatal("expected UUID_ENV_VAR to be set, got empty string")
	}
	if len(e) != 36 {
		t.Fatalf("expected UUID_ENV_VAR to be length 36, got %d", len(e))
	}
}
