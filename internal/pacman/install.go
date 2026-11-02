package pacman

import (
	"fmt"

	"pacseed/internal/exec"
)

func InstallOfficialPackage(name string) error {
	fmt.Printf("Installing official package: %s\n", name)
	_, err := exec.ExecuteCommand("sudo", "pacman", "-S", "--noconfirm", name)
	if err != nil {
		return fmt.Errorf("failed to install %s: %w", name, err)
	}
	return nil
}
