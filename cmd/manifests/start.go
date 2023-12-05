/*
Copyright © 2023 Alikhan Toleubay <alikhan.toleubay@gmail.com>
*/
package manifests

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Start(isDryRun bool, testFolder string) error {
	goFiles, err := ListAllGoFiles()
	check(err)

	pfuncs, err := GetFuncsByFiles(goFiles)
	check(err)

	templateVariables := GenerateVars(pfuncs)

	// Might panic
	GenerateTests(templateVariables, isDryRun, testFolder)

	return nil
}
