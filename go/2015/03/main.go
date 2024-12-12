package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

const relative_input string = "../../../inputs/2015/03/input.txt"
const sampleInput = `<SAMPLE_INPUT>`

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
	part1 := make(grid)
	pointPart1 := newPoint()
	part1.addPoint(pointPart1)

	part2 := make(grid)
	pointsPart2 := [2]point{newPoint(), newPoint()}
	part2.addPoint(pointsPart2[0])
	part2.addPoint(pointsPart2[1])

	for idx, char := range input {
		pointPart1.changePoint(char)
		part1.addPoint(pointPart1)
		pointsPart2[idx%2].changePoint(char)
		part2.addPoint(pointsPart2[idx%2])
	}

	return len(part1), len(part2)
}

// Key is "x,y". Value is the number of times that point was passed
type grid map[string]int

func (g *grid) addPoint(p point) {
	(*g)[p.getPosition()] += 1
}

type point struct {
	x int
	y int
}

func newPoint() point {
	return point{0, 0}
}

func (p *point) changePoint(instruction byte) {
	switch instruction {
	case '<':
		p.x -= 1
	case '>':
		p.x += 1
	case '^':
		p.y += 1
	case 'v':
		p.y -= 1
	}
}

func (p *point) getPosition() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}
