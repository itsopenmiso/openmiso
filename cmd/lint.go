/*
Copyright © 2025 Michal Hodur @ itsopenmiso
*/
package cmd

import (
	"fmt"

	"github.com/itsopenmiso/openmiso/pkg/waypointfile"

	"github.com/spf13/cobra"
)

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("lint called")
		// FIXME: Lint only one file
		// Gets current directory
		jf := waypointfile.Synth()
		fmt.Println("Project name:", jf.Project)
		for a, app := range jf.Apps {
			fmt.Println(a, ". App name:", app.Name)
			fmt.Println(app.Name)
			for i, buildBlock := range app.BuildBlocks {
				fmt.Println("Building Block #", i)
				fmt.Println("Uses plugin")
				fmt.Println(buildBlock.Use.PluginName)
				// find a plugin

				// check if download flag is set

				// download plugin if necessary

				// run the plugin adding step config along with the variables
				// but mark those variables as redacted
				// secrets are also redacted
				// ask the plugin to perform action lint-build
				// It's plugin responsibility to handle it well
				fmt.Println()
				fmt.Println("Calling lint-build")
			}

			for i, deployBlock := range app.DeployBlocks {
				fmt.Println("Deploying Block #", i)
				fmt.Println("Uses plugin")
				fmt.Println(deployBlock.Use.PluginName)
				fmt.Println()

				// find a plugin

				// check if download flag is set

				// download plugin if necessary

				// run the plugin adding step config along with the variables
				// but mark those variables as redacted
				// secrets are also redacted
				// ask the plugin to perform action lint-deploy
				// It's plugin responsibility to handle it well
				fmt.Println("Calling lint-deploy")
			}

			for i, releaseBlock := range app.ReleaseBlocks {
				fmt.Println("Release Block #", i)
				fmt.Println("Uses plugin")
				fmt.Println(releaseBlock.Use.PluginName)
				fmt.Println()

				// find a plugin

				// check if download flag is set

				// download plugin if necessary

				// run the plugin adding step config along with the variables
				// but mark those variables as redacted
				// secrets are also redacted
				// ask the plugin to perform action lint-deploy
				// It's plugin responsibility to handle it well
				fmt.Println("Calling lint-release")
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(lintCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lintCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lintCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
