package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 41
	result, _ := Solver([]byte(sampleInput))

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 6
	_, result := Solver([]byte(sampleInput))

	if result != expected {
		t.Errorf("Expected %d but got %d.", expected, result)
	}
}
