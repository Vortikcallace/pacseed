package aur

import (
	"strings"

	"pacseed/internal/exec"
)

func GetAURPackages() ([]string, error) {
	out, err := exec.ExecuteCommand("pacman", "-Qqm")
	if err != nil {
		return []string{}, nil
	}
	return parseLines(out), nil
}

func parseLines(output string) []string {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	var result []string
	for _, line := range lines {
		if line != "" {
			result = append(result, line)
		}
	}
	return result
}
