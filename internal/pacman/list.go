package pacman

import (
	"fmt"
	"strings"

	"pacseed/internal/exec"
)

func GetOfficialPackages() ([]string, error) {
	out, err := exec.ExecuteCommand("pacman", "-Qqn")
	if err != nil {
		return nil, fmt.Errorf("pacman -Qqn failed: %w", err)
	}
	return parseLines(out), nil
}

func parseLines(output string) []string {
	if strings.TrimSpace(output) == "" {
		return []string{}
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")
	var result []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}
	return result
}
