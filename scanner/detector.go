package scanner

import (
	"math"
	"regexp"
	"strings"
)

type SecretData struct {
	Value    string
	FilePath string
	LineNum  int
}

var secretVariableMap = make(map[string]SecretData)
var linePos = "%s:%d: "

var patterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)(aws_access_key_id|aws_secret_access_key|api_key|token|secret|password)[^\n]*`),
	regexp.MustCompile(`[A-Za-z0-9-_]{20,40}`),
	regexp.MustCompile(`[A-Fa-f0-9]{32,64}`),
	regexp.MustCompile(`[A-Za-z0-9-_]{32}\.[A-Za-z0-9-_]{6,}\.[A-Za-z0-9-_]{27,}`),
}

const entropyThreshold = 4.5

// Regular expression to match valid identifiers
var validIdentifier = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

func parseIdentifier(str string) string {
	// Split by spaces and return the last element
	str = strings.TrimSpace(str)
	parts := strings.Split(str, " ")
	identifier := parts[len(parts)-1]
	// Check if the identifier is a valid identifier
	if validIdentifier.MatchString(identifier) {
		return identifier
	}
	return ""
}

// Calculate entropy to detect high-entropy strings (potential secrets)
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

// Check if the string has high entropy (indicating potential secrets)
func isPotentialSecret(word string) bool {
	return len(word) > 8 && calculateEntropy(word) > entropyThreshold
}