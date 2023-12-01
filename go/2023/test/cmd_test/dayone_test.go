package cmd

import (
	"aoc23/cmd"
	"testing"
)

func TestDayOnePartOne(t *testing.T) {
	tests := []struct {
		input    []byte
		expected uint32
	}{
		{[]byte("1abc2"), 12},
		{[]byte("pqr3stu8vwx"), 38},
		{[]byte("a1b2c3d4e5f"), 15},
		{[]byte("treb7uchet"), 77},
	}
	var total uint32
	var expected uint32 = 142
	for _, test := range tests {
		result := cmd.DayOnePartOne(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %d, but got %d", test.input, test.expected, result)
		}
		total += result
	}
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}

func TestDayOnePartTwo(t *testing.T) {
	input := []struct {
		input    []byte
		expected uint32
	}{
		{[]byte("two1nine"), 29},
		{[]byte("eightwothree"), 83},
		{[]byte("abcone2threexyz"), 13},
		{[]byte("xtwone3four"), 24},
		{[]byte("4nineeightseven2"), 42},
		{[]byte("zoneight234"), 14},
		{[]byte("7pqrstsixteen"), 76},
	}
	var total uint32
	var expected uint32 = 281
	for _, test := range input {
		result := cmd.DayOnePartTwo(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %d, but got %d", test.input, test.expected, result)
		}
		total += result
	}
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}

func TestGetSpelledOutValue(t *testing.T) {
	tests := []struct {
		input    []byte
		expected uint8
	}{
		{[]byte("one"), 1},
		{[]byte("two"), 2},
		{[]byte("three"), 3},
		{[]byte("four"), 4},
		{[]byte("five"), 5},
		{[]byte("six"), 6},
		{[]byte("seven"), 7},
		{[]byte("eight"), 8},
		{[]byte("nine"), 9},
		{[]byte("zero"), 0},
	}
	for _, test := range tests {
		result := cmd.CheckForSpelledOutDigit(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected %d, but got %d", test.input, test.expected, result)
		}
	}
}
