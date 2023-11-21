/*
Copyright © 2023 Alikhan Toleubay <alikhan.toleubay@gmail.com>
*/
package manifests

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type PathFunc struct {
	PathToFile string
	FuncName   string
}

// Traverse through the files, and find the functions
func GetFuncsByFiles(dct map[string]string) (map[PathFunc]*ast.FuncDecl, error) {
	funcs := map[PathFunc]*ast.FuncDecl{}

	// Add someignore file to ignore most of the non-.go files
	for _, fPath := range dct {
		set := token.NewFileSet()
		parsedFiles, err := parser.ParseFile(set, fPath, nil, parser.Mode(0))
		if err != nil {
			fmt.Println("Failed to parse package:", err)
			os.Exit(1)
		}

		for _, d := range parsedFiles.Decls {
			if fn, isFn := d.(*ast.FuncDecl); isFn {
				pf := PathFunc{
					PathToFile: fPath,
					FuncName:   fn.Name.Name,
				}

				funcs[pf] = fn
			}
		}
	}

	return funcs, nil
}
