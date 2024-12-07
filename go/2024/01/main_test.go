package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 11
	result, _ := Solver(sampleInput)
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 31
	_, result := Solver(sampleInput)
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}
