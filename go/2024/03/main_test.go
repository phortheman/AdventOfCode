package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 161
	result, _ := Solver(sampleInput1)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 48
	_, result := Solver(sampleInput2)

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}
