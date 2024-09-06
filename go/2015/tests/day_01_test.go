package solutions_test

import (
	"testing"

	"github.com/phortheman/AdventOfCode_2015_go/solutions"
)

func TestDay1Part1Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "(())", expected: 0},
		{input: "()()", expected: 0},
		{input: "(((", expected: 3},
		{input: "(()(()(", expected: 3},
		{input: "))(((((", expected: 3},
		{input: "())", expected: -1},
		{input: ")))", expected: -3},
		{input: ")())())", expected: -3},
	}
	for i, test := range tests {
		result, _, err := solutions.Day01Solver(test.input)
		if err != nil {
			t.Errorf("solver error %v", err)
		}
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}

func TestDay1Part2Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: ")", expected: 1},
		{input: "()())", expected: 5},
	}
	for i, test := range tests {
		_, result, err := solutions.Day01Solver(test.input)
		if err != nil {
			t.Errorf("solver error %v", err)
		}
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}
