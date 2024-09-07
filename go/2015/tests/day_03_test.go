package solutions_test

import (
	"testing"

	"github.com/phortheman/AdventOfCode_2015_go/solutions"
)

func TestDay3Part1Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: ">", expected: 2},
		{input: "^>v<", expected: 4},
		{input: "^v^v^v^v^v", expected: 2},
	}
	for i, test := range tests {
		result, _ := solutions.Day03Solver([]byte(test.input))
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}

func TestDay3Part2Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "^v", expected: 3},
		{input: "^>v<", expected: 3},
		{input: "^v^v^v^v^v", expected: 11},
	}
	for i, test := range tests {
		_, result := solutions.Day03Solver([]byte(test.input))
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}
