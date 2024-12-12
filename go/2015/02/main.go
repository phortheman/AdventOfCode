package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/2015/02/input.txt"

var bTest = false
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
	part1, part2 := Solver(string(content))
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input string) (int, int) {
	part1 := 0
	part2 := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		val := strings.Split(line, "x")
		rec := newRectangle(val[0], val[1], val[2])
		part1 += rec.getDimension()
		part2 += rec.getRibbon()
	}

	return part1, part2
}

type rectangle struct {
	length int
	width  int
	height int
}

func newRectangle(length, width, height string) rectangle {
	rectangle := rectangle{}
	rectangle.length, _ = strconv.Atoi(length)
	rectangle.width, _ = strconv.Atoi(width)
	rectangle.height, _ = strconv.Atoi(height)

	return rectangle
}

// Surface area of the rectangle plus some slack for the wrapping paper
func (r *rectangle) getDimension() int {
	var sum int
	dimensions := []int{
		r.length * r.width,
		r.width * r.height,
		r.height * r.length,
	}

	slack := dimensions[0]

	for _, dim := range dimensions {
		if slack > dim {
			slack = dim
		}
		sum += 2 * dim
	}

	return sum + slack
}

// Find the shortest distance to wrap the ribbon around then add l*w*h
func (r *rectangle) getRibbon() int {
	sum := r.length * r.width * r.height
	distances := []int{
		2*r.length + 2*r.width,
		2*r.width + 2*r.height,
		2*r.height + 2*r.length,
	}
	return sum + slices.Min(distances)
}
