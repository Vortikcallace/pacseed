package cli

import (
	"fmt"
	"log/slog"
	"os"
	"pacseed/internal/logger"

	"github.com/spf13/cobra"
)

var (
	debugMode bool
)

var rootCmd = &cobra.Command{
	Use:   "pacseed",
	Short: "",
	Long:  ``,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		level := "info"
		if debugMode {
			level = "debug"
		}
		logger.InitLogger(level, "text")
	},
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("welcome message, root command entered")
		fmt.Println("Welcome to E-pacseed! It's a tool for managing your Arch Linux packages.")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debugMode, "debug", false, "enable debug logging")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
