package cmd

import (
	"github.com/spf13/cobra"
)

const (
	architectureFlag = "architecture"
	configPathFlag   = "config-path"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-wizard",
	Short: "Generator of packages with hexagonal architecture.",
	Long: `Generator of packages with hexagonal architecture.

Layers that can be generated:
1. domain
2. handler
3. storage
4. model
5. sqlmigration

Hexagonal Architecture variation that can be used:
1. edhex
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP(architectureFlag, "a", "edhex", "Indicates the architecture that will be use. By default it uses `edhex`.")
	rootCmd.PersistentFlags().StringP(configPathFlag, "c", "config.yaml", "Indicates the path of the config file. By default it uses `./config.yaml`.")
}
