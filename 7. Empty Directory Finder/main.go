package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// go run main.go -d=testdata -nested=true
func main() {
	directory := flag.String("d", ".", "Directory to search")
	nested := flag.Bool("nested", false, "Also report directories that hold nothing but empty directories")
	flag.Parse()

	info, err := os.Stat(*directory)
	if err != nil {
		log.Fatal(err)
	}
	if !info.IsDir() {
		log.Fatalf("'%s' is not a directory", *directory)
	}

	emptyDirs := make([]string, 0)
	extractEmptyDirs(*directory, *nested, &emptyDirs)

	for _, dir := range emptyDirs {
		fmt.Printf("The '%s' directory is empty!\n", dir)
	}

	fmt.Printf("\nFound %d empty director(ies) in '%s'.\n", len(emptyDirs), *directory)
}

// extractEmptyDirs walks the tree under directory and appends every empty
// directory it finds to emptyDirs. It reports whether the subtree holds at
// least one file, which is what lets the -nested mode spot directories that
// contain only other empty directories.
func extractEmptyDirs(directory string, nested bool, emptyDirs *[]string) bool {
	entries, err := os.ReadDir(directory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Skipping '%s': %v\n", directory, err)

		// A directory we cannot read may well hold files, so never call it empty.
		return true
	}

	if len(entries) == 0 {
		*emptyDirs = append(*emptyDirs, directory)
		return false
	}

	hasFiles := false
	for _, entry := range entries {
		if !entry.IsDir() {
			hasFiles = true
			continue
		}

		if extractEmptyDirs(filepath.Join(directory, entry.Name()), nested, emptyDirs) {
			hasFiles = true
		}
	}

	if nested && !hasFiles {
		*emptyDirs = append(*emptyDirs, directory)
	}

	return hasFiles
}
