package scanner

import (
	"bufio"
	"codeleaks/utils"
	"fmt"
	"os"
	"strings"
)

// Check if a line references any known secrets
func checkReferences(line string, lineNum int, filePath string) {
	for varName, data := range secretVariableMap {
		if strings.Contains(line, varName) && filePath != data.FilePath {
			utils.YELLOW.Printf("Potential secret reference '%s' in ", varName)
			utils.CYAN.Printf(linePos, filePath, lineNum)
			fmt.Println(utils.BLUE.Sprintf("%s (originally from %s:%d)", varName, data.FilePath, data.LineNum))
		}
	}
}

// Scan a line of code and detect secrets
func scanLine(line string, lineNum int, filePath string) {
	// Check for string literals (either "..." or '...')
	parts := strings.Split(line, "=")
	if len(parts) > 1 {
		// Only parse the identifier if the value looks like a string
		value := strings.Trim(strings.TrimSpace(parts[1]), `;`)

		// If the value is a string literal, then extract the identifier
		if strings.HasPrefix(value, `"`) || strings.HasPrefix(value, `'`) {
			variable := parseIdentifier(parts[0])
			// Check if the value matches any of the patterns (secrets)
			for _, pattern := range patterns {
				if pattern.MatchString(value) {
					secretVariableMap[variable] = SecretData{Value: value, FilePath: filePath, LineNum: lineNum}
					utils.PURPLE.Print("Potential secret in ")
					utils.CYAN.Printf(linePos, filePath, lineNum)
					fmt.Println(utils.RED.Sprintf("%s", value))
					return
				}
			}

			// Check for high-entropy strings if no regex match
			if isPotentialSecret(value) {
				secretVariableMap[variable] = SecretData{Value: value, FilePath: filePath, LineNum: lineNum}
				utils.PURPLE.Print("High-entropy string in ")
				utils.GREY.Printf(linePos, filePath, lineNum)
				fmt.Println(utils.RED.Sprintf("%s", value))
			}
		}
	}

	// Check for references to known secret variables
	checkReferences(line, lineNum, filePath)
}

func scanFile(filePath string) {
	utils.GREEN.Printf("Scanning %s...\n", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, "//") && !strings.HasPrefix(line, "#") {
			scanLine(line, lineNum, filePath)
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}


func ScanFiles(files []string) {
	for _, file := range files {
		scanFile(file)
	}

	utils.GREEN.Printf("\nScanning complete\n")

	if len(secretVariableMap) == 0 {
		utils.GREEN.Println("No potential secrets found.")
		return
	}
	fmt.Printf("\nFound %d potential secrets\n", len(secretVariableMap))
}