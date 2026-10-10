package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type commandRunner func(name string, args ...string) error

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: kronos <lab-id> [lab-arguments...]")
		os.Exit(2)
	}

	if err := runLab(os.Args[1], os.Args[2:], executeCommand); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runLab(labID string, labArgs []string, execute commandRunner) error {
	switch labID {
	case "linux-timezone":
		args := append([]string{"scripts/run-lab.sh", labID}, labArgs...)
		return execute("bash", args...)
	default:
		return fmt.Errorf("unknown lab %q", labID)
	}
}

func executeCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return fmt.Errorf("lab exited with status %d", exitErr.ExitCode())
		}
		return fmt.Errorf("running lab: %w", err)
	}

	return nil
}
