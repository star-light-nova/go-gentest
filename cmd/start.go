package cmd

import (
	"github.com/spf13/cobra"
	"github.com/star-light-nova/gentest/cmd/manifests"
)

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().Bool("dry-run", false, "Only outputs the result to the terminal without any effect (no file creation).")
	startCmd.Flags().String("test-folder", "", "Generates tests inside of the 'test' folder.")
	startCmd.Flags().String("test-only", "", "Generates only one specified '.go' file.")
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: START_HELP_SHORT,
	Long:  START_HELP_LONG,
	Run: func(cmd *cobra.Command, args []string) {
		flagsValues := manifests.NewFlagsValues()

		flagsChecker(cmd, flagsValues)

		manifests.Start(flagsValues)
	},
}

func flagsChecker(cmd *cobra.Command, flagsValues *manifests.FlagsValues) {
	if isDryRun, err := cmd.Flags().GetBool("dry-run"); err != nil {
		panic("[--dry-run] Can't assign value: " + err.Error())
	} else {
		flagsValues.IsDryRun = isDryRun
	}

	if testFolder, err := cmd.Flags().GetString("test-folder"); err != nil {
		panic("[--test-folder] Can't assign value: " + err.Error())
	} else {
		if len(testFolder) != 0 {
			flagsValues.TestFolder = testFolder
			flagsValues.IsTestFolder = true
		}
	}

	if testOnlyFilePath, err := cmd.Flags().GetString("test-only"); err != nil {
		panic("[--test-only] Can't assign value: " + err.Error())
	} else {
		if len(testOnlyFilePath) != 0 {
			flagsValues.IsTestOnly = true
			flagsValues.TestOnlyFilePath = testOnlyFilePath
		}
	}
}
