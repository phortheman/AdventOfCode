package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "abcdef", expected: 609043},
		{input: "pqrstuv", expected: 1048970},
	}
	for i, test := range tests {
		result, _ := Solver(test.input)
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}

// No sample was provided for part 2
