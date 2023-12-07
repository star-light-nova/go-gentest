/*
Copyright © 2023 Alikhan Toleubay <alikhan.toleubay@gmail.com>
*/
package manifests

// This struct is a collection of the variables
// that changes the functionality of the command.
type FlagsValues struct {
	IsDryRun bool

	IsTestFolder bool
	TestFolder   string
}

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

2. Goes through all collected path of `.go` files and extracts functions

3. Prepares variables for the template

4. Generates `_test.go` files depending on the flag, migh just output to the terminal.
*/
func Start(flagsValues *FlagsValues) error {
	goFiles, err := ListAllGoFiles()
	check(err)

	pfuncs, err := GetFuncsByFiles(goFiles)
	check(err)

	templateVariables := GenerateVars(pfuncs)

	// Might panic
	GenerateTests(templateVariables, flagsValues)

	return nil
}
