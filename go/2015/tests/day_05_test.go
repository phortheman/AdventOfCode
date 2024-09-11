package solutions_test

import (
	"testing"

	"github.com/phortheman/AdventOfCode_2015_go/solutions"
)

func TestDay5Part1Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "ugknbfddgicrmopn", expected: 1},
		{input: "aaa", expected: 1},
		{input: "jchzalrnumimnmhp", expected: 0},
		{input: "haegwjzuvuyypxyu", expected: 0},
		{input: "dvszwmarrgswjxmb", expected: 0},
	}
	for i, test := range tests {
		result, _ := solutions.Day05Solver(test.input)
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}

func TestDay5Part2Example(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "qjhvhtzxzqqjkmpb", expected: 1},
		{input: "xxyxx", expected: 1},
		{input: "uurcxstgmygtbstg", expected: 0},
		{input: "ieodomkazucvgmuy", expected: 0},
		{input: "aaaxyxy", expected: 1},
		{input: "aaaa", expected: 1},
		{input: "aaa", expected: 0},
	}
	for i, test := range tests {
		_, result := solutions.Day05Solver(test.input)
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}
