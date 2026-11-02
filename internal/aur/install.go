package aur

import (
	"fmt"

	"pacseed/internal/config"
	"pacseed/internal/exec"
)

func InstallAURPackages() error {
	if config.UConfig == nil {
		return fmt.Errorf("config not loaded")
	}

	repos := config.UConfig.Repos
	if len(repos) == 0 {
		fmt.Println("Manifest is empty.")
		return nil
	}

	helper := GetHelper()
	hasAUR := false
	for _, r := range repos {
		if r.AUR {
			hasAUR = true
			break
		}
	}

	if hasAUR && helper == "" {
		var err error
		helper, err = InstallHelper()
		if err != nil {
			return fmt.Errorf("failed to setup AUR helper: %w", err)
		}
	}

	for _, repo := range repos {
		if repo.AUR {
			fmt.Printf("Installing AUR package: %s\n", repo.Name)
			if helper == "" {
				return fmt.Errorf("cannot install %s: no AUR helper", repo.Name)
			}
			_, err := exec.ExecuteCommand(helper, "-S", "--noconfirm", repo.Name)
			if err != nil {
				return fmt.Errorf("failed to install %s: %w", repo.Name, err)
			}
		}
	}
	return nil
}
