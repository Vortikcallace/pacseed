package exec

import (
	"fmt"
	"os/exec"
)

func ExecuteCommand(args ...string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("no command provided")
	}

	cmd := exec.Command(args[0], args[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("failed to execute %s: %w", args[0], err)
	}

	fmt.Printf("Output:\n%s", string(output))
	return string(output), nil
}
