package main

import (
	"flag"
	"fmt"

	"codeleaks/cli"
	"codeleaks/scanner"
)

// main is the entry point for the program
func main() {
	// Parse CLI flags
	files, exclude, err := cli.ParseCLI()
	if err != nil {
		// If an error occurs, print the error and show usage
		fmt.Println("Error:", err)
		flag.Usage() // Show help message
		return
	}

	// Print the list of files and extensions to exclude
	fmt.Println("Files to scan:", files)
	fmt.Println("Excluded extensions:", exclude)

	// Call the scanner to scan the files
	scanner.ScanFiles(files)
}
