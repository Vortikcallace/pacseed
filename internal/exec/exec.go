package exec

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

func ExecuteCommandRestore(args ...string) (string, error) {
	slog.Debug("executing command...")
	if len(args) == 0 {
		return "", fmt.Errorf("no command provided")
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute %s: %w", args[0], err)
	}

	slog.Debug("command executed successfully")
	return "", nil
}

func ExecuteCommandCompile(args ...string) (string, error) {
	slog.Debug("executing command with output capture...")
	if len(args) == 0 {
		return "", fmt.Errorf("no command provided")
	}

	cmd := exec.Command(args[0], args[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("failed to execute %s: %w", args[0], err)
	}

	return string(output), nil
}
