package manifests

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var SKIP_DIRS = []string{".git"}

// Walks in the current directory to find `.go` files. [fileName]=path
func walkCurrentDirectory(goFiles map[string]string) error {
	walker := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("[Walking] Prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if info.IsDir() && slices.Contains(SKIP_DIRS, info.Name()) {
			// Skipping directories that configure in the array SKIP_DIRS
			// fmt.Printf("skipping a dir configured in the list: %v\n", SKIP_DIRS)
			return filepath.SkipDir
		} else if strings.LastIndex(info.Name(), ".go") == -1 {
			// Skipping non-.go files
			// fmt.Printf("skipping non-.go files, %q\n", info.Name())
			return nil
		}

		goFiles[info.Name()] = path

		return nil
	}

	return filepath.Walk(".", walker)
}

func ListAllGoFiles() map[string]string {
	dctOfGoFiles := map[string]string{}

	err := walkCurrentDirectory(dctOfGoFiles)
	if err != nil {
		panic("[Listing] Walked through current directory, but got an error: " + err.Error())
	}

	return dctOfGoFiles
}

// Finds the file specified by `--test-only` flag
func FindAndCollectOneFile(path string) map[string]string {
	if file, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic("[--test-only] Didn't find the file in this path: " + path + " \nError: " + err.Error())
	} else {
		return map[string]string{file.Name(): path}
	}
}
