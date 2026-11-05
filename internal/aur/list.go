package aur

import (
	"fmt"
	"strings"

	"pacseed/internal/exec"
)

func GetAURPackages() ([]string, error) {
	out, err := exec.ExecuteCommandCompile("pacman", "-Qqm")
	if err != nil {
		return nil, fmt.Errorf("pacman -Qqm failed: %w", err)
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
