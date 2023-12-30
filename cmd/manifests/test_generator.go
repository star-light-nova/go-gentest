package manifests

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	templ "github.com/star-light-nova/gentest/cmd/manifests/templates"
)

// Template is going to be created using these values.
type Template struct {
	PackageName string
	TV          *[]TemplateVars
}

var temp *template.Template

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

	for ptf, tv := range templateVariables {
		if li := strings.LastIndex(ptf, "_test.go"); li != -1 {
			continue
		}

		pckname := tv[0].PackageName
		filePath := ptf[:len(ptf)-3] + "_test.go"

		// Add `test/` folder as a parent folder.
		if flagsValues.IsTestFolder {
			filePath = flagsValues.TestFolder + "/" + filePath
		}

		if flagsValues.IsDryRun {
			dryRun(pckname, filePath, &tv)
		} else {
			if flagsValues.IsTestFolder {
				err = os.MkdirAll(filePath[:strings.LastIndex(filePath, "/")], os.ModePerm)
				check(err)
			}

			realRun(pckname, filePath, &tv, f)
		}
	}

	// Clsoe the file IO
	defer f.Close()
}

func dryRun(pckname, filePath string, tv *[]TemplateVars) {
	fmt.Println("====================")
	fmt.Printf("PACKAGE NAME: %s, PATH: %s\n\n", pckname, filePath)
	fmt.Println("== BEGIN TEMPLATE ==")

	executeTemplate(os.Stdout, pckname, tv)

	fmt.Println("== END TEMPLATE ==\n")
}

func realRun(pckname, filePath string, tv *[]TemplateVars, f *os.File) {
	var err error

	f, err = os.Create(filePath)
	check(err)

	executeTemplate(f, pckname, tv)
}

// TODO: Description
func executeTemplate(writer *os.File, pckname string, tv *[]TemplateVars) {
	err := temp.Execute(writer, Template{PackageName: pckname, TV: tv})
	if err != nil {
		panic(err)
	}
}
