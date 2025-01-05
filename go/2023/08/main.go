package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/2023/08/input.txt"

var inputFileName string

func init() {
	flag.StringVar(&inputFileName, "i", "",
		"Path to the puzzle input. "+
			"Default to using the internal relative path. "+
			"Pass 'stdin' to use it instead")
}

func main() {
	flag.Parse()
	var content []byte
	var err error
	switch inputFileName {
	case "stdin":
		content, err = io.ReadAll(os.Stdin)
	case "":
		content, err = os.ReadFile(relative_input)
	default:
		content, err = os.ReadFile(inputFileName)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read file: %v", err)
		os.Exit(66)
	}

	start := time.Now()
	part1, part2 := Solver(content)
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input []byte) (int, int) {
	content := bytes.Split(input, []byte("\n"))
	nodeMap := make(map[string]Node)
	startNodes := make([]string, 0)
	instructions := content[0]

	// Create the map
	for _, line := range content[2:] {
		if len(line) == 0 {
			continue
		}
		startNodes = NewNode(nodeMap, line, startNodes)
	}

	var part1Total int = Step("AAA", "ZZZ", instructions, nodeMap)

	// Get all of the min steps for each start node to find an end node
	steps := make([]int, 0, len(startNodes))
	for _, startKey := range startNodes {
		steps = append(steps, Step(startKey, "", instructions, nodeMap))
	}

	// Calculate the least common multiple of all the steps
	part2Total := 1
	for _, n := range steps {
		part2Total = LeastCommonMultiple(part2Total, n)
	}
	return part1Total, part2Total
}

func Step(startKey, endKey string, instructions []byte, nodeMap map[string]Node) int {
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
