package cli

import (
	"fmt"
	"log/slog"
	"pacseed/internal/config"
	"strings"

	"github.com/spf13/cobra"
)

var (
	repos bool
	aur   bool
	repo  string
	pathC string
	nameC string
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "info about repos",
	Long:  `info about repos, you can check a specific repo by name or aur repos`,
	Run: func(cmd *cobra.Command, args []string) {
		if !repos && !aur && strings.TrimSpace(repo) == "" {
			fmt.Println("no repo specified, use --repos or --aur to list all repos, or --repo to check a specific repo")
			slog.Debug("no flags")
		}
		if repos {
			slog.Debug("repos flag set")
			config.PrintRepos()
		}
		if aur {
			fmt.Println("listing all AUR repos")
			slog.Debug("aur flag set")
			config.PrintAURRepos()
		}
		if strings.TrimSpace(repo) != "" {
			fmt.Println("checking repo " + repo)
			slog.Debug("repo flag set")
			config.PrintRepo(repo)
		}
	},
}

func init() {
	checkCmd.Flags().BoolVar(&repos, "repos", false, "list all repos")
	checkCmd.Flags().BoolVar(&aur, "aur", false, "list all AUR repos")
	checkCmd.Flags().StringVarP(&pathC, "path", "p", "", "path to config")
	checkCmd.Flags().StringVarP(&nameC, "name", "n", "", "name of the config")
	checkCmd.Flags().StringVar(&repo, "repo", "", "check a specific repo by name")
	rootCmd.AddCommand(checkCmd)
}
