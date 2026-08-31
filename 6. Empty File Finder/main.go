package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// go run main.go -d=testdata/
func main() {
	directory := flag.String("d", ".", "The directory to search for")
	flag.Parse()

	info, err := os.Stat(*directory)
	if err != nil {
		log.Fatal(err)
	}
	if !info.IsDir() {
		log.Fatalf("'%s' is not a directory", *directory)
	}

	emptyFiles, err := extractEmptyFiles(*directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range emptyFiles {
		fmt.Printf("The '%s' file is empty!\n", file)
	}

	fmt.Printf("\nFound %d empty file(s) in '%s'.\n", len(emptyFiles), *directory)
}

// extractEmptyFiles walks the whole tree under directory and collects the path
// of every zero-byte file. Directories it cannot read are reported and skipped
// instead of aborting the search.
func extractEmptyFiles(directory string) ([]string, error) {
	emptyFiles := make([]string, 0)

	err := filepath.WalkDir(directory, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Skipping '%s': %v\n", path, err)

			if entry != nil && entry.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		// Only regular files have a meaningful size: directories, symlinks and
		// devices are never "empty" in the sense we care about here.
		if !entry.Type().IsRegular() {
			return nil
		}

		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Skipping '%s': %v\n", path, err)
			return nil
		}

		if info.Size() == 0 {
			emptyFiles = append(emptyFiles, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return emptyFiles, nil
}
