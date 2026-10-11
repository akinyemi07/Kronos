package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		root, err := findProjectRoot()
		if err != nil {
			return err
		}

		runnerPath := filepath.Join(root, "scripts", "run-lab.sh")
		args := append([]string{runnerPath, labID}, labArgs...)
		return execute("bash", args...)

	default:
		return fmt.Errorf("unknown lab %q", labID)
	}
}

func findProjectRoot() (string, error) {
	if root := os.Getenv("KRONOS_ROOT"); root != "" {
		root, err := filepath.Abs(root)
		if err != nil {
			return "", fmt.Errorf("resolve KRONOS_ROOT: %w", err)
		}

		if _, err := os.Stat(filepath.Join(root, "scripts", "run-lab.sh")); err != nil {
			return "", fmt.Errorf("KRONOS_ROOT %q does not contain scripts/run-lab.sh", root)
		}

		return root, nil
	}

	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "scripts", "run-lab.sh")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("Kronos project root not found; set KRONOS_ROOT to the repository path")
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
