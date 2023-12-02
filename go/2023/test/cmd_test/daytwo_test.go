package cmd

import (
	"aoc23/cmd"
	"testing"
)

func TestDayTwoPartOne(t *testing.T) {
	tests := []struct {
		input    string
		expected uint
	}{
		{string("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"), 1},
		{string("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"), 2},
		{string("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"), 0},
		{string("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"), 0},
		{string("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"), 5},
	}
	var total uint
	var expected uint = 8
	for _, test := range tests {
		result := cmd.DayTwoPartOne(test.input)
		if result != test.expected {
			t.Errorf("For input %s | expected %v, but got %v", test.input, test.expected, result)
		}
		total += result

	}
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
func TestDayTwoPartTwo(t *testing.T) {
	tests := []struct {
		input    string
		expected uint
	}{
		{string("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"), 48},
		{string("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"), 12},
		{string("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"), 1560},
		{string("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"), 630},
		{string("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"), 36},
	}
	var total uint
	var expected uint = 2286
	for _, test := range tests {
		result := cmd.DayTwoPartTwo(test.input)
		if result != test.expected {
			t.Errorf("For input %s | expected %v, but got %v", test.input, test.expected, result)
		}
		total += result

	}
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
