package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

const relative_input string = "../../../inputs/2015/01/input.txt"

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
	part1, part2 := Solver(content)
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

func Solver(input []byte) (int, int) {
	part1 := 0
	part2 := 0
	inBasement := false

	for _, char := range input {
		if !inBasement {
			part2 += 1
		}
		if char == '(' {
			part1 += 1
		} else if char == ')' {
			part1 -= 1
		}
		if part1 < 0 {
			inBasement = true
		}
	}
	return part1, part2
}
