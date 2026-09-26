package cli

import (
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "help you to compile yaml config of your repos",
	Long:  `help you to compile yaml config of your repos`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)
}
