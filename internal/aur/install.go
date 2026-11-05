package aur

import (
	"fmt"
	"strings"

	"pacseed/internal/exec"
)

func (m *AurManager) Install(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("package name cannot be empty")
	}
	if strings.HasPrefix(name, "-") {
		return fmt.Errorf("invalid package name %s: names cannot start with a dash", name)
	}

	fmt.Printf("Installing AUR package: %s\n", name)

	if m.binPath == "" {
		return fmt.Errorf("cannot install %s: no AUR helper configured", name)
	}

	_, err := exec.ExecuteCommandRestore(m.binPath, "-S", "--noconfirm", "--needed", name)
	if err != nil {
		return fmt.Errorf("failed to install %s via %s: %w", name, m.binPath, err)
	}

	return nil
}
