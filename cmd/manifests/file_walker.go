/*
Copyright © 2023 Alikhan Toleubay <alikhan.toleubay@gmail.com>
*/
package manifests

import (
	"fmt"
	"io/fs"
	"log"
	"slices"
	"path/filepath"
	"strings"
)

func ListAllGoFiles() (map[string]string, error) {
	log.SetPrefix("LOG:")
	log.Println("Starting Function")

	skipDirs := []string{".git"}
	dctOfGoFiles := map[string]string{}

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if info.IsDir() && slices.Contains(skipDirs, info.Name()) {
            // Skipping directories that configure in the array skipDirs
			// fmt.Printf("skipping a dir configured in the list: %v\n", skipDirs)

			return filepath.SkipDir
		} else if strings.LastIndex(info.Name(), ".go") == -1 {
            // Skipping non-.go files
			// fmt.Printf("skipping non-.go files, %q\n", info.Name())

			return nil
		}

        dctOfGoFiles[info.Name()] = path

		return nil
	})

	if err != nil {
		fmt.Printf("Some error occured %q", err)
	}

	return dctOfGoFiles, nil
}
