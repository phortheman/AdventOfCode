package solutions_test

import (
	"testing"

	"github.com/phortheman/AdventOfCode_2015_go/solutions"
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
		result, _ := solutions.Day02Solver(test.input)
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
		_, result := solutions.Day02Solver(test.input)
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}
