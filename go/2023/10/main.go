package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

const relative_input string = "../../../inputs/2023/10/input.txt"
const sampleInput = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

var bTest = false
var inputFileName string

func init() {
	flag.BoolVar(&bTest, "t", false, "Run using the sample input")
	flag.StringVar(&inputFileName, "i", "",
		"Path to the puzzle input. "+
			"Default to using the internal relative path. "+
			"Pass 'stdin' to use it instead")
}

func main() {
	flag.Parse()
	var content []byte = []byte(sampleInput)
	var err error
	if !bTest {
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
	}

	start := time.Now()
	part1, part2 := Solver(content)
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input []byte) (int, int) {
	newInput := bytes.Split(input, []byte("\n"))

	part1, part2 := 0, 0

	tiles, start := MakeGraph(newInput)

	part1 = Part1(tiles, start)
	// Part 2 might be a flood fill algorithm? Need to research

	return part1, part2
}

func Part1(tiles Graph, start Point) int {
	var prevDirectionA, prevDirectionB int

	switch tiles[start] {
	case '|':
		prevDirectionA = NORTH
		prevDirectionB = SOUTH
	case '-':
		prevDirectionA = EAST
		prevDirectionB = WEST
	case 'L':
		prevDirectionA = NORTH
		prevDirectionB = EAST
	case 'J':
		prevDirectionA = NORTH
		prevDirectionB = WEST
	case '7':
		prevDirectionA = SOUTH
		prevDirectionB = WEST
	case 'F':
		prevDirectionA = SOUTH
		prevDirectionB = EAST
	default:
		prevDirectionA = -1
		prevDirectionB = -1
	}

	startA := GetNextPoint(start, prevDirectionA)
	startB := GetNextPoint(start, prevDirectionB)

	return DepthFirstSearch(tiles, startA, startB, prevDirectionA, prevDirectionB, 1)
}

// Returns the graph and the start point
func MakeGraph(input [][]byte) (Graph, Point) {
	outputGraph := make(Graph)
	outputStart := Point{}
	for i, line := range input {
		if len(line) == 0 {
			continue
		}
		for j, pipe := range line {
			p := Point{
				x: j,
				y: i,
			}
			if pipe == 'S' {
				var north, south, east, west byte
				if i-1 < 0 {
					north = '.'
				} else {
					north = input[i-1][j]
				}

				if i+1 >= len(input) {
					south = '.'
				} else {
					south = input[i+1][j]
				}

				if j+1 >= len(input[i]) {
					east = '.'
				} else {
					east = input[i][j+1]
				}

				if j-1 < 0 {
					west = '.'
				} else {
					west = input[i][j-1]
				}
				pipe = TranslateStartPipe(north, south, east, west)
				outputStart = p
			}
			outputGraph[p] = pipe
		}
	}
	return outputGraph, outputStart
}

func TranslateStartPipe(north, south, east, west byte) byte {
	bNorth, bSouth, bEast, bWest := false, false, false, false
	var pipe byte = '.'
	if north == '|' || north == '7' || north == 'F' {
		bNorth = true
	}
	if south == '|' || south == 'L' || south == 'J' {
		bSouth = true
	}
	if east == '-' || east == 'J' || east == '7' {
		bEast = true
	}
	if west == 'L' || west == 'F' || west == '-' {
		bWest = true
	}
	if bNorth && bSouth {
		pipe = '|'
	}
	if bEast && bWest {
		pipe = '-'
	}
	if bNorth && bEast {
		pipe = 'L'
	}
	if bNorth && bWest {
		pipe = 'J'
	}
	if bSouth && bWest {
		pipe = '7'
	}
	if bSouth && bEast {
		pipe = 'F'
	}
	return pipe
}

func DepthFirstSearch(graph Graph, curA, curB Point, prevDirectionA, prevDirectionB int, currentDistance int) int {
	if curA == curB {
		return currentDistance
	}

	nextDirectionA := GetNextDirection(graph[curA], prevDirectionA)
	nextDirectionB := GetNextDirection(graph[curB], prevDirectionB)

	nextA := GetNextPoint(curA, nextDirectionA)
	nextB := GetNextPoint(curB, nextDirectionB)

	return DepthFirstSearch(graph, nextA, nextB, nextDirectionA, nextDirectionB, currentDistance+1)
}

type Point struct {
	x int
	y int
}

type Graph map[Point]byte

func GetNextPoint(p Point, direction int) Point {
	if direction == NORTH {
		return Point{p.x, p.y - 1}
	} else if direction == SOUTH {
		return Point{p.x, p.y + 1}
	} else if direction == EAST {
		return Point{p.x + 1, p.y}
	} else if direction == WEST {
		return Point{p.x - 1, p.y}
	} else {
		fmt.Println("Unsupported direction:", direction)
		os.Exit(1)
	}
	return Point{}
}

const NORTH = 0
const SOUTH = 1
const EAST = 2
const WEST = 3

var DIRECTION_LOOKUP map[int]string = map[int]string{
	NORTH: "NORTH",
	SOUTH: "SOUTH",
	EAST:  "EAST",
	WEST:  "WEST",
}

func GetNextDirection(pipe byte, prevDirection int) int {
	direction := -1
	switch pipe {
	case '|':
		if prevDirection == SOUTH {
			direction = SOUTH
		} else if prevDirection == NORTH {
			direction = NORTH
		}
	case '-':
		if prevDirection == WEST {
			direction = WEST
		} else if prevDirection == EAST {
			direction = EAST
		}
	case 'L':
		if prevDirection == SOUTH {
			direction = EAST
		} else if prevDirection == WEST {
			direction = NORTH
		}
	case 'J':
		if prevDirection == SOUTH {
			direction = WEST
		} else if prevDirection == EAST {
			direction = NORTH
		}
	case '7':
		if prevDirection == EAST {
			direction = SOUTH
		} else if prevDirection == NORTH {
			direction = WEST
		}
	case 'F':
		if prevDirection == WEST {
			direction = SOUTH
		} else if prevDirection == NORTH {
			direction = EAST
		}
	default:
		fmt.Printf("Unexpected GetNextDirection result for %v and prevDirection %v\n", string(pipe), prevDirection)
	}
	return direction
}
