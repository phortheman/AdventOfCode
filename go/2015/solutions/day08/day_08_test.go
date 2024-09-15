package day08

import (
	"testing"
)

func TestDay8CountChars(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: `""`, expected: 0},
		{input: `"abc"`, expected: 3},
		{input: `"aaa\"aaa"`, expected: 7},
		{input: `"\x27"`, expected: 1},
		{input: `"\\"`, expected: 1},
	}
	for i, test := range tests {
		result := countChars(test.input)
		if result != test.expected {
			t.Errorf("Test %d - Expected %d and got %d: %s", i+1, test.expected, result, test.input)
		}
	}
}

func TestDay8Part1Example2(t *testing.T) {
	input := `""
"abc"
"aaa\"aaa"
"\x27"`
	expected := 12
	result, _ := Solver(input)
	if result != expected {
		t.Errorf("Expected %d and got %d: %s", expected, result, input)
	}
}

func TestEncodeString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: `""`, expected: `"\"\""`},
		{input: `"abc"`, expected: `"\"abc\""`},
		{input: `"aaa\"aaa"`, expected: `"\"aaa\\\"aaa\""`},
		{input: `"\x27"`, expected: `"\"\\x27\""`},
	}
	for i, test := range tests {
		result := encodeString(test.input)
		if result != test.expected {
			t.Errorf("Test %d - Expected %s and got %s: %s", i+1, test.expected, result, test.input)
		}
	}
}
func TestDay8Part2Example(t *testing.T) {
	input := `""
"abc"
"aaa\"aaa"
"\x27"`
	expected := 19
	_, result := Solver(input)
	if result != expected {
		t.Errorf("Expected %d and got %d: %s", expected, result, input)
	}
}
