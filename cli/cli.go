package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ParseCLI parses the command-line flags and arguments
func ParseCLI() ([]string, []string, error) {
	// Define CLI flags
	dirname := flag.String("d", "", "Directory to scan (required)")
	onlyScanFiles := flag.String("o", "", "Comma-separated list of files to scan (e.g., file1.txt,file2.txt)")
	onlyScanExts := flag.String("x", "", "Comma-separated list of extensions to scan (e.g., .go,.py)")
	excludeFiles := flag.String("e", "", "Comma-separated list of files to exclude (e.g., file1.txt,file2.txt)")
	excludeExts := flag.String("ex", "", "Comma-separated list of extensions to exclude (e.g., .go,.py)")
	help := flag.Bool("h", false, "Show usage information")
	flag.Parse()

	if *help {
		flag.Usage() // Show help message
		os.Exit(0)
	}

	// Validate that the directory is specified
	if *dirname == "" {
		return nil, nil, fmt.Errorf("the -d flag (directory) is required")
	}

	// Parse exclude extensions if provided
	var exclude []string
	if *excludeExts != "" {
		exclude = strings.Split(*excludeExts, ",")
		// Trim spaces from the extensions
		for i, ext := range exclude {
			exclude[i] = strings.TrimSpace(ext)
		}
	}

	// Parse include files if provided
	var filesToScan []string
	if *onlyScanFiles != "" {
		filesToScan = strings.Split(*onlyScanFiles, ",")
		// Trim spaces from the filenames
		for i, file := range filesToScan {
			filesToScan[i] = strings.TrimSpace(file)
		}
	}

	// Parse include extensions if provided
	var extToScan []string
	if *onlyScanExts != "" {
		extToScan = strings.Split(*onlyScanExts, ",")
		// Trim spaces from the extensions
		for i, ext := range extToScan {
			extToScan[i] = strings.TrimSpace(ext)
		}
	}

	// Find all files in the directory
	var files []string
	err := addToFiles(&files, *dirname, exclude, filesToScan, extToScan, excludeFiles)
	if err != nil {
		return nil, nil, err
	}

	return files, exclude, nil
}

func checkExclusion(path string, excludeFiles *string) error {
	if *excludeFiles != "" {
		filesToExclude := strings.Split(*excludeFiles, ",")
		for _, excludeFile := range filesToExclude {
			if strings.Contains(path, excludeFile) {
				return nil
			}
		}
	}
	return nil
}

func checkFile(path string, files *[]string, filesToScan []string, extToScan []string, ext string) error {
	if len(filesToScan) > 0 {
		for _, file := range filesToScan {
			if strings.Contains(path, file) {
				*files = append(*files, path)
				return nil
			}
		}
	} else if len(extToScan) > 0 {
		for _, ext2s := range extToScan {
			if ext2s == ext {
				*files = append(*files, path)
				return nil
			}
		}
	} else {
		// Add the file to the list if it matches no exclusion
		*files = append(*files, path)
	}
	return nil
}

func addToFiles(files *[]string, dirname string, exclude []string, filesToScan []string, extToScan []string, excludeFiles *string) error {
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if file should be excluded based on filename
		err = checkExclusion(path, excludeFiles)
		if err != nil {
			return nil
		}

		// Check if file has an extension that needs to be excluded
		ext := filepath.Ext(path)
		for _, excludeExt := range exclude {
			if ext == excludeExt {
				return nil
			}
		}

		// If only files are specified, check if the file should be scanned
		err = checkFile(path, files, filesToScan, extToScan, ext)
		if err != nil {
			return nil
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking the directory: %w", err)
	}
	return nil
}