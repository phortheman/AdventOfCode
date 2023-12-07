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

func TestGetDistance(t *testing.T) {
	tests := []struct {
		held     int
		expected int
	}{
		{0, 0},
		{1, 6},
		{2, 10},
		{3, 12},
		{4, 12},
		{5, 10},
		{6, 6},
		{7, 0},
		{8, 0},
	}
	maxTime := 7
	for _, test := range tests {
		result := GetDistance(test.held, maxTime)
		if result != test.expected {
			t.Errorf("Expected %v, but got %v", test.expected, result)
		}
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
