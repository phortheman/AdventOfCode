package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const relative_input string = "../../../inputs/2024/02/input.txt"
const sampleInput = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

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
	part1, part2 := Solver(string(content))
	duration := time.Since(start)

	fmt.Printf("Time: %v\nPart 1: %d\nPart 2: %d\n", duration, part1, part2)
}

// Safely remove the element at the specified index. Create a new slice leaving the input unmodified
func removeIndex[T any](idx int, input []T) []T {
	if idx < 0 || idx >= len(input) {
		return input
	}

	// Create a new slice to return. Odd behavior with append
	output := make([]T, 0, len(input)-1)
	output = append(output, input[:idx]...)
	output = append(output, input[idx+1:]...)

	return output
}

// Start the recursive call to see if the report is safe. Starts with the first index and unknown direction
func isSafe(input []string) bool {
	return _isSafe(0, 0, input)
}

/*
	Do not use. Call 'isSafe' instead

idx: the current index we are checking

direction: 1 means we are ascending. -1 means we are desending. 0 means we don't know yet

input: the report
*/
func _isSafe(idx, direction int, input []string) bool {
	// If we're at the end then it is safe
	if idx == len(input)-1 {
		return true
	}

	currentElement, _ := strconv.Atoi(input[idx])
	nextElement, _ := strconv.Atoi(input[idx+1])

	// Get the difference between the current element and the next one
	diff := nextElement - currentElement

	switch diff {
	// No difference is unsafe
	case 0:
		return false

	// At least 1 and at most 3 difference
	case 1, 2, 3:
		if direction == 0 {
			direction = 1
		}

	case -1, -2, -3:
		if direction == 0 {
			direction = -1
		}

	// Everything else is unsafe
	default:
		return false
	}

	// If we are suppose to increase and it is decreasing, it is unsafe
	if direction == 1 && diff < 0 {
		return false
	}

	// If we are suppose to decrease and it is increasing, it is unsafe
	if direction == -1 && diff > 0 {
		return false
	}

	// Move the index over one and pass the expected direction
	return _isSafe(idx+1, direction, input)
}

func Solver(input string) (int, int) {
	safeReports := 0
	safeReportsWithTolerance := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Split(line, " ")
		if isSafe(report) {
			safeReports += 1
		} else {
			for i := range report {
				modifiedReport := removeIndex(i, report)
				if isSafe(modifiedReport) {
					safeReportsWithTolerance += 1
					break
				}
			}
		}
	}

	return safeReports, safeReports + safeReportsWithTolerance
}
