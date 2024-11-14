package main

import (
	"testing"
)

func TestParseIdentifier(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"variable_name=some_value", "variable_name"},
		{"_underscore=123", "_underscore"},
		{"123invalid=456", ""},
		{"valid123=789", "valid123"},
		{"no_equals_sign", "no_equals_sign"},
		{"=no_identifier", ""},
		{"valid_identifier_123", "valid_identifier_123"},
	}

	for _, test := range tests {
		result := parseIdentifier(test.input)
		if result != test.expected {
			t.Errorf("parseIdentifier(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}
