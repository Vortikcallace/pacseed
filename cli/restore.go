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

	config.LoadConfig(nameR, pathR)

	if config.UConfig == nil {
		fmt.Println("restore failed: could not load config")
		return
	}

	pManager := pacman.NewPacmanManager()
	var aManager *aurPkg.AurManager
	var err error

	for _, repo := range config.UConfig.Repos {
		if repo.AUR {
			if aManager == nil {
				aManager, err = aurPkg.NewAurManager()
				if err != nil {
					slog.Error("Failed to initialize AUR manager", "error", err)
					fmt.Println("restore failed")
					return
				}
			}

			if err := aManager.Install(repo.Name); err != nil {
				slog.Error("Failed to install AUR package", "error", err)
				fmt.Println("restore failed")
				return
			}
			continue
		}

		if err := pManager.Install(repo.Name); err != nil {
			slog.Error("Failed to install official package", "error", err)
			fmt.Println("restore failed")
			return
		}
	}

	fmt.Println("restore completed!")
	slog.Debug("Restore completed")
}

func init() {
	restoreCmd.Flags().StringVarP(&pathR, "path", "p", "", "path to config")
	restoreCmd.Flags().StringVarP(&nameR, "name", "n", "", "name of the config")

	_ = restoreCmd.MarkFlagRequired("path")
	_ = restoreCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(restoreCmd)
}
