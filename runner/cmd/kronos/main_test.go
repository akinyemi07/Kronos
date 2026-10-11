package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunLabRejectsUnknownLab(t *testing.T) {
	called := false

	fakeRunner := func(name string, args ...string) error {
		called = true
		return nil
	}

	err := runLab("unknown-lab", nil, fakeRunner)

	if err == nil {
		t.Fatal("expected an error for an unknown lab")
	}

	if !strings.Contains(err.Error(), `unknown lab "unknown-lab"`) {
		t.Fatalf("unexpected error: %v", err)
	}

	if called {
		t.Fatal("command runner should not run for an unknown lab")
	}
}

func TestRunLabPassesArgumentsToRunner(t *testing.T) {
	root := t.TempDir()

	if err := os.MkdirAll(filepath.Join(root, "scripts"), 0755); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(
		filepath.Join(root, "scripts", "run-lab.sh"),
		[]byte("#!/usr/bin/env bash\n"),
		0644,
	); err != nil {
		t.Fatal(err)
	}

	t.Setenv("KRONOS_ROOT", root)

	var gotName string
	var gotArgs []string

	fakeRunner := func(name string, args ...string) error {
		gotName = name
		gotArgs = args
		return nil
	}

	err := runLab("linux-timezone", []string{"Africa/Lagos"}, fakeRunner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotName != "bash" {
		t.Fatalf("expected command bash, got %q", gotName)
	}

	expectedArgs := []string{
		filepath.Join(root, "scripts", "run-lab.sh"),
		"linux-timezone",
		"Africa/Lagos",
	}

	if len(gotArgs) != len(expectedArgs) {
		t.Fatalf("expected arguments %v, got %v", expectedArgs, gotArgs)
	}

	for i, expected := range expectedArgs {
		if gotArgs[i] != expected {
			t.Fatalf("argument %d: expected %q, got %q", i, expected, gotArgs[i])
		}
	}
}

func TestRunLabReturnsRunnerError(t *testing.T) {
	expectedErr := errors.New("simulated command failure")

	fakeRunner := func(name string, args ...string) error {
		return expectedErr
	}

	err := runLab("linux-timezone", nil, fakeRunner)

	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected runner error %v, got %v", expectedErr, err)
	}
}

func TestFindProjectRootSearchesParentDirectories(t *testing.T) {
	root := t.TempDir()
	scriptsDir := filepath.Join(root, "scripts")

	if err := os.MkdirAll(scriptsDir, 0755); err != nil {
		t.Fatal(err)
	}

	runnerFile := filepath.Join(scriptsDir, "run-lab.sh")
	if err := os.WriteFile(runnerFile, []byte("#!/usr/bin/env bash\n"), 0644); err != nil {
		t.Fatal(err)
	}

	nestedDir := filepath.Join(root, "nested", "directory")
	if err := os.MkdirAll(nestedDir, 0755); err != nil {
		t.Fatal(err)
	}

	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(nestedDir); err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	})

	t.Setenv("KRONOS_ROOT", "")

	got, err := findProjectRoot()
	if err != nil {
		t.Fatalf("findProjectRoot returned an error: %v", err)
	}

	if got != root {
		t.Fatalf("expected project root %q, got %q", root, got)
	}
}

func TestFindProjectRootRejectsInvalidEnvironmentPath(t *testing.T) {
	root := t.TempDir()
	t.Setenv("KRONOS_ROOT", root)

	_, err := findProjectRoot()
	if err == nil {
		t.Fatal("expected an error for KRONOS_ROOT without scripts/run-lab.sh")
	}

	if !strings.Contains(err.Error(), "does not contain scripts/run-lab.sh") {
		t.Fatalf("unexpected error: %v", err)
	}
}
