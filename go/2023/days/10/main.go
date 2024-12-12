package main

import (
	file "aoc23/internal"
	"fmt"
	"os"
)

var EXAMPLE string = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

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

	tiles, start := MakeGraph(content)
	var part1Total int = Part1Solver(tiles, start)

	// Part 2 might be a flood fill algorithm? Need to research
	//var part2Total int = Part2Solver(content)

	fmt.Println(part1Total)
	//fmt.Println(part2Total)
}

func Part1Solver(tiles Graph, start Point) int {
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
