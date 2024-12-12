package main

import (
	file "aoc23/internal"
	"testing"
)

func TestPartOne(t *testing.T) {
	content := file.Read_String_Into_Byte_Slice(EXAMPLE)
	tiles, start := MakeGraph(content)
	expectedStart := Point{0, 2}
	if start != expectedStart {
		t.Errorf("Unexpected start position. Expected %v but got %v", expectedStart, start)
		return
	}
	total := Part1Solver(tiles, start)
	var expected int = 8
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}

func TestTranslateStartPipe(t *testing.T) {
	tests := []struct {
		north    byte
		south    byte
		east     byte
		west     byte
		expected byte
	}{
		// North and South = |
		{north: '|', south: '|', east: 'L', west: '7', expected: '|'},
		{north: '|', south: 'L', east: 'L', west: '7', expected: '|'},
		{north: '|', south: 'J', east: 'L', west: '7', expected: '|'},
		{north: '7', south: '|', east: 'L', west: '7', expected: '|'},
		{north: '7', south: 'L', east: 'L', west: '7', expected: '|'},
		{north: '7', south: 'J', east: 'L', west: '7', expected: '|'},
		{north: 'F', south: '|', east: 'L', west: '7', expected: '|'},
		{north: 'F', south: 'L', east: 'L', west: '7', expected: '|'},
		{north: 'F', south: 'J', east: 'L', west: '7', expected: '|'},

		// East and West = -
		{north: 'L', south: 'F', east: '-', west: 'L', expected: '-'},
		{north: 'L', south: 'F', east: '-', west: 'F', expected: '-'},
		{north: 'L', south: 'F', east: 'J', west: 'L', expected: '-'},
		{north: 'L', south: 'F', east: 'J', west: 'F', expected: '-'},
		{north: 'L', south: 'F', east: '7', west: 'L', expected: '-'},
		{north: 'L', south: 'F', east: '7', west: 'F', expected: '-'},

		// North and East = L
		{north: '|', south: '7', east: '7', west: 'J', expected: 'L'},
		{north: '|', south: '7', east: 'J', west: 'J', expected: 'L'},
		{north: '|', south: '7', east: '-', west: 'J', expected: 'L'},
		{north: '7', south: '7', east: '7', west: 'J', expected: 'L'},
		{north: '7', south: '7', east: 'J', west: 'J', expected: 'L'},
		{north: '7', south: '7', east: '-', west: 'J', expected: 'L'},
		{north: 'F', south: '7', east: '7', west: 'J', expected: 'L'},
		{north: 'F', south: '7', east: 'J', west: 'J', expected: 'L'},
		{north: 'F', south: '7', east: '-', west: 'J', expected: 'L'},

		// North and West = J
		{north: '|', south: '7', east: 'L', west: 'L', expected: 'J'},
		{north: '|', south: '7', east: 'L', west: 'F', expected: 'J'},
		{north: '7', south: '7', east: 'L', west: 'L', expected: 'J'},
		{north: '7', south: '7', east: 'L', west: 'F', expected: 'J'},
		{north: 'F', south: '7', east: 'L', west: 'L', expected: 'J'},
		{north: 'F', south: '7', east: 'L', west: 'F', expected: 'J'},

		// South and West = 7
		{north: 'L', south: '|', east: '|', west: 'L', expected: '7'},
		{north: 'L', south: '|', east: '|', west: 'F', expected: '7'},
		{north: 'L', south: 'L', east: '|', west: 'L', expected: '7'},
		{north: 'L', south: 'L', east: '|', west: 'F', expected: '7'},
		{north: 'L', south: 'J', east: '|', west: 'L', expected: '7'},
		{north: 'L', south: 'J', east: '|', west: 'F', expected: '7'},

		// South and East = F
		{north: 'L', south: '|', east: '7', west: '|', expected: 'F'},
		{north: 'L', south: '|', east: '-', west: '|', expected: 'F'},
		{north: 'L', south: '|', east: 'J', west: '|', expected: 'F'},
		{north: 'L', south: 'L', east: '7', west: '|', expected: 'F'},
		{north: 'L', south: 'L', east: '-', west: '|', expected: 'F'},
		{north: 'L', south: 'L', east: 'J', west: '|', expected: 'F'},
		{north: 'L', south: 'J', east: '7', west: '|', expected: 'F'},
		{north: 'L', south: 'J', east: '-', west: '|', expected: 'F'},
		{north: 'L', south: 'J', east: 'J', west: '|', expected: 'F'},
	}
	for _, test := range tests {
		result := TranslateStartPipe(test.north, test.south, test.east, test.west)
		if result != test.expected {
			t.Errorf("Expected '%v' but got '%v'", string(test.expected), string(result))
		}
	}

}

func TestGetNextDirection(t *testing.T) {
	tests := []struct {
		Pipe          byte
		PrevDirection int
		Expected      int
	}{
		{'|', NORTH, NORTH},
		{'|', SOUTH, SOUTH},
		{'-', EAST, EAST},
		{'-', WEST, WEST},
		{'L', SOUTH, EAST},
		{'L', WEST, NORTH},
		{'J', SOUTH, WEST},
		{'J', EAST, NORTH},
		{'7', EAST, SOUTH},
		{'7', NORTH, WEST},
		{'F', WEST, SOUTH},
		{'F', NORTH, EAST},
	}
	for _, test := range tests {
		result := GetNextDirection(test.Pipe, test.PrevDirection)
		if result != test.Expected {
			t.Errorf("Expected %v but got %v", DIRECTION_LOOKUP[test.Expected], DIRECTION_LOOKUP[result])
		}
	}
}
