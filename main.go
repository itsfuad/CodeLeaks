package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "regexp"
    "strings"

	"codeleaks/utils"
)

// Map to store variables identified as secrets
var secretVariableMap = make(map[string]string)

var linePos = "%s:%d: "

var patterns = []*regexp.Regexp{
    regexp.MustCompile(`(?i)(aws_access_key_id|aws_secret_access_key|api_key|token|secret|password)[^\n]*`),
    regexp.MustCompile(`[A-Za-z0-9-_]{20,40}`),
    regexp.MustCompile(`[A-Fa-f0-9]{32,64}`),
    regexp.MustCompile(`[A-Za-z0-9-_]{32}\.[A-Za-z0-9-_]{6,}\.[A-Za-z0-9-_]{27,}`),
}

const entropyThreshold = 4.5

func parseIdentifier(str string) string {
	// If has '=', we take the left side
	if strings.Contains(str, "=") {
		str = strings.Split(str, "=")[0]
	}
	// First character must be a letter or an underscore, followed by any number of letters, digits, or underscores
	identifier := regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*`)
	return identifier.FindString(str)
}

func calculateEntropy(str string) float64 {
    freq := make(map[rune]float64)
    for _, char := range str {
        freq[char]++
    }
    entropy := 0.0
    for _, count := range freq {
        p := count / float64(len(str))
        entropy -= p * math.Log2(p)
    }
    return entropy
}

func isPotentialSecret(word string) bool {
    return len(word) > 8 && calculateEntropy(word) > entropyThreshold
}

// Store secrets in a map if identified as sensitive
func scanLine(line string, lineNum int, filePath string) {
	parts := strings.Split(line, "=")
	if len(parts) > 1 {
		variable := parseIdentifier(parts[0])
		value := strings.TrimSpace(parts[1])

		// Check regex patterns for secrets
		for _, pattern := range patterns {
			if pattern.MatchString(value) {
				secretVariableMap[variable] = value
				fmt.Print("Potential secret in ")
				utils.CYAN.Printf(linePos, filePath, lineNum)
				fmt.Println(utils.RED.Sprintf("%s", value))
				return
			}
		}

		// Check entropy if not matched by regex
		if isPotentialSecret(value) {
			secretVariableMap[variable] = value
			fmt.Print("High-entropy string in ")
			utils.GREY.Printf(linePos, filePath, lineNum)
			fmt.Println(utils.RED.Sprintf("%s", value))
		}
	}

	fmt.Printf("Found %d secrets in %s\n", len(secretVariableMap), filePath)
	fmt.Printf("Checking for secrets in %s\n", line)

	// Check if line references any known secret variables
	for varName := range secretVariableMap {
		fmt.Printf("Checking for '%s' in '%s'\n", varName, line)
		if strings.Contains(line, varName) {
			fmt.Print("Secret reference in ")
			utils.GREY.Printf(linePos, filePath, lineNum)
			fmt.Println(utils.BLUE.Sprintf("%s", line))
		}
	}
}

func scanFile(filePath string) {
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

func main() {
    files := []string{"tests/app.js", "tests/app2.js"} // Replace with your files
    for _, file := range files {
        scanFile(file)
    }
}