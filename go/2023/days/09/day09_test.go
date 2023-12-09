package main

import (
	file "aoc23/internal"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(EXAMPLE)
	total, _ := Solver(input)
	var expected int = 114
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
func TestPartTwo(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(EXAMPLE)
	_, total := Solver(input)
	var expected int = 2
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
