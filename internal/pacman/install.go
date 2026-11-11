package pacman

import (
	"fmt"
	"strings"

	"pacseed/internal/exec"
)

type PacmanManager struct {
	binPath string
}

func NewPacmanManager() *PacmanManager {
	return &PacmanManager{
		binPath: "pacman",
	}
}

func (m *PacmanManager) Install(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("package name cannot be empty")
	}
	if strings.HasPrefix(name, "-") {
		return fmt.Errorf("invalid package name %s: names cannot start with a dash", name)
	}

	fmt.Printf("Installing official package: %s\n", name)

	_, err := exec.ExecuteCommandRestore("sudo", m.binPath, "-S", "--noconfirm", "--needed", name)
	if err != nil {
		return fmt.Errorf("failed to install %s via pacman: %w", name, err)
	}

	return nil
}
