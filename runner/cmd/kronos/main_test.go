package main

import (
	"errors"
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
		"scripts/run-lab.sh",
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
