/*
Copyright © 2023 Alikhan Toleubay <alikhan.toleubay@gmail.com>
*/
package manifests

// This struct is a collection of the variables
// that changes the functionality of the command.
type FlagsValues struct {
	IsDryRun bool
	TestFolder string
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func NewFlagsValues() *FlagsValues {
    return &FlagsValues{}
}

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
