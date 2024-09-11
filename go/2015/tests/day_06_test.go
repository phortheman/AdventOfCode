package solutions_test

import (
	"testing"

	"github.com/phortheman/AdventOfCode_2015_go/solutions"
)

func TestDay6Part1Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		changes  int
	}{
		{input: "turn on 0,0 through 999,999", expected: 1_000_000},
		{input: "toggle 0,0 through 999,0", expected: 1_000},
		{input: "turn off 499,499 through 500,500", expected: 0},
	}
	for i, test := range tests {
		result, _ := solutions.Day06Solver(test.input)
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}

func _TestDay6Part2Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "turn on 0,0 through 0,0", expected: 1},
		{input: "toggle 0,0 through 999,999", expected: 2_000_000},
	}
	for i, test := range tests {
		_, result := solutions.Day06Solver(test.input)
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}
