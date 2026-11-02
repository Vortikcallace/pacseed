package cli

import (
	"fmt"
	"log/slog"

	aurPkg "pacseed/internal/aur"
	"pacseed/internal/config"
	"pacseed/internal/pacman"

	"github.com/spf13/cobra"
)

var (
	pathR string
	nameR string
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "help you to restore your repos",
	Long:  `help you to restore your repos`,
	Run: func(cmd *cobra.Command, args []string) {
		Restore()
	},
}

func Restore() {
	slog.Debug("Beginning restore")
	fmt.Println("restore...")

	if pathR == "" || nameR == "" {
		fmt.Println("Error: --path and --name are required for restore")
		return
	}

	config.ConfigPath = pathR
	config.ConfigName = nameR

	config.LoadConfig()

	if config.UConfig == nil {
		fmt.Println("restore failed: could not load config")
		return
	}

	for _, repo := range config.UConfig.Repos {
		if repo.AUR {
			continue
		}
		if err := pacman.InstallOfficialPackage(repo.Name); err != nil {
			slog.Error("Failed to install official package", "error", err)
			fmt.Println("restore failed")
			return
		}
	}

	if err := aurPkg.InstallAURPackages(); err != nil {
		slog.Error("Restore failed (AUR)", "error", err)
		fmt.Println("restore failed")
		return
	}

	fmt.Println("restore completed!")
	slog.Debug("Restore completed")
}

func init() {
	restoreCmd.Flags().StringVarP(&pathR, "path", "p", "", "path to config")
	restoreCmd.Flags().StringVarP(&nameR, "name", "n", "", "name of the config")
	rootCmd.AddCommand(restoreCmd)
}
