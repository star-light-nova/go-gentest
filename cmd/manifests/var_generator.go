package manifests

import (
	"go/ast"
	"strings"
)

const (
	/*
	   The package name is considered to be the directory where the `.go` file exist.

	   EXAMPLE: path/to/file.go in this case the package name will be `to`
	   and the template will output `package to` for this test file.

	   If we can't go any further:

	   EXAMPLE: main.go (folder that has one file), this const DEFAULT_PACKAGE
	   is written as a package name
	*/
	DEFAULT_PACKAGE = "main"
)

// Holds all information for template.
type TemplateVars struct {
	PackageName  string
	FuncName     string
	FuncTestName string
}

/*
Parse the current path of the file
and returns the last folder as a package name.

If the length of the split path array is less than 2
the DEFAULT_PACKAGE const returned.
*/
func parsePackageNameFromPath(path string) string {
	packagePathArray := strings.Split(path, "/")

	if len(packagePathArray) >= 2 {
		return packagePathArray[len(packagePathArray)-2]
	}

	return DEFAULT_PACKAGE
}

// Groups information about every function in every file.
func GenerateVars(pfuncs map[PathFunc]*ast.FuncDecl) map[string][]TemplateVars {
	generatedVars := map[string][]TemplateVars{}

	for pf, _func := range pfuncs {
		myFunc := TemplateVars{
			PackageName:  parsePackageNameFromPath(pf.PathToFile),
			FuncName:     _func.Name.Name,
			FuncTestName: strings.Title(_func.Name.Name),
		}

		if _, exist := generatedVars[pf.PathToFile]; exist {
			generatedVars[pf.PathToFile] = append(generatedVars[pf.PathToFile], myFunc)
		} else {
			generatedVars[pf.PathToFile] = []TemplateVars{myFunc}
		}

		// TODO: Do something with return values/type and params values/types
	}

	return generatedVars
}
