package main

import (
	"bytes"
	"testing"
)

func TestPart1(t *testing.T) {
	input := bytes.Split([]byte(sampleInput), []byte("\n"))
	seeds, data := ParseData(input)
	result := Part1(seeds, data)
	var expected int = 35
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := bytes.Split([]byte(sampleInput), []byte("\n"))
	seeds, data := ParseData(input)
	result := Part2(seeds, data)
	var expected int = 46
	if result != expected {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestSourceDestinationLogic(t *testing.T) {
	input := SourceDestination{}
	input.AddData(98, 50, 2)
	input.AddData(50, 52, 48)
	result := input.GetDestination(79)
	expected := 81
	if expected != result {
		t.Errorf("Expected %d and got %d", expected, result)
	}

	result = input.GetDestination(14)
	expected = 14
	if expected != result {
		t.Errorf("Expected %d and got %d", expected, result)
	}

	result = input.GetDestination(55)
	expected = 57
	if expected != result {
		t.Errorf("Expected %d and got %d", expected, result)
	}

	result = input.GetDestination(13)
	expected = 13
	if expected != result {
		t.Errorf("Expected %d and got %d", expected, result)
	}
}

func TestByteContains(t *testing.T) {
	input := "soil-to-fertilizer map:"
	lookup := "map:"
	result := ByteContains([]byte(input), lookup)
	if !result {
		t.Errorf("Expected to find '%v' in '%v'", lookup, input)
	}
}
