package cmd

import (
	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := cmd.Flag(configPathFlag)
		architecture := cmd.Flag(architectureFlag)

		run(configPath.Value.String(), architecture.Value.String(), "")
	},
}

func init() {
	addCmd.AddCommand(packageCmd)
}
