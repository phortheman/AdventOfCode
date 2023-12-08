package main

import (
	file "aoc23/internal"
	"fmt"
	"os"
	"strings"
)

var EXAMPLE string = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func main() {
	var content [][]byte
	if len(os.Args) != 2 {
		content = file.Read_String_Into_Byte_Slice(EXAMPLE)
	} else {
		var err error
		content, err = file.Read_File_Into_Memory(os.Args[1])
		if err != nil {
			fmt.Println("Error reading file into memory: ", err)
			os.Exit(1)
		}
	}

	nodeMap := make(map[string]Node)
	startNodes := make([]string, 0)
	instructions := content[0]

	for _, line := range content[2:] {
		startNodes = NewNode(nodeMap, line, startNodes)
	}

	var part1Total int = Solver("AAA", "ZZZ", instructions, nodeMap)

	steps := make([]int, 0, len(startNodes))
	for _, startKey := range startNodes {
		steps = append(steps, Solver(startKey, "", instructions, nodeMap))
	}

	part2Total := 1
	for _, n := range steps {
		part2Total = LeastCommonMultiple(part2Total, n)
	}

	fmt.Println(part1Total)
	fmt.Println(part2Total)
}

func Solver(startKey, endKey string, instructions []byte, nodeMap map[string]Node) int {
	curKey := startKey
	i := 0
	steps := 0
	for {
		if endKey != "" && curKey == endKey {
			break
		}
		if endKey == "" && curKey[2] == 'Z' {
			break
		}
		direction := instructions[i]
		switch direction {
		case 'L':
			curKey = nodeMap[curKey].Left
		case 'R':
			curKey = nodeMap[curKey].Right
		}
		i++
		steps++
		if i >= len(instructions) {
			i = 0
		}
	}
	return steps
}

type Node struct {
	Left  string
	Right string
}

func NewNode(nodeMap map[string]Node, input []byte, startNodes []string) []string {
	s := string(input)
	key, s, _ := strings.Cut(s, " = ")
	if _, ok := nodeMap[key]; ok {
		return startNodes
	}
	s = strings.Trim(s, "()")
	values := strings.Split(s, ", ")
	if len(values) != 2 {
		fmt.Print("Unexpected error. To make a node we need to have exactly 2 strings return from the space split")
		os.Exit(1)
	}
	nodeMap[key] = Node{
		Left:  values[0],
		Right: values[1],
	}
	if strings.HasSuffix(key, "A") {
		startNodes = append(startNodes, key)
	}
	return startNodes
}

func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LeastCommonMultiple(a, b int) int {
	return (a * b) / GreatestCommonDivisor(a, b)
}
