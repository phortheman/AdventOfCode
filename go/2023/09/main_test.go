package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 114
	result, _ := Solver([]byte(sampleInput))
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 2
	_, result := Solver([]byte(sampleInput))
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}
