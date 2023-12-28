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

// TODO: Maybe remove?
func check(err error) {
	if err != nil {
		panic(err)
	}
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
func Start(flagsValues *FlagsValues) error {
	var goFiles = map[string]string{}
	var err error

	if flagsValues.IsTestOnly {
		goFiles, err = FindAndCollectOneFile(flagsValues.TestOnlyFilePath)
	} else {
		goFiles, err = ListAllGoFiles()
	}

	pfuncs, err := GetFuncsByFiles(goFiles)
	check(err)

	templateVariables := GenerateVars(pfuncs)

	// Might panic
	// todo return an error
	GenerateTests(templateVariables, flagsValues)

	return nil
}
