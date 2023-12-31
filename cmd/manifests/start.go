package manifests

// This struct is a collection of the variables
// that changes the functionality of the command.
type FlagsValues struct {
	IsDryRun bool

	IsTestFolder bool
	TestFolder   string

	IsTestOnly       bool
	TestOnlyFilePath string
}

func NewFlagsValues() *FlagsValues {
	return &FlagsValues{}
}

/*
What is going to be done:

1. Gather all `.go` files in the current directory with `[fileName]=path` format

2. Goes through all collected path of `.go` files and extracts functions, if `--test-only` flag used
will only goes through the specified `.go` file.

3. Prepares variables for the template

4. Generates `_test.go` files depending on the flag, migh just output to the terminal.
*/
func Start(flagsValues *FlagsValues) {
	var goFiles = map[string]string{}

	if flagsValues.IsTestOnly {
		goFiles = FindAndCollectOneFile(flagsValues.TestOnlyFilePath)
	} else {
		goFiles = ListAllGoFiles()
	}

	pfuncs := GetFuncsByFiles(goFiles)
	templateVariables := GenerateVars(pfuncs)

	GenerateTests(templateVariables, flagsValues)
}
