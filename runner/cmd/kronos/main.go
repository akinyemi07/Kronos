package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: kronos <lab-id> [lab-arguments...]")
		os.Exit(2)
	}

	labID := os.Args[1]

	switch labID {
	case "linux-timezone":
		args := append([]string{"scripts/run-lab.sh", labID}, os.Args[2:]...)
		cmd := exec.Command("bash", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				os.Exit(exitErr.ExitCode())
			}

			fmt.Fprintf(os.Stderr, "Error running lab: %v\n", err)
			os.Exit(1)
		}

	default:
		fmt.Fprintf(os.Stderr, "Error: unknown lab %q\n", labID)
		os.Exit(1)
	}
}
