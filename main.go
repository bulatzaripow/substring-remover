package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	var dirPath, substr string

	// getting flags
	flag.StringVar(&dirPath, "dir", "", "Path to the root directory")
	flag.StringVar(&substr, "substr", "", "Substring to match")
	flag.Parse()

	if dirPath == "" || substr == "" {
		fmt.Println("Usage: go run main.go -dir=<directory> -substr=<substring_to_remove>")
		os.Exit(1)
	}

	absPath, err := filepath.Abs(dirPath)
	if err != nil {
		fmt.Printf("Invalid directory path: %v\n", err)
		os.Exit(1)
	}

	// collect all paths
	var paths []string
	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		paths = append(paths, path)
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	// sorting paths desc by deep
	sort.Slice(paths, func(i, j int) bool {
		return strings.Count(paths[i], string(filepath.Separator)) > strings.Count(paths[j], string(filepath.Separator))
	})

	// renaming
	for _, path := range paths {
		base := filepath.Base(path)
		if !strings.Contains(base, substr) {
			continue
		}

		newBase := strings.ReplaceAll(base, substr, "")
		if newBase == "" {
			fmt.Printf("Skipping %s\n", path)
			continue
		}

		dir := filepath.Dir(path)
		newPath := filepath.Join(dir, newBase)

		if _, err := os.Stat(newPath); err == nil {
			fmt.Printf("Skipping target already exists:  %s -> %s\n", path, newPath)
			continue
		}

		fmt.Printf("Renaming %s to %s\n", path, newPath)
		if err := os.Rename(path, newPath); err != nil {
			fmt.Printf("Error renaming %s to %s: %v\n", path, newPath, err)
		}
	}
}
