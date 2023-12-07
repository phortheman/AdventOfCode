package main

import (
	file "aoc23/internal"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(EXAMPLE)
	timeData := ParseData(input[0])
	distanceData := ParseData(input[1])
	total := Part1Solver(timeData, distanceData)
	var expected int = 288
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}

func TestPartTwo(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(EXAMPLE)
	timeData := ParseData(input[0])
	distanceData := ParseData(input[1])
	total := Part2Solver(timeData, distanceData)
	var expected int = 71503
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}

func TestParseData(t *testing.T) {
	tests := []struct {
		time     int
		distance int
	}{
		{7, 9},
		{15, 40},
		{30, 200},
	}
	content := file.Read_String_Into_Byte_Slice(EXAMPLE)
	timeData := ParseData(content[0])
	distanceData := ParseData(content[1])
	if len(timeData) != len(distanceData) {
		t.Errorf("Data mis-match. Time len: %v | Distance len: %v", len(timeData), len(distanceData))
		return
	}
	for i, test := range tests {
		if timeData[i] != test.time {
			t.Errorf("Expected %v, but got %v", test.time, timeData[i])
		}
		if distanceData[i] != test.distance {
			t.Errorf("Expected %v, but got %v", test.distance, distanceData[i])
		}
	}
}

func TestQuadraticForm(t *testing.T) {
	tests := []struct {
		time     int
		distance int
		expected int
	}{
		{7, 9, 4},
		{15, 40, 8},
		{30, 200, 9},
	}
	for _, test := range tests {
		result := QuadraticFormula(test.time, test.distance)
		if result != test.expected {
			t.Errorf("Expected %v, but got %v", test.expected, result)
		}
	}
}
