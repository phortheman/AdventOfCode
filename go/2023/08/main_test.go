package main

import (
	"bytes"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`, 2},
		{`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`, 6},
	}
	for idx, test := range tests {
		input := bytes.Split([]byte(test.input), []byte("\n"))
		nodeMap := make(map[string]Node)
		startNodes := make([]string, 0)
		instructions := input[0]
		for _, line := range input[2:] {
			if len(line) == 0 {
				continue
			}
			startNodes = NewNode(nodeMap, line, startNodes)
		}

		total := Step("AAA", "ZZZ", instructions, nodeMap)

		if total != test.expected {
			t.Errorf("Test %d: Expected %d and got %d", idx+1, test.expected, total)
		}
	}
}

func TestPart2(t *testing.T) {
	input := bytes.Split([]byte(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`), []byte("\n"))
	nodeMap := make(map[string]Node)
	startNodes := make([]string, 0)
	instructions := input[0]
	for _, line := range input[2:] {
		if len(line) == 0 {
			continue
		}
		startNodes = NewNode(nodeMap, line, startNodes)
	}

	steps := make([]int, 0, len(startNodes))
	for _, startKey := range startNodes {
		steps = append(steps, Step(startKey, "", instructions, nodeMap))
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
