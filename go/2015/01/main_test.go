package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
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
		result, _ := Solver([]byte(test.input))
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: ")", expected: 1},
		{input: "()())", expected: 5},
	}
	for i, test := range tests {
		_, result := Solver([]byte(test.input))
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}
