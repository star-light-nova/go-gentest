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
	DEFAULT_PACKAGE = "main"
)

type TemplateVars struct {
	PackageName  string
	FuncName     string
	FuncTestName string
}

// Parse the current path of the file
// and returns the last folder as a package name.
// If the length of the split path array is less than 2
// the default package returned (main)
func parsePackageNameFromPath(path string) string {
	packagePathArray := strings.Split(path, "/")

	if len(packagePathArray) >= 2 {
		return packagePathArray[len(packagePathArray)-2]
	}

	return DEFAULT_PACKAGE
}

// Generates variables for the templates.
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

		// fmt.Println("Params\n", _func.Type.Params)
		// if _func.Type.Params != nil && _func.Type.Params.List != nil {
		// 	fmt.Println("DOES HAVE PARAMS")

		// 	fmt.Println("LIST", _func.Type.Params.List[0])
		// }
		// fmt.Println("Params", _func.Type.Params)
		// fmt.Println("Type Params", _func.Type.TypeParams)
		// fmt.Println("Returns", _func.Type.Results)
	}

	return generatedVars
}

type Temporary struct {
	PackageName string
	TV          *[]TemplateVars
}

var temp *template.Template

// Move it to another file?
func init() {
	temp = template.Must(templ.SimpleTemplate())
}

// Generates *_test.go files for the non _test.go files and ignored ones.
func GenerateTests(templateVariables map[string][]TemplateVars) {
	for ptf, tv := range templateVariables {
		if li := strings.LastIndex(ptf, "_test.go"); li != -1 {
			continue
		}

		pckname := tv[0].PackageName
		filePath := ptf[:len(ptf)-3] + "_test.go"

		f, err := os.Create(filePath)
		check(err)

		err = temp.Execute(f, Temporary{PackageName: pckname, TV: &tv})
		if err != nil {
			panic(err)
		}
	}

	fmt.Println()
}
