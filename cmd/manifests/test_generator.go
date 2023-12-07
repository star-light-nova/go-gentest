/*
Copyright © 2023 Alikhan Toleubay <alikhan.toleubay@gmail.com>
*/
package manifests

import (
	"fmt"
	"go/ast"
	"os"
	"strings"
	"text/template"

	templ "github.com/star-light-nova/gentest/cmd/manifests/templates"
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

// Template is going to be created using these values.
type Temporary struct {
	PackageName string
	TV          *[]TemplateVars
}

var temp *template.Template

// Move it to another file? (func#init, functGenerateTests)
func init() {
	temp = template.Must(templ.SimpleTemplate())
}

// Generates `*_test.go` files for the non `_test.go` files and ignored ones.
func GenerateTests(templateVariables map[string][]TemplateVars, flagsValues *FlagsValues) {
	// This one is here to aviod `os.Create` and `f.Defer` being in the loop.
	f, err := os.Create(os.DevNull)
	if err != nil {
		fmt.Println("Error /dev/null")
		panic(err)
	}

	// TODO: Refactor
	for ptf, tv := range templateVariables {
		if li := strings.LastIndex(ptf, "_test.go"); li != -1 {
			continue
		}

		pckname := tv[0].PackageName
		filePath := ptf[:len(ptf)-3] + "_test.go"

		// Add test/ folder as a parent folder.
		if flagsValues.IsTestFolder {
			filePath = flagsValues.TestFolder + "/" + filePath
		}

		if flagsValues.IsDryRun {
			fmt.Println("====================")
			fmt.Printf("PACKAGE NAME: %s, PATH: %s\n\n", pckname, filePath)
			fmt.Println("== BEGIN TEMPLATE ==")
			err := temp.Execute(os.Stdout, Temporary{PackageName: pckname, TV: &tv})
			if err != nil {
				panic(err)
			}
			fmt.Println("== END TEMPLATE ==\n")
		} else {
            if flagsValues.IsTestFolder {
                err = os.MkdirAll(filePath[:strings.LastIndex(filePath, "/")], os.ModePerm)
                check(err)
            }

			f, err = os.Create(filePath)
			check(err)

			err = temp.Execute(f, Temporary{PackageName: pckname, TV: &tv})
			if err != nil {
				panic(err)
			}

		}
	}

	// Clsoe the writing IO
	defer f.Close()
}
