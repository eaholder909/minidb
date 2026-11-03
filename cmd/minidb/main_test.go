package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestVersionFlag(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"--version"}, &output)
	if err != nil {
		t.Fatalf("run() returned unexpected error: %v", err)
	}
	got := output.String()
	want := "MiniDB v0.1.0\n"

	if got != want {
		t.Errorf("unexpected output:\ngot:  %q\nwant %q", got, want)
	}
}

func TestHelpOutput(t *testing.T) {
	var output bytes.Buffer
	err := run([]string{"--help"}, &output)
	if err != nil {
		t.Fatalf("run() returned unexpected error: %v", err)
	}

	if !strings.Contains(output.String(), "Usage:") {
		t.Error("expected help output to contain Usage:")
	}
}

func TestUnexpectedArgument(t *testing.T) {
	var output bytes.Buffer

	err := run([]string{"something"}, &output)
	if err == nil {
		t.Fatal("expected an error for unexpected argument")
	}
}
