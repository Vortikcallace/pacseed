package cli

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	aurpkg "pacseed/internal/aur"
	"pacseed/internal/config"
	"pacseed/internal/config/models"
	"pacseed/internal/exec"

	"github.com/spf13/cobra"
)

var (
	path string
	name string
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "help you to compile yaml config of your repos",
	Long:  `help you to compile yaml config of your repos`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting compilation...")

		officialOut, err := exec.ExecuteCommand("pacman", "-Qqn")
		if err != nil {
			slog.Error("Failed to get official packages", "error", err)
			fmt.Println("compilation failed")
			return
		}
		officialPkgs := parseLines(officialOut)

		aurPkgs, err := aurpkg.GetAURPackages()
		if err != nil {
			slog.Error("Failed to get AUR packages", "error", err)
			fmt.Println("compilation failed")
			return
		}

		var repos []models.Repo
		for _, pkg := range officialPkgs {
			repos = append(repos, models.Repo{Name: pkg, AUR: false})
		}
		for _, pkg := range aurPkgs {
			repos = append(repos, models.Repo{Name: pkg, AUR: true})
		}

		config.UConfig = &models.Config{Repos: repos}

		config.ConfigPath = path
		config.ConfigName = name

		if err := config.SaveConfig(); err != nil {
			slog.Error("Failed to save config", "error", err)
			fmt.Println("compilation failed")
			return
		}

		actualPath := config.ConfigPath
		if actualPath == "" {
			actualPath, _ = os.UserHomeDir()
		}
		actualName := config.ConfigName
		if actualName == "" {
			actualName = "repos"
		}

		fullPath := filepath.Join(actualPath, actualName+".yaml")

		fmt.Printf("Successfully compiled %d packages\n", len(repos))
		fmt.Printf("Saved to: %s\n", fullPath)
	},
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

func init() {
	compileCmd.Flags().StringVarP(&path, "path", "p", "", "path to save the compiled config")
	compileCmd.Flags().StringVarP(&name, "name", "n", "", "name of the config")
	rootCmd.AddCommand(compileCmd)
}
