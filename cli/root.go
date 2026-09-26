package cli

import (
	"fmt"
	"os"
	"pacseed/internal/config"
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
	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogger("info", "text")
		config.LoadConfig()
		fmt.Println("Welcome to pacseed! It's a tool for managing your Arch Linux packages.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
