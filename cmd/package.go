package cmd

import (
	"github.com/edteamlat/go-wizard/domain/runner"
	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Adds a package to the project.",
	Run: func(cmd *cobra.Command, args []string) {
		run(cmd, args, runner.NewPackage)
	},
}

func init() {
	addCmd.AddCommand(packageCmd)
}
