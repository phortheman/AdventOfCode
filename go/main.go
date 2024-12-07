package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/<YEAR>/<DAY>/input.txt"
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

	// NOTE: Default to converting puzzle input into a string representation of the whole file.
	// Change as needed for the puzzle
	start := time.Now()
	part1, part2 := Solver(string(content))
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

// NOTE: Base structure for the solver. Should handle the line by line operation for those puzzles
func Solver(input string) (int, int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	return -1, -1
}
