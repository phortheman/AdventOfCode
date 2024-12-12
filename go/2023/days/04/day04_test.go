package main

import (
	file "aoc23/internal"
	"bytes"
	"testing"
)

func TestPartOne(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(EXAMPLE)
	total, _ := PartSolver(input)
	var expected int = 13
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
func TestPartTwo(t *testing.T) {
	input := file.Read_String_Into_Byte_Slice(EXAMPLE)
	_, total := PartSolver(input)
	var expected int = 30
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}

func TestBytesToUint(t *testing.T) {
	input := []byte("11  3 55")
	var total int
	for _, b := range bytes.Split(input, []byte(" ")) {
		total += BytesToInt(b)
	}
	var expected int = 69
	if total != expected {
		t.Errorf("Expected %v and got %v", expected, total)
	}
}
