package main

import (
	file "aoc23/internal"
	"testing"
)

func TestPartOne(t *testing.T) {
	example := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`
	input := file.Read_String_Into_Byte_Slice(example)
	nodeMap := make(map[string]Node)
	startNodes := make([]string, 0)
	instructions := input[0]
	for _, line := range input[2:] {
		startNodes = NewNode(nodeMap, line, startNodes)
	}
	total := Solver("AAA", "ZZZ", instructions, nodeMap)
	var expected int = 2
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
func TestPartTwo(t *testing.T) {
	example := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
	input := file.Read_String_Into_Byte_Slice(example)
	nodeMap := make(map[string]Node)
	startNodes := make([]string, 0)
	instructions := input[0]
	for _, line := range input[2:] {
		startNodes = NewNode(nodeMap, line, startNodes)
	}
	steps := make([]int, 0, len(startNodes))
	for _, startKey := range startNodes {
		steps = append(steps, Solver(startKey, "", instructions, nodeMap))
	}
	total := 1
	for _, n := range steps {
		total = LeastCommonMultiple(total, n)
	}
	var expected int = 6
	if total != expected {
		t.Errorf("Expected %d and got %d", expected, total)
	}
}
