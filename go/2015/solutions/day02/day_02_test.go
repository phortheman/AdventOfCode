package day02

import (
	"testing"
)

func TestDay2Part1Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "2x3x4", expected: 58},
		{input: "1x1x10", expected: 43},
	}
	for i, test := range tests {
		result, _ := Solver(test.input)
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}

func TestDay2Part2Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "2x3x4", expected: 34},
		{input: "1x1x10", expected: 14},
	}
	for i, test := range tests {
		_, result := Solver(test.input)
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}
