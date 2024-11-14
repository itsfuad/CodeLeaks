package scanner

import (
	"testing"
)

const fileName = "testfile.go"

func TestScanLine(t *testing.T) {
	tests := []struct {
		line     string
		lineNum  int
		filePath string
		expected bool
	}{
		{
			line:     `password := "supersecret"`,
			lineNum:  1,
			filePath: fileName,
			expected: true,
		},
		{
			line:     `apiKey = "myapikey"`,
			lineNum:  2,
			filePath: fileName,
			expected: true,
		},
		{
			line:     `username := "user"`,
			lineNum:  3,
			filePath: fileName,
			expected: false,
		},
		{
			line:     `// This is a comment`,
			lineNum:  4,
			filePath: fileName,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			// Reset secretVariableMap before each test
			secretVariableMap = make(map[string]SecretData)

			scanLine(tt.line, tt.lineNum, tt.filePath)

			if tt.expected {
				if len(secretVariableMap) == 0 {
					t.Errorf("Expected secret to be found in line: %s", tt.line)
				}
			} else {
				if len(secretVariableMap) > 0 {
					t.Errorf("Did not expect secret to be found in line: %s", tt.line)
				}
			}
		})
	}
}
