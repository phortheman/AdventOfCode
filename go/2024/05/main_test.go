package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 143
	result, _ := Solver(sampleInput)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 123
	_, result := Solver(sampleInput)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}

func TestGetMiddle(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"75", "47", "61", "53", "29"}, 61},
		{[]string{"97", "61", "53", "29", "13"}, 53},
		{[]string{"75", "29", "13"}, 29},
		{[]string{"75", "97", "47", "61", "53"}, 47},
		{[]string{"61", "13", "29"}, 13},
		{[]string{"97", "13", "75", "29", "47"}, 75},
	}

	for idx, test := range tests {
		result := getMiddle(test.input)
		if result != test.expected {
			t.Errorf("Test #%d expected %d but got %d: %v", idx+1, test.expected, result, test.input)
		}
	}
}
