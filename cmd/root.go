/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
	Short: "Generador de paquetes con Arquitectura Hexagonal.",
	Long: `Generador de paquetes con Arquitectura Hexagonal.

Las capas que genera son:
1. domain
2. handler
3. storage
4. model
5. sqlmigration`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
