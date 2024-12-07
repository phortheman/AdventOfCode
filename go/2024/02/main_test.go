package main

import (
	"strings"
	"testing"
)

func TestDay2IsSafe(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"7 6 4 2 1", true},
		{"1 2 7 8 9", false},
		{"9 7 6 2 1", false},
		{"1 3 2 4 5", false},
		{"8 6 4 4 1", false},
		{"1 3 6 7 9", true},
	}

	for idx, test := range tests {
		data := strings.Split(test.input, " ")
		if isSafe(data) != test.expected {
			t.Errorf("Expected test #%d to be %v. Input: %s", idx+1, test.expected, test.input)
		}
	}
}

func TestDay2IsSafeWithTolerance(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"7 6 4 2 1", true},
		{"1 2 7 8 9", false},
		{"9 7 6 2 1", false},
		{"1 3 2 4 5", true},
		{"8 6 4 4 1", true},
		{"1 3 6 7 9", true},
	}

	for idx, test := range tests {
		report := strings.Split(test.input, " ")
		if !isSafe(report) {
			safe := false
			for i := 0; i < len(report) && !safe; i++ {
				modifiedReport := removeIndex(i, report)
				if isSafe(modifiedReport) {
					safe = true
					if !test.expected {
						t.Errorf("Expected test #%d to be false but was true. Input: %s", idx+1, test.input)
					}
				}
			}
			if !safe && test.expected {
				t.Errorf("Expected test #%d to be true but it was false. Input: %s", idx+1, test.input)
			}
		}
	}
}

func TestDay2Part1Example(t *testing.T) {
	expected := 2
	result, _ := Solver(sampleInput)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}

func TestDay2Part2Example(t *testing.T) {
	expected := 4
	_, result := Solver(sampleInput)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}
