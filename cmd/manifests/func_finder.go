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

var funcs map[PathFunc]*ast.FuncDecl

func init() {
	funcs = map[PathFunc]*ast.FuncDecl{}
}

// Traverse through the files, and find the functions
func GetFuncsByFiles(dct map[string]string) (map[PathFunc]*ast.FuncDecl, error) {
	// TODO: O(N^2), try to make it faster if needed.
	// Add some ignore file to ignore most of the non-.go files
	for _, fPath := range dct {
		set := token.NewFileSet()
		parsedFiles, err := parser.ParseFile(set, fPath, nil, parser.Mode(0))
		if err != nil {
			fmt.Println("Failed to parse package:", err)
			os.Exit(1)
		}

		populateFuncs(parsedFiles.Decls, fPath)
	}

	return funcs, nil
}

func populateFuncs(parsedFilesDecl []ast.Decl, fPath string) {
	for _, d := range parsedFilesDecl {
		collectFuncsFromParsedFileDecls(d, fPath)
	}
}

func collectFuncsFromParsedFileDecls(declaration ast.Decl, fPath string) {
	if fn, isFn := declaration.(*ast.FuncDecl); isFn {
		pf := PathFunc{
			PathToFile: fPath,
			FuncName:   fn.Name.Name,
		}

		funcs[pf] = fn
	}
}
