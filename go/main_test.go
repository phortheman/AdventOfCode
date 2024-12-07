package main

import (
	"testing"
)

/* Example function for more granular tests. Remove if not applicable
func TestMultiTestExample(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", -1},
	}
	for idx, test := range tests {
		result := funcToTest(test.input)
		if result != test.expected {
			t.Errorf("Expected test #%d to be %d but was %d. Input: %s", idx+1, test.expected, result, test.input)
		}
	}
} */

func TestPart1(t *testing.T) {
	expected := -1
	result, _ := Solver(string(sampleInput))
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := -1
	_, result := Solver(string(sampleInput))
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}
