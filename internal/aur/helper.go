package aur

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pacseed/internal/exec"
)

type AurManager struct {
	binPath string
}

func NewAurManager() (*AurManager, error) {
	m := &AurManager{}

	helper := m.GetHelper()
	if helper == "" {
		var err error
		helper, err = m.InstallHelper()
		if err != nil {
			return nil, fmt.Errorf("failed to setup AUR helper: %w", err)
		}
	}

	m.binPath = helper
	return m, nil
}

func (m *AurManager) GetHelper() string {
	if _, err := exec.ExecuteCommand("which", "yay"); err == nil {
		return "yay"
	}
	if _, err := exec.ExecuteCommand("which", "paru"); err == nil {
		return "paru"
	}
	return ""
}

func (m *AurManager) InstallHelper() (string, error) {
	fmt.Println("No AUR helper found (yay or paru).\nChoose AUR helper to install:\n[1] yay\n[2] paru")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var helper string
	switch input {
	case "1":
		helper = "yay"
	case "2":
		helper = "paru"
	default:
		return "", fmt.Errorf("invalid choice")
	}

	fmt.Printf("Installing %s...\n", helper)

	tmpDir := fmt.Sprintf("/tmp/%s_build", helper)
	os.RemoveAll(tmpDir)

	_, err := exec.ExecuteCommand("git", "clone", fmt.Sprintf("https://aur.archlinux.org/%s.git", helper), tmpDir)
	if err != nil {
		return "", fmt.Errorf("failed to clone: %w", err)
	}

	_, err = exec.ExecuteCommand("sh", "-c", fmt.Sprintf("cd %s && makepkg -si --noconfirm", tmpDir))
	if err != nil {
		return "", fmt.Errorf("failed to build: %w", err)
	}

	return helper, nil
}
