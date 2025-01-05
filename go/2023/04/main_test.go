package main

import (
	"bytes"
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 13
	result, _ := Solver([]byte(sampleInput))
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	expected := 30
	_, result := Solver([]byte(sampleInput))
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestBytesToInt(t *testing.T) {
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
