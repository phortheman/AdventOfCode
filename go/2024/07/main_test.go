package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 3749
	result, _ := Solver(sampleInput)
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 11387
	_, result := Solver(sampleInput)
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestGranularPartTwo(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"156: 15 6", 156},
		{"7290: 6 8 6 15", 7290},
		{"192: 17 8 14", 192},
	}
	for idx, test := range tests {
		_, result := Solver(test.input)
		if result != test.expected {
			t.Errorf("Expected test #%d to be %d but was %d. Input: %s", idx+1, test.expected, result, test.input)
		}
	}
}
