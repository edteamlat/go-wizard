package cmd

import (
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a package to the project.",
}

func init() {
	rootCmd.AddCommand(addCmd)
}
