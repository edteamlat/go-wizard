package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the project structure.",
	Long: `Initialize the project structure in the given architecture.

It generates the layers:
1. cmd
2. domain
3. infrastructure/handler
4. infrastructure/storage
5. model
6. sqlmigration

By default will use edhex if no arquitecture is specified.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called", cmd.Flag("architecture").Value)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}