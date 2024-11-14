package scanner

import (
	"testing"
)

func TestParseIdentifier(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var myVar", "myVar"},
		{"var myVar = 123", ""},
		{" PASSWORD ", "PASSWORD"},
		{"1245pass", ""},
		{"_secret", "_secret"},
		{"pass123", "pass123"},
	}

	for _, test := range tests {
		result := parseIdentifier(test.input)
		if result != test.expected {
			t.Errorf("parseIdentifier(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}