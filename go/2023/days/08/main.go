package main

import (
	file "aoc23/internal"
	"fmt"
	"os"
	"strings"
)

func main() {
	var content [][]byte
	if len(os.Args) != 2 {
		fmt.Println("This day didn't have a common example input. Please pass in the input")
		os.Exit(1)
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

	// Create the map
	for _, line := range content[2:] {
		startNodes = NewNode(nodeMap, line, startNodes)
	}

	var part1Total int = Solver("AAA", "ZZZ", instructions, nodeMap)

	// Get all of the min steps for each start node to find an end node
	steps := make([]int, 0, len(startNodes))
	for _, startKey := range startNodes {
		steps = append(steps, Solver(startKey, "", instructions, nodeMap))
	}

	// Calculate the least common multiple of all the steps
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
		// Case where the end key was specified
		if endKey != "" && curKey == endKey {
			break
		}
		// Case where the end key is assumed as a key ending with 'Z'
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
	key, left, right := string(input[:3]), string(input[7:10]), string(input[12:15])
	nodeMap[key] = Node{
		Left:  left,
		Right: right,
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
