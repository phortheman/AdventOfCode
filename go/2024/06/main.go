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

const relative_input string = "../../../inputs/2024/06/input.txt"
const sampleInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

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
	gridMap := newGrid(input)
	gridMap.patrol()

	return gridMap.visited(), gridMap.obstructions()
}

func formatCoords(x, y int) string {
	return fmt.Sprintf("%d, %d", x, y)
}

var _directions = [4][2]int{
	{0, -1}, // North
	{1, 0},  // East
	{0, 1},  // South
	{-1, 0}, // West
}

func getDirectionName(direction int) string {
	switch direction {
	case 0:
		return "North"
	case 1:
		return "East"
	case 2:
		return "South"
	case 3:
		return "West"
	}
	return ""
}

// Get the next direction after rotating 90 degrees right
func getNextDirection(curDirection int) int {
	if curDirection == len(_directions)-1 {
		return 0
	}
	return curDirection + 1
}

// Return the new x, y coord for taking a step in the specified direction
func step(x, y, direction int) (int, int) {
	return x + _directions[direction][0], y + _directions[direction][1]
}

type grid struct {
	GridMap        [][]byte
	GuardX         int
	GuardY         int
	GuardDirection int
	Visited        map[string]bool
	Obstructions   map[string]bool
	StartX         int
	StartY         int

	debugToggle bool // false = part1, true = part2
}

// Grid constructor
func newGrid(input []byte) grid {
	grid := grid{}
	rows := bytes.Split(input, []byte("\n"))
	grid.GridMap = make([][]byte, len(rows))
	for y, row := range rows {
		if len(row) == 0 {
			continue
		}
		grid.GridMap[y] = make([]byte, len(row))
		for x, col := range row {
			grid.GridMap[y][x] = col
			if col == '^' {
				grid.GuardX, grid.GuardY = x, y
				grid.StartX, grid.StartY = x, y
			}
		}
	}
	grid.GuardDirection = 0

	grid.Visited = make(map[string]bool)
	grid.Obstructions = make(map[string]bool)

	return grid
}

// Return the number of visited tiles in the grid
func (g *grid) visited() int {
	return len(g.Visited)
}

// Return the number of obstructions that can be placed
func (g *grid) obstructions() int {
	count := 0
	for _, val := range g.Obstructions {
		if val {
			count++
		}
	}
	return count
}

// String representation of the grid along with additional debug data
func (g grid) String() string {
	builder := strings.Builder{}
	builder.WriteString("Map:\n")
	builder.WriteString(fmt.Sprintf("Guard Position: %s\n", formatCoords(g.GuardX+1, g.GuardY+1)))
	builder.WriteString(fmt.Sprintf("Guard Visited: %d\n", g.visited()))
	builder.WriteString(fmt.Sprintf("# of Obstructions: %d\n", g.obstructions()))
	for y, row := range g.GridMap {
		for x, col := range row {
			coord := formatCoords(x, y)
			if y == g.GuardY && x == g.GuardX {
				builder.WriteByte('^')
				continue
			}

			// Indicates where the guard started
			if x == g.StartX && y == g.StartY {
				builder.WriteByte('*')
				continue
			}

			// Visualize what coords were visited
			if _, ok := g.Visited[coord]; ok && !g.debugToggle {
				builder.WriteByte('X')
				continue
			}

			// Visualize where obstructions where places
			if g.debugToggle && g.Obstructions[coord] {
				builder.WriteByte('O')
				continue
			}

			builder.WriteByte(col)
		}
		builder.WriteRune('\n')
	}
	return builder.String()
}

// Check if the given x,y coords are within the grid bounds
func (g *grid) inbounds(x, y int) bool {
	return !(x < 0 || y < 0 || y >= len(g.GridMap) || x >= len(g.GridMap[y]))
}

// Safely get the value at the specified x and y coords. Return nul byte if out of bounds
func (g *grid) at(x, y int) byte {
	if g.inbounds(x, y) {
		return g.GridMap[y][x]
	}
	return byte(0)
}

// Return the next X, Y coord for the guard
func (g *grid) nextStep() (int, int) {
	return step(g.GuardX, g.GuardY, g.GuardDirection)
}

// Rotate the guard's direction. 90 degrees right
func (g *grid) rotate() {
	g.GuardDirection = getNextDirection(g.GuardDirection)
}

/*
	Main simulation

# Once the guard is out of bounds then the guard has completed its cycle for part 1
*/
func (g *grid) patrol() {
	for g.inbounds(g.GuardX, g.GuardY) {
		g.Visited[formatCoords(g.GuardX, g.GuardY)] = true

		g.scanForLoop()

		nextX, nextY := g.nextStep()
		if g.at(nextX, nextY) == '#' {
			g.rotate()
			continue
		}

		g.GuardX, g.GuardY = nextX, nextY
	}
}

/*
	Scan the next direction line as if the guard turned at this spot.

Returns true if there is a '#' in the path indicating that placing a 'O'
in front of the current position could create a loop.
*/
func (g *grid) scanNextDirectionLine() bool {
	x, y, direction := g.GuardX, g.GuardY, g.GuardDirection
	direction = getNextDirection(direction)
	for g.inbounds(x, y) {
		x, y = step(x, y, direction)
		if g.at(x, y) == '#' {
			return true
		}
	}
	return false
}

/*
Scan if placing an obstruction will create a loop
*/
func (g *grid) scanForLoop() {

	// Get the coords for where the 'O' would be placed
	oX, oY := step(g.GuardX, g.GuardY, g.GuardDirection)

	// Can't place the 'O' out of bounds
	if !g.inbounds(oX, oY) {
		return
	}

	// If the spot we'd put a 'O' is already a '#' we won't test it
	if g.at(oX, oY) == '#' {
		return
	}

	// Check if we've already tried these coords
	oCoords := formatCoords(oX, oY)
	if _, ok := g.Obstructions[oCoords]; ok {
		return
	}

	// If there isn't a '#' along the line then we won't test it
	if !g.scanNextDirectionLine() {
		g.Obstructions[oCoords] = false
		return
	}

	// Initalize local variables
	x, y, direction := g.StartX, g.StartY, 0
	bLoop := false

	// Traveling over the same points going in the same direction indicates a loop
	traversalCache := make(map[string]bool)

	for g.inbounds(x, y) {

		// Normal traversal loop logic
		nextX, nextY := step(x, y, direction)
		if !g.inbounds(nextX, nextY) {
			break
		}

		// If we are where we would put the 'O' treat it like a '#'
		if g.at(nextX, nextY) == '#' || oX == nextX && oY == nextY {

			// If we have already traveled to this spot in the same direction we looped
			key := fmt.Sprintf("%s | %s", formatCoords(x, y), getDirectionName(direction))
			if _, ok := traversalCache[key]; ok {
				bLoop = true
				break
			}
			traversalCache[key] = true
			direction = getNextDirection(direction)
			continue
		}

		// Cache the traversal and move along
		x, y = nextX, nextY
	}

	// Cache the result
	g.Obstructions[oCoords] = bLoop
}
