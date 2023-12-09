/*
Copyright © 2023 Alikhan Toleubay <alikhan.toleubay@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/star-light-nova/gentest/cmd/manifests"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		flagsValues := manifests.NewFlagsValues()

		if isDryRun, err := cmd.Flags().GetBool("dry-run"); err != nil {
			panic("Something wrong with flags. [DRY-RUN]")
		} else {
			flagsValues.IsDryRun = isDryRun
		}

		if testFolder, err := cmd.Flags().GetString("test-folder"); err != nil {
			panic("Something wrong with flags [TEST-FOLDER]")
		} else {
			if len(testFolder) != 0 {
				flagsValues.TestFolder = testFolder
				flagsValues.IsTestFolder = true
			}
		}

		if testOnlyFolder, err := cmd.Flags().GetString("test-only"); err != nil {
			panic("Seomthing wrong with flag [TEST-ONLY]")
		} else {
			flagsValues.IsTestOnly = true
			flagsValues.TestOnlyFolder = testOnlyFolder
		}

		err := manifests.Start(flagsValues)

		if err != nil {
			panic("Something went wrong")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")
	startCmd.Flags().Bool("dry-run", false, "Only outputs the result to the terminal without any effect (no file creation).")
	startCmd.Flags().String("test-folder", "", "Generates tests inside of the 'test' folder.")
	startCmd.Flags().String("test-only", "", "Generates only one specified `.go` file.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
