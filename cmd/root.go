/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"github.com/tgs266/neon/neon"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "neon",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		postgres, _ := cmd.Flags().GetString("postgres")
		username, _ := cmd.Flags().GetString("pg-username")
		password, _ := cmd.Flags().GetString("pg-password")
		port, _ := cmd.Flags().GetString("port")
		ui, _ := cmd.Flags().GetBool("ui")
		reset, _ := cmd.Flags().GetBool("reset")
		inCluster, _ := cmd.Flags().GetBool("in-cluster")
		kubePath, _ := cmd.Flags().GetString("kube-path")
		neon.Start(postgres, username, password, port, ui, reset, inCluster, kubePath)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

	rootCmd.Flags().String("postgres", "127.0.0.1:5432", "The address postgres binds to.")
	rootCmd.Flags().String("pg-username", "admin", "Postgres username.")
	rootCmd.Flags().String("pg-password", "admin", "Postgres password.")
	rootCmd.Flags().String("port", "5000", "The port the warehouse listens on.")
	rootCmd.Flags().Bool("ui", true, "host ui")
	rootCmd.Flags().Bool("reset", false, "reset database")
	rootCmd.Flags().Bool("in-cluster", false, "run in cluster mode or out of cluster")
	rootCmd.Flags().String("kube-path", "", "path to kube config")
}
