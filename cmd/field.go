package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fieldCmd represents the field command
var fieldCmd = &cobra.Command{
	Use:   "field",
	Short: "Adds a field to a package",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Command not supported yet")
	},
}

func init() {
	fieldCmd.Flags().StringP("package", "p", "", "Indicates the package to add the field.")

	addCmd.AddCommand(fieldCmd)
}
